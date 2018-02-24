package main;

import (
	"net/rpc"
	"net/http"
	"log"
	"os/exec"
	"io/ioutil"
	"fmt"
	"context"
	//"syscall"
	"os"
	"time"
)

//go对RPC的支持，支持三个级别：TCP、HTTP、JSONRPC
//go的RPC只支持GO开发的服务器与客户端之间的交互，因为采用了gob编码

//注意字段必须可导出
type Params struct {
	Width, Height int;
}
//type Cmd struct {
//	cmdline string
//}
type CommandParam struct {
	Commandname string
	Commandargs []string
}
type Rect struct{}
var processmap =make(map[int]CommandParam)
//函数必须是导出的
//必须有两个导出类型参数
//第一个参数是接收参数
//第二个参数是返回给客户端参数，必须是指针类型
//函数还要有一个返回值error
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height;
	return nil;
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2;
	return nil;
}
/*
*执行cmd命令的函数Run
*/
func (r *Rect) Run(params CommandParam,ret *string) error{
	fmt.Println(params)
	go func(){
		cmd :=exec.Command(params.Commandname,params.Commandargs...)
		//fmt.Println(cmd.Process.Pid)
		stdout, err := cmd.StdoutPipe()
		if err != nil {

			*ret =err.Error()
		}else {
			// 运行命令
			if err := cmd.Start(); err != nil {
				*ret =err.Error()
			}else {
				// 读取输出结果
				opBytes, err := ioutil.ReadAll(stdout)
				if err != nil {
					*ret =err.Error()
				}else {
					*ret =string(opBytes)
				}
			}
		}
		// 保证关闭输出流
		defer stdout.Close()
	}()
	return nil
}

func (r *Rect) RunBack(params CommandParam,ret *string) error {
	ctx,_ := context.WithCancel(context.Background())
	cmd :=exec.CommandContext(ctx,params.Commandname,params.Commandargs...)
	//cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine:""} //linux  not compat
	cmd.Stdout = os.Stdout
	cmd.Start()
	time.Sleep(1 * time.Second)
	processid :=cmd.Process.Pid
	processmap[processid]=params
	fmt.Println("退出程序中...", cmd.Process.Pid)
	//cancel()   //是否杀死进程
	//cmd.Wait() //是否等待进程结束
	return nil
}

func (r *Rect) Runcmd(params CommandParam,ret *string) error {
	cmd :=exec.Command(params.Commandname,params.Commandargs...)
	//将其他命令传入生成出的进程
	//给新进程设置文件描述符，可以重定向到文件中
	cmd.Stdin=os.Stdin
	cmd.Stdout=os.Stdout
	cmd.Stderr=os.Stderr
	//开始执行新进程，不等待新进程退出
	cmd.Start()
	return nil
}

func healthcheck(){
	/*
	*进程状态健康检测函数
	*/
	for{
		for k,v :=range processmap{
			fmt.Println(k,v)
		}
		time.Sleep(10*time.Second)
	}
}
func main() {
	rect := new(Rect);
	//注册一个rect服务
	rpc.Register(rect);
	//把服务处理绑定到http协议上
	rpc.HandleHTTP();
	log.Println("start rpc server on 8080!")
	go healthcheck()
	err := http.ListenAndServe(":8080", nil);
	if err != nil {
		log.Fatal(err);
	}
}