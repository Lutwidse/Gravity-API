package gravity_api

import "github.com/go-resty/resty/v2"

type GravityClient struct {
	httpClient *resty.Client

	User   *User
	Common *Common
	Push *Push
}

func NewGravityClient(httpClient *resty.Client) *GravityClient {
	c := &GravityClient{httpClient: httpClient}
	c.User = &User{client: c}
	c.Common = &Common{client: c}
	c.Push = &Push{client: c}
	return c
}
