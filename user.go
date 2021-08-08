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

func (p *User) UpdateDevice() interface{} {
	p.client.DeviceInfo.SetSignWithTimestamp()
	params := p.client.DeviceInfo
	resp, err := p.client.userClient.R().SetBody(params).SetResult(&updateDeviceResponse{}).
	Post("/updateDevice")
	if err != nil {
		panic(err)
	}

	return resp
}
