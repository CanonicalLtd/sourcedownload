# Source Download Tool

Downloads the sources of an Ubuntu Core snap from the OSS Compliance service
for Ubuntu Core.

```
sourcedownload -help
  -path string
        Location to store the download files (default "downloads")
  -revision int
        The revision of the snap to process
  -snap string
        The name of the snap to process
  -url string
        The URL of the Compliance Service (default "https://sources.iotdevice.io")
```
If the application is called with no parameters, it will list the available
snap revisions for each architecture.


## Build
The application has been developed using Go v13.* and uses Go Modules for dependencies.
```
go build sourcedownload.go
```
