package gravity_api

import "github.com/go-resty/resty/v2"

type DeviceInfo struct {
	Country, Product, Sys_lang, Uwd, App_version, Sign, Pkg, Referrer, Sub_referrer, System_version, Model, Device, Brand, Push_token, Address string
	Zone, Sdk_version                                                                                                                          int
	Ts                                                                                                                                         uint32
}

type GravityClient struct {
	httpClient *resty.Client

	DeviceInfo *DeviceInfo

	User   *User
	Common *Common
	Push   *Push
}

func NewGravityClient(httpClient *resty.Client) *GravityClient {
	c := &GravityClient{httpClient: httpClient, DeviceInfo: &DeviceInfo{}}
	c.User = &User{client: c}
	c.Common = &Common{client: c}
	c.Push = &Push{client: c}
	return c
}
