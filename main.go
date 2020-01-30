package main

import (
	"fmt"
	_ "gin-skeleton/config"
	"gin-skeleton/route"
)

func main() {
	r := route.InitRouter()
	err := r.Run() // default 8080
	if err != nil {
		fmt.Printf("启动失败 %s", err.Error())
	}
}
