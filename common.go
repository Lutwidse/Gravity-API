package gravity_api

import (
	"strconv"
)

type ICommon interface {
	Getimpornword(country string, product string, sys_lang string, uwd string, app_version string, sign string, pkg string, referrer string, zone int, system_version string, skd_version int, model string, device string, brand string, sub_referrer string, ts uint32) getimpornwordResponse
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

func (p *Common) Getimpornword(country string, product string, sys_lang string, uwd string, app_version string, sign string, pkg string, referrer string, zone int, system_version string, skd_version int, model string, device string, brand string, sub_referrer string, ts uint32) interface{} {
	resp, err := p.client.httpClient.R().SetQueryParams(map[string]string{
		"country":        country,
		"product":        product,
		"sys_lang":       sys_lang,
		"uwd":            uwd,
		"app_version":    app_version,
		"sign":           sign,
		"pkg":            pkg,
		"referrer":       referrer,
		"zone":           strconv.Itoa(zone),
		"system_version": system_version,
		"skd_version":    strconv.Itoa(skd_version),
		"model":          model,
		"device":         device,
		"brand":          brand,
		"sub_referrer":   sub_referrer,
		"ts":             strconv.FormatUint(uint64(ts), 10),
	}).SetResult(&getimpornwordResponse{}).Post(userUrl + "/getimpornword")
	if err != nil {
		panic(err)
	}

	return resp.Result()
}
