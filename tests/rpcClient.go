package main;

import (
	"net/rpc"
	"log"
	"fmt"
)

type Params struct {
	Width, Height int;
}
type Commandparam struct {
	Commandname string
	Commandargs []string
}

func main() {
	//连接远程rpc服务
	rpc, err := rpc.DialHTTP("tcp", "127.0.0.1:8080");
	if err != nil {
		log.Fatal(err);
	}
	ret := 0;
	//调用远程方法
	//注意第三个参数是指针类型
	err2 := rpc.Call("Rect.Area", Params{50, 100}, &ret);
	if err2 != nil {
		log.Fatal(err2);
	}
	fmt.Println(ret,">>>>");
	err3 := rpc.Call("Rect.Perimeter", Params{50, 100}, &ret);
	if err3 != nil {
		log.Fatal(err3);
	}

	fmt.Println(ret);
	var result string
	err4 := rpc.Call("Rect.Run", Commandparam{"ls",[]string{"-ll"}}, &result);
	if err4 != nil {
		log.Fatal(err3);
	}
	fmt.Println(result)
	//
	var result1 string
	err5 := rpc.Call("Rect.RunBack", Commandparam{"ls",[]string{"-ll"}}, &result1);
	if err5 != nil {
		log.Fatal(err3);
	}
	fmt.Println(result1)

	//var result1 string
	//err5 := rpc.Call("Rect.Runcmd", Commandparam{"python",[]string{"test.py"}}, &result1);
	//if err5 != nil {
	//	log.Fatal(err3);
	//}
	//fmt.Println(result1)
}