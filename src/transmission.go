package main

import (
	"fmt"
	"github.com/cdelorme/go-transmission-api"
)


func main()  {
	fmt.Println("hello world!")
	trans := transmission.Transmission{}
	//trans.Configure("/Users/liuweipeng/go/src/github.com/transmission-rpc/src/conf/transmission-daemon/settings.json")

	results, error := trans.Get()
	if (error != nil) {
		fmt.Println(error)
	}
	fmt.Println(len(results))
	//trans.Add("")

}
