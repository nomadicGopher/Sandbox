# Tools

## nmap (Network Mapper)
* `nmap -sV {target_IP}` - Probe open ports to determine service/version info

## smbclient
* **Arguments**
  * `-L` specifies the host. ie. `smbclient -L {target_IP}`
  * `-U` specified a username. If not provided the current system's username or anonymouse will be used (depending on configuration). ie. `... -U guest`
    * Can press enter to try an empty password
* To access a share: `smbclient \\\\{target_IP}\\{sharename}` which can also include the username argument.
* _Once smp shell is activated, can use `help` to see list of commands just like ftp._
  * `exit` to escape