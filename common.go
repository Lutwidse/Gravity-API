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

func (p *Common) Getimpornword(deviceInfo *DeviceInfo) interface{} {
	deviceInfo.SetSignWithTimestamp()
	resp, err := p.client.httpClient.R().SetQueryParams(map[string]string{
		"country":        deviceInfo.Country,
		"product":        deviceInfo.Product,
		"sys_lang":       deviceInfo.Sys_lang,
		"uwd":            deviceInfo.Uwd,
		"app_version":    deviceInfo.App_version,
		"sign":           deviceInfo.Sign,
		"pkg":            deviceInfo.Pkg,
		"referrer":       deviceInfo.Referrer,
		"zone":           strconv.Itoa(deviceInfo.Zone),
		"system_version": deviceInfo.System_version,
		"skd_version":    strconv.Itoa(deviceInfo.Sdk_version),
		"model":          deviceInfo.Model,
		"device":         deviceInfo.Device,
		"brand":          deviceInfo.Brand,
		"ts":             strconv.FormatUint(uint64(deviceInfo.Ts), 10),
	}).SetResult(&getimpornwordResponse{}).Get(commonUrl + "/getimpornword")
	if err != nil {
		panic(err)
	}

	return resp
}
