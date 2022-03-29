#!/bin/sh

OS=''
#########################################################################
MAIN_USER="hap"
# lf-26 is the Go terminal file manager
PACKAGES2INSTALL="git vim curl go lf-26"
# hey is used to send load to web applications
WEBDEVPACKAGES="mariadb105-server-10.5.15 hey-0.1.4_1"
#########################################################################
# Colors for the text (they do not set the background color)
COLOR_GREEN='\e[0;32m'
NO_COLOR='\033[0m'
COLOR_RED='\033[0;31m'
COLOR_MAGENTA='\033[0;35m'
COLOR_CYAN='\033[0;36m'

#########################################################################

# Print an info message to stdout. First parameter is string to be logged
print_info(){
	local INFO_COLOR=${COLOR_CYAN}
	printf "[${INFO_COLOR}INFO${NO_COLOR}] $1\n"
}

# Print an error message to stdout. First parameter is string to be logged
print_error(){
	printf "[${COLOR_RED}ERROR${NO_COLOR}] $1\n"
}

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
		OS="debian"
		return
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
	pkg install --yes ${WEBDEVPACKAGES} && print_info "${WEBDEVPACKAGES} were successfully installed!" || return -1

	# Installing mariaDB on FreeBSD:
	# https://www.osradar.com/how-to-install-mariadb-on-freebsd-12/

	# Enable the mariadb service to start with the system
	sysrc mysql_enable="yes"

	# Start mariadb service
	service mysql-server start

	# Run the mysql configuration script
	/usr/local/bin/mysql_secure_installation

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
		pkg install --yes ${PACKAGES2INSTALL} && print_info "${PACKAGES2INSTALL} were successfully installed!" || { print_error "Packages installation failed!"; exit -1; }
		
		install_webdev_packages  || { print_error "Webdev packages installation failed!"; exit -1; }

		return
	fi

	if [ ${OS} = "debian" ]; then
		return	
	fi

}

# configure vim, its plugins, tmux, lf, all dotfiles
configure_dotfiles(){
	# Install vim-plug to handle plugins
	# Reference: https://github.com/junegunn/vim-plug
	curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
		    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

	local REPO_PATH="/dotfiles"
	# change to home directory
	cd ${HOME}
	git clone https://github.com/erodrigufer/dotfiles.git || { rm -fR ${HOME}${REPO_PATH}; print_error "Cloning GitHub dotfiles repo failed!"; exit -1; }
	cd ${HOME}${REPO_PATH}
	make || { rm -fR ${HOME}${REPO_PATH}; print_error "Creating symlinks to dotfiles failed!"; exit -1; }
	
	# we do not need the zshrc file
	rm -f ${HOME}/.zshrc
	
}

# Configure sshd properly, do not disconnect client after short idle time
configure_ssh(){
	cd /etc/ssh
	# Create a copy of the sshd original configuration, before changing it
	cp ./sshd_config ./sshd_config.bak

	# Change the configuration of sshd, so that it does not automatically close
	# after the client has been idle for a while, it will now only disconnect
	# after 2 hours of inactivity.
	# - Allow PubKey Authentication
	# Reference: https://www.simplified.guide/ssh/disable-timeout
	# (sed) -i: Replace in-place, and append commands with -e
	# replace everything after and before the pattern with .*
	sed -i '' -e 's/^.*TCPKeepAlive.*/TCPKeepAlive no/' -e 's/^.*ClientAliveInterval.*/ClientAliveInterval 30/' -e 's/^.*ClientAliveCountMax.*/ClientAliveCountMax 240/' -e 's/^.*PubkeyAuthentication.*/PubkeyAuthentication yes/' ./sshd_config	
	# Disable SSH timeout from the server. The server will not send any 
	# TCPKeepAlive packages and will only disconnect the client after 2 hours 
	# of inactivity.
	# Reference: https://www.simplified.guide/ssh/disable-timeout

	# Restart sshd with new configuration
	service sshd restart

}

main(){
	check_os
	install_bare_packages	
	configure_dotfiles
	configure_ssh
}

main

# install lf, golang, git
# setup vim, setup lf, setup dotfiles (tmux, etc)
