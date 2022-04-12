#!/bin/sh
# Eduardo Rodriguez [@erodrigufer] 2022 (c) 
# Spawn a VM in Vultr Cloud using the Go client for the Vultr API

#########################################################################
HOSTNAME="eee"
LABEL="logwatch-test"
REGION="fra"
#########################################################################

# Eliminate credentials file, if it already exists
rm -f ./vm_credentials.secrets
cd ./vultrClient
make build
source  ./secrets/vultrAPI.secrets
cd ..
./vultrClient/build/vultrClient.bin -hostname ${HOSTNAME} -label ${LABEL} -sshKey \
	${SSH_KEY} -tokenAPI ${API_TOKEN} -region ${REGION}

# Source the newly acquired credentials for the VM
source ./vm_credentials.secrets
# Ping the newly created VM every 7 seconds, until it is up and running
ping -o -i 7 ${HOST}
# SSH into the VM, do not ask to check key fingerprinting for new host
# ssh -o "StrictHostKeyChecking no" root@${HOST}
./setupFirstRun.sh
# After finishing setup, ssh into the VM
ssh ${USER}@${HOST}
