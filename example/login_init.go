package main

import (
	"encoding/base64"
	"fmt"
	"strings"

	"crypto/md5"
	"encoding/hex"
	"github.com/Lutwidse/gravity-api"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"strconv"
	"time"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetTimestamp() uint32 {
	return uint32(time.Now().Unix())
}

func GetSign(ts uint32) string {
	return GetMD5Hash(strconv.FormatUint(uint64(ts), 10))
}

type DeviceInfo struct {
	country, product, sys_lang, uwd, app_version, sign, pkg, referrer, sub_referrer string
	zone                                                                            int
	system_version                                                                  string
	sdk_version                                                                     int
	model, device, brand, push_token                                                string
	ts                                                                              uint32
}

func (di *DeviceInfo) SetSignWithTimestamp() {
	di.ts = GetTimestamp()
	di.sign = GetSign(di.ts)
}

func (di *DeviceInfo) SetUWD() {
	// Should be stored and static after UUID generated.
	userId := uuid.New().String()
	userId = strings.ToUpper(userId)

	aes := gravity_api.NewAESEncrypter()

	aes.SetKeys("qrstuvwxyz123456", "baisimeji9262019")
	/*
			public static final String a(String p0){
		       i.IllegalArgumentHandling(p0, "content");
		       i.IllegalArgumentHandling(p0, "encData");
		       String str = "qrstuvwxyz123456";
		       i.IllegalArgumentHandling(str, "vector");
		       Cipher cInstance = Cipher.getInstance("AES/CBC/PKCS5Padding");
		       byte[] bBytes = "baisimeji9262019".getBytes(a.a);
		       String str1 = "\(this as java.lang.String\).getBytes\(charset\)";
		       i.NullHandling(bBytes, str1);
		       byte[] bBytes1 = str.getBytes(a.a);
		       i.NullHandling(bBytes1, str1);
		       cInstance.init(1, new SecretKeySpec(bBytes, "AES"), new IvParameterSpec(bBytes1));
		       Charset cName = Charset.forName("utf-8");
		       i.NullHandling(cName, "Charset.forName\(charsetName\)");
		       byte[] bBytes2 = p0.getBytes(cName);
		       i.NullHandling(bBytes2, str1);
		       String sstr = Base64.encodeToString(cInstance.doFinal(bBytes2), 0);
		       i.IllegalStateHandle(sstr, "Base64.encodeToString\(encrypted, 0\)");
		       return sstr;
		    }
	*/
	_, ph := aes.GetKeys()

	encryptedData := aes.AESEncrypt(userId, []byte(ph))
	di.uwd = base64.StdEncoding.EncodeToString(encryptedData)
}

func main() {

	di := &DeviceInfo{
		country:        "JP",
		product:        "gravity",
		sys_lang:       "ja-JP",
		uwd:            "",
		app_version:    "2.8.0",
		sign:           "",
		pkg:            "anonymous.sns.community.gravity",
		referrer:       "default",
		sub_referrer:   "",
		zone:           0,
		system_version: "8.0.0",
		sdk_version:    26,
		model:          "Google_Nexus_6",
		device:         "android",
		brand:          "unknown",
		// Firebase
		push_token: "-1",
	}

	// legitimate communication
	client := gravity_api.NewGravityClient(resty.New())
	var resp interface{}

	// I'll look up other API functions when I'm free and refactor this dirty code.
	di.SetSignWithTimestamp()
	resp = client.User.UpdateDevice(di.country, di.product, di.sys_lang, di.uwd, di.app_version, di.sign, di.pkg, di.referrer, di.zone, di.system_version, di.sdk_version, di.model, di.device, di.brand, di.sub_referrer, di.ts)
	fmt.Printf("updateDevice: %v\n", resp)

	di.SetSignWithTimestamp()
	resp = client.Common.Getimpornword(di.country, di.product, di.sys_lang, di.uwd, di.app_version, di.sign, di.pkg, di.referrer, di.zone, di.system_version, di.sdk_version, di.model, di.device, di.brand, di.ts)
	fmt.Printf("getimpornword: %v\n", resp)

	di.SetSignWithTimestamp()
	resp = client.Push.BindToken(di.country, di.product, di.sys_lang, di.uwd, di.app_version, di.sign, di.pkg, di.referrer, di.zone, di.system_version, di.sdk_version, di.model, di.device, di.brand, di.push_token, di.ts)
	fmt.Printf("bindToken: %v\n", resp)

	di.SetSignWithTimestamp()
	resp = client.Push.BindToken(di.country, di.product, di.sys_lang, di.uwd, di.app_version, di.sign, di.pkg, di.referrer, di.zone, di.system_version, di.sdk_version, di.model, di.device, di.brand, di.push_token, di.ts)
	fmt.Printf("bindToken: %v\n", resp)
}
