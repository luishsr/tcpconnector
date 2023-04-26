package main

import (
	"fmt"

	"github.com/luishsr/tcpconnector/connector"
)

func main() {
	fmt.Println("main() method")

	//calls tcpconnector method
	response := tcpconnector.Connect()

	fmt.Println(response)
}
