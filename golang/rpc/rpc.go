package main

import (
	"net/http"

	"github.com/hprose/hprose-golang/rpc"
)

func sayJson(name map[string]string) map[string]string {
	return name
}

func helloJson(name map[string]string) string {
	return "Hello " + name["a"] + "!"
}

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello)
	service.AddFunction("sayJson", sayJson)
	http.ListenAndServe(":8080", service)
}
