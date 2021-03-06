#!/bin/sh
# Eduardo Rodriguez [@erodrigufer] 2022 (c) 
# This script automates setting up a FreeBSD VM in the cloud after it has been
# successfully spawned.
# - It configures the ecosystem for development; vim, go, dotfiles in general
# - It configures sshd to not close a connection so quickly, and accept public
#	key authentication
# - It installs common useful packages and, if desired, packages for web dev.

# Default values, just in case that configVM.env does not exist, or it
# even exists, but does not contain these variables. So the program would still
# work with some default values even if the .env file is missing some variables.
#########################################################################
# Default user-defined input parameters

MAIN_USER="hap"

# If the .env file is configured correctly, its variables should over-write the
# default ones.
if [ -f ./configVM.env ]; then
	. ./configVM.env
fi

if [ -z ${PACKAGES2INSTALL} ]; then 
# Package naming used by FreeBSD pkg
# lf-26 is the Go terminal file manager
# duf is a modern version of 'df', free disk utility
PACKAGES2INSTALL="git vim curl go lf duf htop"
# hey is used to send load to web applications
fi

if [ -z ${WEBDEVPACKAGES} ]; then 
WEBDEVPACKAGES="mariadb105-server hey sqlite3"
fi

#########################################################################
# Internal global variables, which should not be modified by user

OS=''
# Colors for the text (they do not set the background color)
COLOR_GREEN='\e[0;32m'
NO_COLOR='\033[0m'
COLOR_RED='\033[0;31m'
COLOR_MAGENTA='\033[0;35m'
COLOR_CYAN='\033[0;36m'

SSHD_CONFIG_FILE="./etc/sshd_config"

#########################################################################
# Internal global variables, which can be modified by user

FILE_VM_CREDENTIALS=vm_credentials.secrets # file with ssh credentials for VM
FILE_USER_CONFIG="./configVM.env" # file that describes which packages to 
# install, which instance to initiate and so on...
FILE_NAME=$0 # name of this same file
PATH_IN_VM="/root/" # where to store setup script in VM, from which it is then
					# run inside the VM after establishing ssh connection

#########################################################################
# Helper functions

# Print an info message to stdout. First parameter is string to be logged
print_info(){
	local INFO_COLOR=${COLOR_CYAN}
	printf "[${INFO_COLOR}INFO${NO_COLOR}] $1\n"
}

# Print an error message to stdout. First parameter is string to be logged
print_error(){
	printf "[${COLOR_RED}ERROR${NO_COLOR}] $1\n"
}

#########################################################################

# Check if the script is running in a valid OS (eventual support should be for
# primary FreeBSD and Ubuntu)
check_os(){
	echo "* Checking system distribution..."
	if [ $(uname) = "Linux" ]; then
		# print Linux_Standard_Base release ID use grep in egrep mode (-e), 
		# being case-insensitive (-i) and check for either ubuntu or debian
		# About egrep: with egrep it is easier to write the OR logic ( | ) 
		# otherwise all these characters must be escaped with the normal regex 
		# grammar of grep
		lsb_release -i | grep -i -E "(ubuntu|debian|kali)" || { print_error "System must be debian-based to run script."; exit -1 ; }
		print_error "{$(uname)} System not supported!"	
		OS="debian"
		exit -1
	fi
	if [ $(uname) = "FreeBSD" ]; then
		OS="freebsd"
		return
	fi

	print_error "{$(uname)} System not supported!"	
	exit -1
}

configure_users(){
	# Check if username already exists
	
	adduser ${MAIN_USER}
	# give sudo rights to main user
	adduser ${MAIN_USER} sudo	

}

install_webdev_packages(){
	pkg install --yes ${WEBDEVPACKAGES} && print_info "${WEBDEVPACKAGES} were successfully installed!" || { print_error "Webdev packages installation failed!"; return; }

	# Installing mariaDB on FreeBSD:
	# https://www.osradar.com/how-to-install-mariadb-on-freebsd-12/

	# Enable the mariadb service to start with the system
	sysrc mysql_enable="yes"

	# Start mariadb service
	service mysql-server start

	# Run the mysql configuration script
	# /usr/local/bin/mysql_secure_installation

}

# Install the bare minimum of necessary packages
install_bare_packages(){
	if [ ${OS} = "freebsd" ]; then
		
		# Update the available remote repositories
		pkg update	
		# Install packages without further confirmation (--yes)
		# For more information take a look at 'man pkg-install'
		# To get more info about a particular package run:
		# pkg search -R <name>
		pkg install --yes ${PACKAGES2INSTALL} && print_info "${PACKAGES2INSTALL} were successfully installed!" || { print_error "Packages installation failed!"; }
		
		return
	fi

	if [ ${OS} = "debian" ]; then
		return	
	fi

}

# configure vim, its plugins, tmux, lf, all dotfiles
configure_dotfiles(){
	# If git, vim or curl are not present, return without installing the dotfiles
	which git > /dev/null || return
	which curl > /dev/null || return
	which vim > /dev/null || return
	# Install vim-plug to handle plugins
	# Reference: https://github.com/junegunn/vim-plug
	curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
		    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

	local REPO_PATH="/dotfiles"
	# change to home directory
	cd ${HOME}
	git clone https://github.com/erodrigufer/dotfiles.git || { rm -fR ${HOME}${REPO_PATH}; print_error "Cloning GitHub dotfiles repo failed!"; exit -1; }
	cd ${HOME}${REPO_PATH}
	# Eliminate vim-go in the local installation
	# Reference to escape single quotes:
	# https://stackoverflow.com/questions/24509214/how-to-escape-single-quote-in-sed
	sed -i '' -e 's/^Plug.*fatih/"Plug '\''fatih/' ./vimrc
	make || { rm -fR ${HOME}${REPO_PATH}; print_error "Creating symlinks to dotfiles failed!"; exit -1; }

	# Install vim plugins in the background
	vim -E -s +PlugInstall +qall
	
	# we do not need the zshrc file
	rm -f ${HOME}/.zshrc
	
}

# Configure sshd properly, do not disconnect client after short idle time
configure_ssh(){
	# Create a copy (backup) of the sshd original configuration, before changing it
	cp /etc/ssh/sshd_config /etc/ssh/sshd_config.bak

	# Replace original file with new config file
	cp ${PATH_IN_VM}sshd_config /etc/ssh/sshd_config

	# Change the configuration of sshd, so that it does not automatically close
	# after the client has been idle for a while, it will now only disconnect
	# after 2 hours of inactivity.
	# - Allow PubKey Authentication
	# Reference: https://www.simplified.guide/ssh/disable-timeout
	# (sed) -i: Replace in-place, and append commands with -e
	# replace everything after and before the pattern with .*
	# sed -i '' -e 's/^#TCPKeepAlive.*/TCPKeepAlive no/' -e 's/^#ClientAliveInterval.*/ClientAliveInterval 30/' -e 's/^#ClientAliveCountMax.*/ClientAliveCountMax 240/' -e 's/^#PubkeyAuthentication.*/PubkeyAuthentication yes/' ./sshd_config	
	# Disable SSH timeout from the server. The server will not send any 
	# TCPKeepAlive packages and will only disconnect the client after 2 hours 
	# of inactivity.
	# Reference: https://www.simplified.guide/ssh/disable-timeout

	# Disable password authentication
	# In the default file from Vultr, there is an uncommented instance of 
	# PasswordAuthentication
	# sed -i '' -e 's/^PasswordAuthentication.*/PasswordAuthentication no/' -e 's/^#PasswordAuthentication.*/PasswordAuthentication no/' -e 's/^#ChallengeResponseAuthentication.*/ChallengeResponseAuthentication no/' ./sshd_config
	# If it is commented out
	# sed -i '' -e 's/^#PasswordAuthentication.*/PasswordAuthentication no/' ./sshd_config

	# Restart sshd with new configuration
	service sshd restart

}

# Connect through ssh to the remote VM and run all the commands there
connectVM(){
	# check if file with ssh secrets for VM exists, if so source file, transfer
    # script to VM with scp, and run the transferred script inside the VM after 	
	# establishing an ssh connection
	# if there is no file with credentials, return, since we are probably 
	# running the script inside the VM

	FILES_TO_SEND="${FILE_NAME} ${SSHD_CONFIG_FILE} ${FILE_USER_CONFIG}"
	# Use the "StrictHostKeyChecking no" ssh option, to not have to agree with
	# the key fingerprints of the server we are trying to ssh into
	[ -f ${FILE_VM_CREDENTIALS} ] && { source ${FILE_VM_CREDENTIALS}; scp -o "StrictHostKeyChecking no" ${FILES_TO_SEND} ${USER}@${HOST}:${PATH_IN_VM} && ssh ${USER}@${HOST} ${PATH_IN_VM}$(basename ${FILE_NAME}) && exit 0; } || return
}

main(){
	connectVM
	check_os
	install_bare_packages	
	#install_webdev_packages
	configure_dotfiles
	configure_ssh
}

main 
