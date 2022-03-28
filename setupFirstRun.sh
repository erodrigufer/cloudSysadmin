#!/bin/sh

#########################################################################
MAIN_USER="hap"

#########################################################################

COLOR_GREEN='\e[0;32m'
NO_COLOR='\033[0m'
COLOR_RED='\033[0;31m'


# Properly log an error message to stdout
print_error(){
	printf "[${COLOR_RED}ERROR${NO_COLOR}] $1\n"
}

# Check if the script is running in a valid OS (eventual support should be for
# primary FreeBSD and Ubuntu)
check_os(){
# print Linux_Standard_Base release ID
# use grep in egrep mode (-e), being 
# case-insensitive (-i) and check for either 
# ubuntu or debian
echo "* Checking system distribution..."
lsb_release -i | grep -i -E "(ubuntu|debian|kali)" || { print_error "System must be debian-based to run script."; exit -1 ; }

# About egrep: with egrep it is easier to write the OR
# logic ( | ) otherwise all these characters must be escaped
# with the normal regex grammar of grep
}

configure_users(){
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

main(){
	check_os
}

main
