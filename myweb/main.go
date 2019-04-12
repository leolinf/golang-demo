package main

import (
	"fmt"
	"golang-demo/myweb/service"
)

func main() {
	fmt.Printf("start .....\n")
	service.StartWebServer("6767")
}
