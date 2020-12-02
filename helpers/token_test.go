package helpers

import (
	"testing"
)

func TestGetSign(t *testing.T) {
	params := map[string]string{
		"app_key":   "XdfetWtesLs1",
		"user_name": "zhangsan@tt.com",
		"url":       "http://www.google.com",
		"time":      "2015-03-05",
	}
	secret := "XzfeGWQ2df"
	expectedSign := "d394697fcfe9ee02cc58220f17858695"
	if expectedSign != GetSign(params, secret) {
		t.Error("not equal")
	}

}

func TestCheckSign(t *testing.T) {
	params := map[string]string{
		"app_key":   "XdfetWtesLs1",
		"user_name": "zhangsan@tt.com",
		"url":       "http://www.google.com",
		"time":      "2015-03-05",
	}
	secret := "XzfeGWQ2df"
	resSign := "d394697fcfe9ee02cc58220f17858695"
	if !CheckSign(params, secret, resSign) {
		t.Error("not equal")
	}
}
