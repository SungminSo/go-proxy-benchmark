package req


import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"time"
)

func SetProxyBasic(proxy *goproxy.ProxyHttpServer) {
	ip := "114.203.110.70"

	proxy.OnRequest(goproxy.Not(goproxy.SrcIpIs(ip))).DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx)(*http.Request, *http.Response) {
			log.Println("1")
			if h,_,_ := time.Now().Clock(); h >= 8 && h <= 19 {
				return r, goproxy.NewResponse(r,
					goproxy.ContentTypeText,http.StatusForbidden,
					"Don't waste your time!")
			}
			return r, nil
		})

	proxy.OnRequest().DoFunc(func(r *http.Request, ctx *goproxy.ProxyCtx)(*http.Request, *http.Response) {
		log.Println("2")
		r.Header.Set("X-GoProxy","yxorPoG-X")
		return r, nil
	})

	proxy.OnRequest(goproxy.SrcIpIs(ip)).HandleConnectFunc(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		log.Println("3")
		return goproxy.OkConnect, host
	})
}