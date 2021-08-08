package gravity_api

import (
	"github.com/go-resty/resty/v2"
)

type GravityClient struct {
	userClient   *resty.Client
	commonClient *resty.Client
	pushClient   *resty.Client

	DeviceInfo DeviceInfo

	User   *User
	Common *Common
	Push   *Push
}

func NewGravityClient() *GravityClient {

	var uc, cc, pc *resty.Client = resty.New(), resty.New(), resty.New()
	/*
		TODO: move SetTimestampWithSign() onto the middleware OnBeforeRequest(...) if other API is also restful and it's capable.
		since SetQuery or SetPath will overwrite old params.
	*/
	uc.SetHostURL(userUrl)
	cc.SetHostURL(commonUrl)
	pc.SetHostURL(pushUrl)

	deviceInfo := DeviceInfo{}
	// 
	if (Uwd) == "" {
		deviceInfo.SetUWD()
	}
	deviceInfo.initDefaultBodyParam()

	c := &GravityClient{userClient: uc, commonClient: cc, pushClient: pc}
	c.DeviceInfo = deviceInfo
	c.User = &User{client: c}
	c.Common = &Common{client: c}
	c.Push = &Push{client: c}

	return c
}
