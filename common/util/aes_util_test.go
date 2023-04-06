package util

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"testing"
)

func TestCBCEncrypt(t *testing.T) {
	origData := []byte("hello world") // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOP") // 加密的密钥
	log.Println("原文：", string(origData))

	encrypted := AesEncryptCBC(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AesDecryptCBC(encrypted, key)
	log.Println("解密结果：", string(decrypted))
}
