package main

import (
	"github.com/SungminSo/simple-go-proxy/req"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func main() {
	username, password := "qwerasdf", "qwerasdf"

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	req.SetProxyBasic(proxy, func(user, pwd string) bool {
		return user == username && password == pwd
	})

	log.Println("asdfasdfasdf")
	log.Fatal(http.ListenAndServe(":8080", proxy))
}