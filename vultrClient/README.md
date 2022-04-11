# VultrClient
Go client to interact with Vultr API.

* Create instances
* Delete instances
* Get info from instances
* List all SSH keys in the system

## Usage
```bash
vultrClient [-h|-help] -tokenAPI <token> [-SSHkey] <sshkey> 

	-h -help	Display usage
	-tokenAPI	Personal Authentication Token to interact with Vultr API
	-SSHkey		SSH key to initialize per default in new instance

```
