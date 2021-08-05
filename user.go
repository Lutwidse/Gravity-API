package gravity_api

type IUser interface {
	UpdateDevice(country string, product string, sys_lang string, uwd string, app_version string, sign string, pkg string, referrer string, zone int, system_version string, skd_version int, model string, device string, brand string, sub_referrer string, ts uint32) updateDeviceResponse
}

type User struct {
	client *GravityClient
}

type updateDeviceResponse struct {
	Data struct {
	} `json:"data"`
	Errmsg string `json:"errmsg"`
	Errno  int    `json:"errno"`
}

func (p *User) UpdateDevice(deviceInfo *DeviceInfo) interface{} {
	deviceInfo.SetSignWithTimestamp()
	resp, err := p.client.httpClient.R().SetBody(map[string]interface{}{
		"country":        deviceInfo.Country,
		"product":        deviceInfo.Product,
		"sys_lang":       deviceInfo.Sys_lang,
		"uwd":            deviceInfo.Uwd,
		"app_version":    deviceInfo.App_version,
		"sign":           deviceInfo.Sign,
		"pkg":            deviceInfo.Pkg,
		"referrer":       deviceInfo.Referrer,
		"zone":           deviceInfo.Zone,
		"system_version": deviceInfo.System_version,
		"skd_version":    deviceInfo.Sdk_version,
		"model":          deviceInfo.Model,
		"device":         deviceInfo.Device,
		"brand":          deviceInfo.Brand,
		"sub_referrer":   deviceInfo.Sub_referrer,
		"ts":             deviceInfo.Ts,
	}).SetResult(&updateDeviceResponse{}).Post(userUrl + "/updateDevice")
	if err != nil {
		panic(err)
	}

	return resp
}
