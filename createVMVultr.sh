#!/bin/sh
# Eduardo Rodriguez [@erodrigufer] 2022 (c) 
# Spawn a VM in Vultr Cloud using the Go client for the Vultr API

#########################################################################
HOSTNAME="eee"
LABEL="eee"
REGION="cdg"
#########################################################################

cd ./vultrClient
make build
source  ./secrets/vultrAPI.secrets
./build/vultrClient.bin -hostname ${HOSTNAME} -label ${LABEL} -sshKey \
	${SSH_KEY} -tokenAPI ${API_TOKEN} -region ${REGION}

