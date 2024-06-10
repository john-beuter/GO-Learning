package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:9803")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Close()
}
