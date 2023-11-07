# govs
Golang Version Switcher  
Easily switch between golang version.  
Inspired by [awsp](https://github.com/johnnyopao/awsp).  

![govs](https://github.com/kane8n/govs/assets/4223926/d8412642-7cad-453e-a37b-e9c3c8aced78)

## Setup
```
git clone git@github.com:kane8n/govs.git
cd govs
make install
```
Add the following to your .bashrc or .zshrc config
```
alias govs="source _govs"
export PATH=$PATH:~/.govs/bin
```
## Install other versions
![govs_install](https://github.com/kane8n/govs/assets/4223926/795a4524-f433-475c-aab8-cc5702dd338f)
It wraps the command to install multiple versions of golang described in the official GO documentation [Managing Go installations](https://go.dev/doc/manage-install).  
The list of installable versions is obtained by scraping GO's [All releases page](https://go.dev/dl/) to get all versions that match the OS/ARC you are running.
