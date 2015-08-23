# Amazon File Uploader (AWS S3)#

Uploading some files (backups, etc) into Amazon s3 cloud

## Compile & run ##

`go build amazon.go` - Creating new build

`./amazon directory/file` - Starting upload to amazon cloud. On remote server will be created current path with new file.

**If you restart the loader with the same file then the file will be uploaded to the Amazon and overwritten by the new version**

## Configuration ##

`app/config/config.go` - configuration file for package.
