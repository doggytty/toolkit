package encrypt

import "encoding/base64"

/**
 * base64加密字符串
 */
func Base64Encode(str string) string {
	var src []byte = []byte(str)
	return base64.RawStdEncoding.EncodeToString(src)
}

/**
 * base64解密字符串
 */
func Base64Decode(str string) (string, error) {
	//var coder = base64.NewEncoding(BASE64TABLE)
	var src []byte = []byte(str)
	by, err := base64.RawStdEncoding.DecodeString(string(src))
	return string(by), err
}

/**
 * base64加密url
 */
func Base64EncodeURL(urlString string) string {
	var src []byte = []byte(urlString)
	return base64.RawURLEncoding.EncodeToString(src)
}

/**
 * base64解密url
 */
func Base64DecodeURL(urlString string) (string, error) {
	var src []byte = []byte(urlString)
	by, err := base64.RawURLEncoding.DecodeString(string(src))
	return string(by), err
}
