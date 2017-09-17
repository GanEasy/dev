package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/robertkrimen/otto"
)

//"io/ioutil"

//	"net/http/cookiejar"

// "bytes"

func main() {
	f, err := os.Open("plugins.js")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	buff := bytes.NewBuffer(nil)

	if _, err := buff.ReadFrom(f); err != nil {
		panic(err)
	}
	runtime := otto.New()
	if _, err := runtime.Run(buff.String()); err != nil {
		panic(err)
	}

	a := 1
	b := 2
	jsa, err := runtime.ToValue(a)
	if err != nil {
		panic(err)
	}
	jsb, err := runtime.ToValue(b)
	if err != nil {
		panic(err)
	}
	result, err := runtime.Call("addnum", nil, jsa, jsb)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	// out, err := result.ToInterger()
	// if err != nil {
	// 	panic(err)
	// }
}
