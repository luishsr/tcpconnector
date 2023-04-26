package main

import (
	"fmt"

	"github.com/luishsr/tcpconnector"
)

func main() {
	fmt.Println("main() method")

	//calls tcpconnector method
	response := tcpconnector.Connect()

	fmt.Println(response)
}
