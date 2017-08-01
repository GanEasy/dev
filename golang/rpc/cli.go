package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
	_ "github.com/hprose/hprose-golang/rpc/fasthttp"
)

type HelloService struct {
	Hello func(string) string
}

func main() {

	client := rpc.NewClient("http://127.0.0.1:8080/")

	var hello *HelloService
	client.UseService(&hello)
	result := hello.Hello("World")
	fmt.Println(result)
}
