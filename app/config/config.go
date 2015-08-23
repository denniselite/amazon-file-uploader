package config

import "../app"

func Set(p *app.Client) {
	p.Login = "LOGIN"
	p.Pass = "PASS"
	p.BucketName = "BucketName"
	p.Permissions = "public-read"
	p.RegionName = "ServerRegion"
}
