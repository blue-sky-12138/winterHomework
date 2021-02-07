package utilities

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
	"time"
)

//MD5即时加密的快捷方式
func CryptographyNow(Data string) (string,int64) {
	Md5salt:=time.Now().Unix()
	return Cryptography(Data,Md5salt),Md5salt
}

//MD5加密
func Cryptography(Data string,Md5salt int64)string{
	has:=md5.New()
	io.WriteString(has,Data+strconv.FormatInt(Md5salt,10))
	tem:= has.Sum(nil)
	Result:=hex.EncodeToString(tem)
	return Result
}