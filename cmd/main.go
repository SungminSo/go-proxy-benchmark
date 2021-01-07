package main

import (
	"github.com/SungminSo/simple-go-proxy/req"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	req.SetProxyBasic(proxy)

	log.Println("asdfasdfasdf")
	log.Fatal(http.ListenAndServe(":8080", proxy))
}