# cloudSysadmin

A Go program (client) and shell script to automate setting up VMs in the cloud(s) [Vultr]. 

The scripts should be focused on running on FreeBSD. Properly tested on **FreeBSD 13.0**

## Configuration
### ssh Keys
If an ssh key is automatically installed in the server at deployment, to use it from the client:
1. `ssh-add` : add key to keys agent
2. `ssh -i <path_of_key> root@<IP_address>` : explicitly state where the file with the key is located

* Reference: [Vultr Docs - ssh keys](https://www.vultr.com/docs/connect-to-a-server-using-an-ssh-key)

# FreeBSD: General information
## Networking
* [Check open ports on FreeBSD](https://linuxhint.com/check-open-ports-freebsd/): _sockstat_ command

## SSH
* [How to configure ssh key authentication on FreeBSD](https://www.digitalocean.com/community/tutorials/how-to-configure-ssh-key-based-authentication-on-a-freebsd-server)

## MariaDB
* [How to install MariaDB on FreeBSD](https://www.osradar.com/how-to-install-mariadb-on-freebsd-12/)
