package main;

import (
	"net/rpc"
	"log"
	"fmt"
)

type Params struct {
	Width, Height int;
}
var Cmd string = "ls"
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
	err4 := rpc.Call("Rect.Run", Cmd, &result);
	if err4 != nil {
		log.Fatal(err3);
	}
	fmt.Println(result)
}