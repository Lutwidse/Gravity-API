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

func (p *User) UpdateDevice(country string, product string, sys_lang string, uwd string, app_version string, sign string, pkg string, referrer string, zone int, system_version string, skd_version int, model string, device string, brand string, sub_referrer string, ts uint32) interface{} {
	resp, err := p.client.httpClient.R().SetBody(map[string]interface{}{
		"country":        country,
		"product":        product,
		"sys_lang":       sys_lang,
		"uwd":            uwd,
		"app_version":    app_version,
		"sign":           sign,
		"pkg":            pkg,
		"referrer":       referrer,
		"zone":           zone,
		"system_version": system_version,
		"skd_version":    skd_version,
		"model":          model,
		"device":         device,
		"brand":          brand,
		"sub_referrer":   sub_referrer,
		"ts":             ts,
	}).SetResult(&updateDeviceResponse{}).Post(userUrl + "/updateDevice")
	if err != nil {
		panic(err)
	}

	return resp
}
