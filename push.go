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
	params := p.client.DeviceInfo
	params["push_token"] = Push_token
	resp, err := p.client.pushClient.R().SetBody(params).SetResult(&bindTokenResponse{}).
	Post("/bindToken")
	if err != nil {
		panic(err)
	}

	return resp
}
