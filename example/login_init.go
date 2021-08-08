package main

import (
	"fmt"
	"github.com/Lutwidse/gravity-api"
)

func main() {

	client := gravity_api.NewGravityClient()
	var resp interface{}

	// legitimate communication
	resp = client.User.UpdateDevice()
	fmt.Printf("updateDevice: %v\n", resp)

	resp = client.Common.Getimpornword()
	fmt.Printf("getimpornword: %v\n", resp)

	resp = client.Push.BindToken()
	fmt.Printf("bindToken: %v\n", resp)

	resp = client.Push.BindToken()
	fmt.Printf("bindToken: %v\n", resp)
}
