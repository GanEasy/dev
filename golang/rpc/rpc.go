package main

import (
	"net/http"

	"github.com/hprose/hprose-golang/rpc"
)

func helloJson(name map[string]string) string {
	return "Hello " + name["a"] + "!"
}

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello)
	http.ListenAndServe(":8080", service)
}
