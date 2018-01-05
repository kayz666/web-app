package utils

import (
	"net/http"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
	"io/ioutil"
	"encoding/json"
	"errors"
)

type HTTP_Body struct {
	Success	bool		`json:"success"`
	Error	string		`json:"error"`
	Return	HTTP_Return		`json:"return"`
}
type HTTP_Return struct {
	Code	string		`json:"code"`
	Link	string		`json:"link"`
}

func HTTP_ConfirmWithEmail(e string,c string) error{
	postform := fmt.Sprintf("code=%s",c)
	url:= fmt.Sprintf("/v1/email/%s",e)
	body,err:= HTTP_Resquest("POST",url,"",postform)
	if err != nil{ return err}
	result :=&HTTP_Body{}
	if err := json.Unmarshal(body,result);err!=nil{
		return err
	}
	//fmt.Println(result)
	if !result.Success{ return errors.New(result.Error) }
	return nil
}

func HTTP_RegWithEmail(e string,p string) error{
	postform := fmt.Sprintf("email=%s&password=%s",e,p)
	body ,err := HTTP_Resquest("POST","/v1/user","mode=email",postform)
	if err != nil{ return err}
	result:=&HTTP_Body{}
	if err := json.Unmarshal(body,result);err!=nil{
		return err
	}
	fmt.Println(result)
	if !result.Success{ return errors.New(result.Error) }
	return nil
}

func HTTP_Resquest(method string,router string,query string,postform string) ([]byte,error){
	ip:=beego.AppConfig.String("resfapi_ip")
	port:=beego.AppConfig.String("resfapi_port")
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s%s?%s",ip,port,router,query)
	request,err := http.NewRequest(method,url,strings.NewReader(postform))
	if err != nil{return nil,err}
	if method =="POST"{
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	defer request.Body.Close()
	resp,err := client.Do(request)
	defer resp.Body.Close()
	if err !=nil {return nil,err}
	body,err :=ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	return body,nil
}
