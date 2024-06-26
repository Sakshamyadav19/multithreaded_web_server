package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"github.com/Sakshamyadav19/multithreaded_web_server/cache"
	"github.com/Sakshamyadav19/multithreaded_web_server/utils"
)

func HandleRequest(conn net.Conn, wg *sync.WaitGroup,cache *cache.LRU) {
	defer wg.Done()
	defer conn.Close()

	reader := bufio.NewReader(conn)

	req, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("ERror Reading Reques")
	}

	url := utils.ParseUrl(req)

	var response string

	if cachedResp, ok := cache.Get(url); ok {
		response = cachedResp.(string)
		fmt.Println("Cached Respones!!!!")
	} else {
		res, err := http.Get("http://" + url)

		if err != nil {
			fmt.Println("Error fetching URL:", err)
			response = "Error: Unable to fetch URL"
		} else {
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			response = string(body)
			cache.Set(url,response)
			fmt.Println("Cached")
		}

	}

	httpResponse := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\n\r\n%s", len(response), response)
	_, err = conn.Write([]byte(httpResponse))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}