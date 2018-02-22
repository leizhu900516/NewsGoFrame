package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"reflect"
)

func main(){
	resp,err := http.Get("http://www.baidu.com/")
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()
	body,err :=ioutil.ReadAll(resp.Body)
	fmt.Println(reflect.TypeOf(body),string(body))
}
