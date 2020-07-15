package encrypt

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

/**
 * 对一个字符串进行MD5加密,不可解密
 * salt: 散列值
 */
func Md5StringWithSalt(s, salt string) string {
	h := md5.New()
	h.Write([]byte(s + salt))
	return hex.EncodeToString(h.Sum(nil))
}

/**
 * 对一个字符串进行MD5加密,不可解密
 */
func Md5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}


/**
 * 获取一个Guid值
 */
func GoUUID(len int) (string,error) {
	rawByte := make([]byte, len)
	if _, err := io.ReadFull(rand.Reader, rawByte); err != nil {
		return "", errors.New("create uuid failed")
	}
	saltByte := make([]byte, 6)
	if _, err := io.ReadFull(rand.Reader, saltByte); err != nil {
		saltByte = []byte("GoUUID")
	}
	md5String := Md5StringWithSalt(base64.RawStdEncoding.EncodeToString(rawByte), base64.RawStdEncoding.EncodeToString(saltByte))
	return md5String[:len], nil
}