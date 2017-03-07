package helpers

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func GetSign(params map[string]string, secret string) string {
	strs := make([]string, 0)
	for key, _ := range params {
		strs = append(strs, key)
	}
	sort.Strings(strs)

	sign := ""
	for _, key := range strs {
		sign += key + "=" + params[key]
	}
	sign += secret
	return Md5(sign)
}

func CheckSign(params map[string]string, secret, resSign string) bool {
	expectedSign := GetSign(params, secret)
	fmt.Println("expectedSign:" + expectedSign)
	if resSign == expectedSign {
		return true
	}
	return false
}

func GenerateAppKey() string {
	return CreateNonceStr(6) + CreateSecond()
}

func GenerateAppSecret(appKey string) string {
	unixTime := time.Now().Unix()
	return Sha1(appKey + string(unixTime))
}

func CreateSecond() string {
	return fmt.Sprintf("%.2d", time.Now().Second())
}

func CreateNonceStr(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	var randStr string
	for i := 1; i <= length; i++ {
		index := i
		if i >= len(chars) {
			index = 1
		}

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randNum := r.Intn(len(chars) - index)
		randStr += string(chars[randNum])
	}

	return string(randStr)
}
