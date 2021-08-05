package main

import (
	"fmt"
	"github.com/Lutwidse/gravity-api"
	"github.com/go-resty/resty/v2"
)

func main() {

	client := gravity_api.NewGravityClient(resty.New())
	var resp interface{}

	deviceInfo := gravity_api.DeviceInfo{
		Country:        gravity_api.Country,
		Product:        gravity_api.Product,
		Sys_lang:       gravity_api.Sys_lang,
		Uwd:            "",
		App_version:    gravity_api.App_version,
		Sign:           "",
		Pkg:            gravity_api.Pkg,
		Referrer:       gravity_api.Referrer,
		Sub_referrer:   "",
		Zone:           gravity_api.Zone,
		System_version: gravity_api.System_version,
		Sdk_version:    gravity_api.Sdk_version,
		Model:          gravity_api.Model,
		Device:         gravity_api.Device,
		Brand:          gravity_api.Brand,
		// Firebase
		Push_token: gravity_api.Push_token,
	}

	// legitimate communication
	resp = client.User.UpdateDevice(&deviceInfo)
	fmt.Printf("updatedeviceInfo: %v\n", resp)

	resp = client.Common.Getimpornword(&deviceInfo)
	fmt.Printf("getimpornword: %v\n", resp)

	resp = client.Push.BindToken(&deviceInfo)
	fmt.Printf("bindToken: %v\n", resp)

	resp = client.Push.BindToken(&deviceInfo)
	fmt.Printf("bindToken: %v\n", resp)
}
