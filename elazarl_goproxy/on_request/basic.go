package on_request

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func DoBasic() goproxy.ReqHandler {
	log.Println("Do")
	return goproxy.FuncReqHandler(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		return req, nil
	})
}

func handleConnectBasic() goproxy.HttpsHandler {
	log.Println("handleConnect")
	return goproxy.FuncHttpsHandler(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		return goproxy.OkConnect, host
	})
}

func SetProxyBasic(proxy *goproxy.ProxyHttpServer, f func(user, passwd string) bool) {
	ip := "114.203.110.70"

	proxy.OnRequest().HandleConnectFunc(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		log.Println("HandleConnectFunc")
		return goproxy.OkConnect, host
	})

	proxy.OnRequest().HandleConnect(handleConnectBasic())

	proxy.OnRequest().DoFunc(func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		log.Println("DoFunc0")
		r.Header.Set("X-GoProxy","yxorPoG-X")

		return r, nil
	})

	proxy.OnRequest().Do(DoBasic())

	proxy.OnRequest(goproxy.Not(goproxy.SrcIpIs(ip))).DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			log.Println("DoFunc1")
			if !f("qwerasdf", "qwerasdf") {
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