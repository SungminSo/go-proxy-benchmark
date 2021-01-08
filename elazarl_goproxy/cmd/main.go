package main

import (
	"github.com/SungminSo/simple-go-proxy/elazarl_goproxy/on_request"
	"github.com/SungminSo/simple-go-proxy/elazarl_goproxy/on_response"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"os"
)

func main() {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	on_request.SetProxyBasic(proxy, func(user, pwd string) bool {
		return user == username && password == pwd
	})

	on_response.SetProxyStats(proxy)

	log.Println("start proxy server at :8080")
	log.Fatal(http.ListenAndServe(":8080", proxy))
}