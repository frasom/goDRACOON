# goDRACOON
clients for public API (upload / download shares, software information) - DRACOON (https://dracoon.com)

## About the project
Some examples of using the DRACOON API.
A description of the DRACOON API can be accessed via the link https://dracoon.team/api. Current available client SDKs and code examples are referenced in the development portal at https://developer.dracoon.com/ and distributed via the DRACOON GitHub account https://github.com/dracoon.

## Tools

### drstatus
  Querying the status and the software version of a DRACOON Server instance. The following information is displayed.  
  * Cloud System (true or false)
  * Default Language
  * API Version (Release)
  * Server Version  (Release)
  * Use S3 Storage (true or false)
  * S3 Hosts (URL)
  * Authentication methods

### drdownload
  Querying the public information of a Download Share. The following information is displayed. 
  * Filename
  * Filetype
  * Filesize
  * Creatorname
  * Linkname
  * Notes
  * Downloadlimit

If desired the file can be downloaded

____
_Disclaimer: This is an unofficial project and is not supported by DRACOON. Use is at your own risk and without any guarantee_  
Only the public API commands are used, therefore no authentication is required and is also not used in the examples.
