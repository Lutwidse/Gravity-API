package gravity_api

import (
)

type ICommon interface {
	Getimpornword(country string, product string, sys_lang string, uwd string, app_version string, sign string, pkg string, referrer string, zone int, system_version string, skd_version int, model string, device string, brand string, ts uint32) getimpornwordResponse
}

type Common struct {
	client *GravityClient
}

type getimpornwordResponse struct {
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
	Errmsg string `json:"errmsg"`
	Errno  int    `json:"errno"`
}

func (p *Common) Getimpornword() interface{} {
	p.client.DeviceInfo.SetSignWithTimestamp()
	params := p.client.DeviceInfo
	resp, err := p.client.commonClient.R().SetPathParams(params).SetResult(&getimpornwordResponse{}).
		Get("/getimpornword?country={country}&product={product}&sys_lang={sys_lang}&uwd={uwd}&app_version={app_version}&sign={sign}&pkg={pkg}&referrer={referrer}&zone={zone}&system_version={system_version}&sdk_version={sdk_version}&model={model}&device={device}&brand={brand}&ts={ts}")
	if err != nil {
		panic(err)
	}

	return resp
}
