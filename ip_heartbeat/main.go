package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func ip(w http.ResponseWriter, req *http.Request) {

	fmt.Println(req.RemoteAddr)
	fmt.Fprintf(w, "%s", strings.TrimSpace(req.RemoteAddr))
}

func main() {
	isServer := flag.Bool("s", false, "服务端启动")
	isClient := flag.Bool("c", false, "服务端启动")
	host := flag.String("host", "127.0.0.1:8080", "server host")
	listen := flag.String("listen", "127.0.0.1:8080", "server ListenAndServe")
	flag.Parse()
	if *isServer {
		http.HandleFunc("/ip", ip)
		log.Fatal(http.ListenAndServe(*listen, nil))
	}
	if *isClient {
		client := http.Client{
			Timeout: 30 * time.Second,
		}
		resp, err := client.Get(fmt.Sprintf("http://%s/ip", *host))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
		return
	}
	fmt.Println("choose a mode s or c")
}
