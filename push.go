package gravity_api

type IPush interface {
	BindToken(country string, product string, sys_lang string, uwd string, app_version string, sign string, pkg string, referrer string, zone int, system_version string, skd_version int, model string, device string, brand string, sub_referrer string, ts uint32) bindTokenResponse
}

type Push struct {
	client *GravityClient
}

type bindTokenResponse struct {
	Data struct {
	} `json:"data"`
	Errmsg string `json:"errmsg"`
	Errno  int    `json:"errno"`
}

func (p *Push) BindToken() interface{} {
	p.client.DeviceInfo.SetSignWithTimestamp()
	resp, err := p.client.httpClient.R().SetBody(map[string]interface{}{
		"country":        p.client.DeviceInfo.Country,
		"product":        p.client.DeviceInfo.Product,
		"sys_lang":       p.client.DeviceInfo.Sys_lang,
		"uwd":            p.client.DeviceInfo.Uwd,
		"app_version":    p.client.DeviceInfo.App_version,
		"sign":           p.client.DeviceInfo.Sign,
		"pkg":            p.client.DeviceInfo.Pkg,
		"referrer":       p.client.DeviceInfo.Referrer,
		"zone":           p.client.DeviceInfo.Zone,
		"system_version": p.client.DeviceInfo.System_version,
		"skd_version":    p.client.DeviceInfo.Sdk_version,
		"model":          p.client.DeviceInfo.Model,
		"device":         p.client.DeviceInfo.Device,
		"brand":          p.client.DeviceInfo.Brand,
		"sub_referrer":   p.client.DeviceInfo.Sub_referrer,
		"ts":             p.client.DeviceInfo.Ts,
	}).SetResult(&bindTokenResponse{}).Post(pushUrl + "/bindToken")
	if err != nil {
		panic(err)
	}

	return resp
}
