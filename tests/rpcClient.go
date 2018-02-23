package main

import (
	"fmt"
	"time"
)

func daemon(){
	for{
		fmt.Println("hello")
		time.Sleep(time.Second*1)
	}
}
func main(){
	go daemon()
	for i:=1;i<10;i++{
		fmt.Println(i)
		time.Sleep(time.Second*2)
	}
}
