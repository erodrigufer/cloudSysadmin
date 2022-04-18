## v0.3.1
* Fix minor issue: sed was not commenting out the line for _vim-go_ initialization in the vimrc file.
## v0.3.0
* Properly implement deleting an instance into the main.go file
## v0.2.2
* Add unit test for checkResponseAPI()
## v0.2.1
* Add untested deleteInstance method, and refactor code of main.go to enable 'action' flag
## v0.2.0
* Dramatically improve the start-up performance by removing the installation of the _vim-go_ plug from the vimrc file with sed.

## v0.1.0
* First fully automated version. If the API token is placed in the right place, it should all work fully automated.
* [BUG] It still takes to long after the first time it is set up, to get into vim, because _vim-go_ tries to update all required binaries.
