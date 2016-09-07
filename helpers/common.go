//helpers包用来写一些通用辅助函数
package helpers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func MyHttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		ret := map[string]string{"code": "500", "msg": "get error occur"}
		ret_json, _ := json.Marshal(ret)
		return string(ret_json)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ret := map[string]string{"code": "500", "msg": "body error occur"}
		ret_json, _ := json.Marshal(ret)
		return string(ret_json)
	}

	return string(body)
}

func MyHttpPost(url string, param interface{}) string {
	query := ""
	switch param.(type) {
	case string:
		query = param.(string)
	default:
		query = GetUrlQuery(param.(map[string]string))
	}

	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(query))
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body error occur")
	}

	fmt.Println(string(body))
	return string(body)
}

func GetUrlQuery(params map[string]string) string {
	query := ""
	for key, val := range params {
		query += key + "=" + val + "&"
	}

	return query[0 : len(query)-1]

}

func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func SiteUrl(baseName string) string {
	return beego.AppConfig.String("siteUrl") + "/" + baseName
}
