## v0.2.0
* Dramatically improve the start-up performance by removing the installation of the _vim-go_ plug from the vimrc file with sed.

## v0.1.0
* First fully automated version. If the API token is placed in the right place, it should all work fully automated.
* [BUG] It still takes to long after the first time it is set up, to get into vim, because _vim-go_ tries to update all required binaries.
