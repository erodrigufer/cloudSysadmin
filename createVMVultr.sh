#!/bin/sh
# Eduardo Rodriguez [@erodrigufer] 2022 (c) 
# Spawn a VM in Vultr Cloud using a self-made Go client for the Vultr API
#
#########################################################################

#########################################################################
# Default values, just in case that configVM.env does not exist, or it
# even exists, but does not contain these variables. So the program would still
# work with some default values even if the .env file is missing some variables.
HOSTNAME="eee"
LABEL="newInstance"
REGION="fra"
PLAN="vc2-1c-1gb"	# cheapest/most basic plan
OS="447"	# FreeBSD-13
#########################################################################
# If the .env file is configured correctly, its variables should over-write the
# default ones.
if [ -f ./configVM.env ]; then
	. ./configVM.env
fi

# Eliminate credentials file, if it already exists
rm -f ./vm_credentials.secrets
cd ./vultrClient
make build
source  ./secrets/vultrAPI.secrets
cd ..
./vultrClient/build/vultrClient.bin -action "create" -hostname ${HOSTNAME} -label ${LABEL} -sshKey \
	${SSH_KEY} -tokenAPI ${API_TOKEN} -region ${REGION} -plan ${PLAN} -os ${OS} || exit -1

# Source the newly acquired credentials for the VM
source ./vm_credentials.secrets
# Ping the newly created VM every 10 seconds, until it is up and running
ping -o -i 10 ${HOST}
# SSH into the VM, do not ask to check key fingerprinting for new host
# ssh -o "StrictHostKeyChecking no" root@${HOST}
./setupFirstRun.sh 
# After finishing setup, ssh into the VM
ssh ${USER}@${HOST}
