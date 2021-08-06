package gravity_api

import (
	"strconv"
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
	resp, err := p.client.httpClient.R().SetQueryParams(map[string]string{
		"country":        p.client.DeviceInfo.Country,
		"product":        p.client.DeviceInfo.Product,
		"sys_lang":       p.client.DeviceInfo.Sys_lang,
		"uwd":            p.client.DeviceInfo.Uwd,
		"app_version":    p.client.DeviceInfo.App_version,
		"sign":           p.client.DeviceInfo.Sign,
		"pkg":            p.client.DeviceInfo.Pkg,
		"referrer":       p.client.DeviceInfo.Referrer,
		"zone":           strconv.Itoa(p.client.DeviceInfo.Zone),
		"system_version": p.client.DeviceInfo.System_version,
		"skd_version":    strconv.Itoa(p.client.DeviceInfo.Sdk_version),
		"model":          p.client.DeviceInfo.Model,
		"device":         p.client.DeviceInfo.Device,
		"brand":          p.client.DeviceInfo.Brand,
		"ts":             strconv.FormatUint(uint64(p.client.DeviceInfo.Ts), 10),
	}).SetResult(&getimpornwordResponse{}).Get(commonUrl + "/getimpornword")
	if err != nil {
		panic(err)
	}

	return resp
}
