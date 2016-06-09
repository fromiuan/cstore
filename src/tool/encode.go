package tool

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

/**
*把字节转化为MD5字符串
 */
func MD5(b []byte) string {
	vCrypto := md5.New()
	vCrypto.Write(b)
	return hex.EncodeToString(vCrypto.Sum(nil))
}

/*
*加密用户密码
 */
func EncodeUserPwd(uname, pwd string) string {
	// return MD5([]byte(strings.Join([]string{uname, "$user$", pwd}, "")))
	return MD5([]byte(pwd))
}

/*
*加密客户密码
 */
func EncodeMemberPwd(uname, pwd string) string {
	return MD5([]byte(strings.Join([]string{uname, "$member$", pwd}, "")))
}

/*
*加密商户密码
 */
func EncodeStorerPwd(uname, pwd string) string {
	return MD5([]byte(strings.Join([]string{uname, "$storer$", pwd}, "")))
}
