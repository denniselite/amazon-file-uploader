package models

import (
	"../config"
	aws "github.com/ashleydw/goamz/aws"
	s3 "github.com/ashleydw/goamz/s3"
	"os"
)

type Client struct {
	Login       string
	Pass        string
	BucketName  string
	RegionName  string
	Permissions string
	connection  *s3.S3
	Bucket      *s3.Bucket
	region      aws.Region
}

func (c *Client) Init() {
	var C Client
	config.Set(&C)
	return C
}

func (c *Client) Connect() {
	AWSAuth := aws.Auth{
		AccessKey: c.Login,
		SecretKey: c.Pass,
	}
	c.region = aws.EUWest
	c.connection = s3.New(AWSAuth, a.region)
	c.region.S3BucketEndpoint = c.RegionName
	c.Bucket = c.connection.Bucket(c.BucketName)
}

func (c *Client) SendFile(backupFile *file) (error, uint8, string) {
	err := c.Bucket.Put(
		backupFile.Name,
		backupFile.Bytes,
		backupFile.Type,
		s3.ACL(c.Permissions))
	if err != nil {
		return err, 8, backupFile.Name
	} else {
		return nil, 0, ""
	}

}
