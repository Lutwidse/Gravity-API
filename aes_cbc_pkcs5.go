package gravity_api

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// Author: https://gist.github.com/hothero/7d085573f5cb7cdb5801d7adcf66dcf3

type aesEncrypter struct {
	initialVector string
	passphrase    string
}

func NewAESEncrypter() *aesEncrypter {
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
	return &aesEncrypter{initialVector: "qrstuvwxyz123456", passphrase: "baisimeji9262019"}
}

func (x *aesEncrypter) GetKeys() (string, string) {
	return x.initialVector, x.passphrase
}

func (x aesEncrypter) AESEncrypt(src string, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(x.initialVector))
	content := []byte(src)
	content = x.PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

func (x aesEncrypter) AESDecrypt(crypt []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, []byte(x.initialVector))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return x.PKCS5Trimming(decrypted)
}

func (x aesEncrypter) PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (aesEncrypter) PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
