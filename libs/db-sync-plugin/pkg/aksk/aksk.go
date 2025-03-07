package aksk

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"os"
)

func GetAkSK(skPart string) (string, string) {
	//skPart := "<TODO:从应用开发平台获取应用秘钥>"
	appId := os.Getenv("AK")

	var appKey string
	if os.Getenv("SK_SALT") == "" {
		appKey = os.Getenv("SK")
	} else {
		by, err := base64.StdEncoding.DecodeString(os.Getenv("SK_SALT"))
		if err != nil {
			return "", ""
		}
		key := GetContentHash(skPart)[:16]
		iv := GetContentHash(appId)[:16]
		appKeyByte, err := CbcDecrypt(by, key, iv)
		if err != nil {
			return "", ""
		}
		appKey = string(appKeyByte)
	}
	return appId, appKey
}

func GetContentHash(content string) string {
	hash := sha256.New()
	hash.Write([]byte(content))
	sum := hash.Sum(nil)
	hashStr := hex.EncodeToString(sum)
	return hashStr
}

func CbcDecrypt(data []byte, key, iv string) ([]byte, error) {
	bKey := []byte(key)
	bIv := []byte(iv)
	block, err := aes.NewCipher(bKey)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, bIv)
	dData := make([]byte, len(data))
	mode.CryptBlocks(dData, data)
	length := len(dData)
	unPadding := int(dData[length-1])
	if length-unPadding < 0 {
		return dData, nil
	}
	return dData[:(length - unPadding)], nil
}
