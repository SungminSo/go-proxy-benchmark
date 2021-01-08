package on_request

import (
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func DoLogTest() goproxy.ReqHandler {
	log.Println("5")
	return goproxy.FuncReqHandler(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		return req, nil
	})
}

func handleConnectLogTest() goproxy.HttpsHandler {
	log.Println("4")
	return goproxy.FuncHttpsHandler(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		return goproxy.OkConnect, host
	})
}

func SetProxyBasic(proxy *goproxy.ProxyHttpServer, f func(user, passwd string) bool) {
	ip := "114.203.110.70"

	proxy.OnRequest().DoFunc(func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		log.Println("2")
		r.Header.Set("X-GoProxy","yxorPoG-X")

		return r, nil
	})

	proxy.OnRequest().Do(DoLogTest())

	proxy.OnRequest().HandleConnectFunc(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		fmt.Println("여기로는 안오냐ㅑㅑㅑㅑ")
		log.Println("3")
		return goproxy.OkConnect, host
	})

	proxy.OnRequest().HandleConnect(handleConnectLogTest())

	proxy.OnRequest(goproxy.Not(goproxy.SrcIpIs(ip))).DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			log.Println("1")
			if !f("qwerqwer", "qwerqwer") {
				return r, goproxy.NewResponse(r,
					goproxy.ContentTypeText, http.StatusForbidden,
					"Don't waste your time!")
			}
			return r, goproxy.NewResponse(r,
				goproxy.ContentTypeText,
				http.StatusOK,
				"Asdfasdfas")
		})
}