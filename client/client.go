package main

import (
	"fmt"
)

func main() {
	fmt.Println("main() method")

	//calls tcpconnector method
	response := tcpconnector.Connect()

	fmt.Println(response)
}
