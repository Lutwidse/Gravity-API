package gravity_api

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

/*
type DeviceInfo struct {
	country, product, sys_lang, uwd, app_version, sign, pkg, referrer, sub_referrer, system_version, model, device, brand, push_token, address string
	zone, sdk_version                                                                                                                          int
	ts                                                                                                                                         uint32
}
*/

type DeviceInfo map[string]string

func (di DeviceInfo) initDefaultBodyParam() {
	di["country"] = Country
	di["product"] = Product
	di["sys_lang"] = Sys_lang
	di["uwd"] = Uwd
	di["app_version"] = App_version
	di["sign"] = Sign
	di["pkg"] = Pkg
	di["referrer"] = Referrer
	di["zone"] = strconv.Itoa(Zone)
	di["system_version"] = System_version
	di["sdk_version"] = strconv.Itoa(Sdk_version)
	di["mode"] = Model
	di["device"] = Device
	di["brand"] = Brand
	di["sub_referrer"] = Sub_referrer
	di["ts"] = strconv.FormatUint(uint64(Ts), 10)
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func getTimestamp() uint32 {
	return uint32(time.Now().Unix())
}

func getSign(ts uint32) string {
	return getMD5Hash(strconv.FormatUint(uint64(ts), 10))
}

func (di DeviceInfo) SetSignWithTimestamp() {
	ts := getTimestamp()
	di["ts"] = strconv.FormatUint(uint64(getTimestamp()), 10)
	di["sign"] = getSign(ts)
}

func (di DeviceInfo) SetAddress(address string) {
	aes := NewAESEncrypter()

	_, ph := aes.GetKeys()

	encryptedData := aes.AESEncrypt(address, []byte(ph))
	di["address"] = base64.StdEncoding.EncodeToString(encryptedData)
}

func (di DeviceInfo) SetUWD() {
	// Should be stored and static after UUID generated.
	userId := uuid.New().String()
	userId = strings.ToUpper(userId)

	aes := NewAESEncrypter()
	_, ph := aes.GetKeys()

	encryptedData := aes.AESEncrypt(userId, []byte(ph))
	di["uwd"] = base64.StdEncoding.EncodeToString(encryptedData)
}
