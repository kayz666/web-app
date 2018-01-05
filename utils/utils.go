package utils

import (
	"net/url"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
)

//发送post请求
//@param apiUrl api地址
//@param postParam post参数
//@param result map格式json数据， err error对象
func UrlPost(apiUrl string,postParam map[string]string)   (result map[string]interface{}, err error) {
	postValue := url.Values{}
	for key, value := range postParam{
		postValue.Set(key, value)
	}
	fmt.Println("<POST>" + apiUrl)
	fmt.Println("post param : " + postValue.Encode())
	response, err := http.Post(apiUrl, "application/x-www-form-urlencoded", strings.NewReader(postValue.Encode()))
	obj := make(map[string]interface{})
	if err != nil{
		return nil, err
	}
	text, err2 := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err2 != nil{
		return nil, err2
	}
	err3 :=  json.Unmarshal(text, &obj)
	return obj, err3
}
