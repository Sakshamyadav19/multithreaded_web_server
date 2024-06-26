package main

import (
	"fmt"
	"net"
	"sync"
	"github.com/Sakshamyadav19/multithreaded_web_server/cache"
	"github.com/Sakshamyadav19/multithreaded_web_server/server"
)




func main() {
	fmt.Println("Welcome!!")

	var wg sync.WaitGroup

	listner, err := net.Listen("tcp", ":8080")
	cache:=cache.NewLRUCache(10)

	if err != nil {
		fmt.Print("Error Starting server")
	}

	defer listner.Close()

	for {

		conn, err := listner.Accept()

		if err != nil {
			fmt.Print("Error Connecting")
			continue
		}

		wg.Add(1)

		go server.HandleRequest(conn, &wg,cache)

	}
}
