package lib

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)
//////////////////////// AesEncrypt AES加密  CBC模式自设定秘钥,需要设置偏移量/////////////////////////////

func getAesKeys() string {
	str := ""
	return str
}

func AesEncrypt(encodeStr string, key []byte, iv []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	//根据key 生成密文
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(encrypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// PKCS7Padding PKCS#7填充，Buf需要被填充为K的整数倍，
// 在buf的尾部填充(K-N%K)个字节，每个字节的内容是(K- N%K)
func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	//填充
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padText...)
}

func AesDecrypt(decodeStr string, key []byte, iv []byte) ([]byte, error) {
	//先解密base64
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, fmt.Errorf("base64.StdEncoding.DecodeString err %v", err)
	}
	if len(decodeBytes) < len(key) {
		return nil, fmt.Errorf("the length of encrypted message too short: %d", len(decodeBytes))
	}
	if len(decodeBytes)&(len(key)-1) != 0 { // or len(enc)%len(key) != 0
		return nil, fmt.Errorf("encrypted message is not a multiple of the key size(%d), the length is %d", len(key), len(decodeBytes))
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("aes.NewCipher err %v", err)
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData, err = PKCS5UnPadding(origData)
	return origData, err
}
// PKCS7UnPadding 去掉PKCS#7填充，Buf需要被填充为K的整数倍，
// 在buf的尾部填充(K-N%K)个字节，每个字节的内容是(K- N%K)
func PKCS5UnPadding(origData []byte) (tmp []byte, err error) {
	length := len(origData)
	unPadding := int(origData[length-1])
	if length < unPadding {
		err = fmt.Errorf("PKCS5UnPadding out of range ")
		return
	}
	tmp = origData[:(length - unPadding)]
	return
}

func MakeKeyAndIV() (string, string) {
	chars := "abcdefghijkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ23456789"
	charLen := float64(len(chars))
	key := ""
	iv := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 32; i++ {
		rfi := int(charLen * rand.Float64())
		key += fmt.Sprintf("%c", chars[rfi])
	}
	for i := 0; i < 16; i++ {
		rfi := int(charLen * rand.Float64())
		iv += fmt.Sprintf("%c", chars[rfi])
	}
	return key, iv
}

//解密
func DecodePassword(passwordEncrypt string, saltEncrypt string) (password string, err error) {
	if passwordEncrypt == "" || saltEncrypt == "" {
		return "", fmt.Errorf("password_encrypt or salt_encrypt is nil")
	}
	keys := getAesKeys()
	passwordSalt, err := DecodeSalt(saltEncrypt, keys)
	if err != nil {
		return "", fmt.Errorf("DecodeSalt wrong. %v", err)
	}
	b := []byte(passwordSalt)
	if len(b) != 48 {
		return "", fmt.Errorf("the password salt len is not 48")
	}
	passByte, err := AesDecrypt(passwordEncrypt, b[0:32], b[32:])
	if err != nil {
		return "", fmt.Errorf("AesDecrypt wrong. %v", err)
	}
	return string(passByte), nil
}

//加密
func EncodePassword(password string) (passwordEncrypt string, saltEncrypt string, err error) {
	key, iv := MakeKeyAndIV()
	passwordEncrypt, err = AesEncrypt(password, []byte(key), []byte(iv))
	if err != nil {
		return "", "", err
	}
	passwordSalt := key + iv
	//对keysalt进行一次ase256加密
	keys := getAesKeys()
	saltEncrypt, err = EncodeSalt(passwordSalt, keys)
	if err != nil {
		return "", "", err
	}

	return passwordEncrypt, saltEncrypt, nil
}
func EncodeSalt(salt string, keyVi string) (saltEncrypt string, err error) {
	if len(keyVi) != 48 {
		return "", fmt.Errorf("the password salt len is not 48")
	}
	b := []byte(keyVi)
	saltEncrypt, err = AesEncrypt(salt, b[0:32], b[32:])
	if err != nil {
		return "", err
	}
	return saltEncrypt, nil

}
func DecodeSalt(saltEncrypt string, keyVi string) (passwordSalt string, err error) {
	b := []byte(keyVi)
	if len(b) != 48 {
		return "", fmt.Errorf("the key vi len is not 48")
	}
	passwordSaltByte, err := AesDecrypt(saltEncrypt, b[0:32], b[32:])
	if err != nil {
		return "", err
	}
	return string(passwordSaltByte), nil
}
