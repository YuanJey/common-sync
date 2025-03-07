package enc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func GenerateKey(uid, salt string) ([]byte, error) {
	// 计算 SHA-256 哈希
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%s%s", uid, salt)))
	hashBytes := hash.Sum(nil)
	// 转换为十六进制字符串
	hashHex := hex.EncodeToString(hashBytes)
	// 取前 32 个字符作为密钥
	return []byte(hashHex[:32]), nil

}
func Encrypt(plaintext []byte, key []byte) (string, error) {
	// 创建 AES GCM 加密模式
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
