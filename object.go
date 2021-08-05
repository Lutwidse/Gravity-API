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

type DeviceInfo struct {
	Country, Product, Sys_lang, Uwd, App_version, Sign, Pkg, Referrer, Sub_referrer, System_version, Model, Device, Brand, Push_token, Address string
	Zone, Sdk_version                                                                                                                          int
	Ts                                                                                                                                         uint32
}

func (di *DeviceInfo) SetSignWithTimestamp() {
	di.Ts = getTimestamp()
	di.Sign = getSign(di.Ts)
}

func (di *DeviceInfo) SetAddress(address string) {
	aes := NewAESEncrypter()

	_, ph := aes.GetKeys()

	encryptedData := aes.AESEncrypt(address, []byte(ph))
	di.Address = base64.StdEncoding.EncodeToString(encryptedData)
}

func (di *DeviceInfo) SetUWD() {
	// Should be stored and static after UUID generated.
	userId := uuid.New().String()
	userId = strings.ToUpper(userId)

	aes := NewAESEncrypter()
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
	di.Uwd = base64.StdEncoding.EncodeToString(encryptedData)
}