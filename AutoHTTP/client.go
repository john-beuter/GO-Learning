package main

import (
	"fmt"
	"net/http"
)

func main() {

	resp, err := http.Get("https://localhost:3333/One")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
}
