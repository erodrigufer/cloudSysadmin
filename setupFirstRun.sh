#!/bin/sh

OS=''
#########################################################################
MAIN_USER="hap"

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
		OS="linux"
		return
	fi
	if [ $(uname) = "FreeBSD" ]; then
		OS="freebsd"
		return
	fi

}

configure_users(){
	# Check if username already exists
	
	adduser ${MAIN_USER}
	# give sudo rights to main user
	adduser ${MAIN_USER} sudo	

}

setup_sshd(){
	# Disable SSH timeout from the server. The server will not send any 
	# TCPKeepAlive packages and will only disconnect the client after 2 hours 
	# of inactivity.
	# Reference: https://www.simplified.guide/ssh/disable-timeout
	echo "Done"
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
		pkg install --yes git vim curl go lf-26
		# lf-26 is the Go terminal file manager
	fi

}

main(){
	check_os
	echo $OS
	install_bare_packages	
}

main

# install lf, golang, git
# setup vim, setup lf, setup dotfiles (tmux, etc)
