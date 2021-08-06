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
	_, ph := aes.GetKeys()

	encryptedData := aes.AESEncrypt(userId, []byte(ph))
	di.Uwd = base64.StdEncoding.EncodeToString(encryptedData)
}