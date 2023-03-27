package main

import (
	"bytes"
	"github.com/rainrcn/gproxy/config"
	"net/http"
	"strings"
)

type GProxyHandler struct {
	Config *config.AppConfig
}

func doProxy(url string, w http.ResponseWriter, r *http.Request) {
	newReq, errReq := http.NewRequest(r.Method, url, r.Body)

	if errReq != nil {
		return
	}

	//设置Req Header
	for k, _ := range r.Header {
		newReq.Header.Set(k, r.Header.Get(k))
	}

	newRes, errRes := http.DefaultClient.Do(newReq)

	if errRes != nil {
		return
	}

	//设置Res Header
	for k, _ := range newRes.Header {
		w.Header().Set(k, newRes.Header.Get(k))
	}

	defer newRes.Body.Close()

	//设置Res Body
	for true {
		buff := make([]byte, bytes.MinRead)

		n, err := newRes.Body.Read(buff)

		if err != nil || n < 1 {
			break
		}

		_, err = w.Write(buff[0:n])

		if err != nil {
			break
		}
	}
}

func (handler *GProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path

	//-- 判断是否只传入域名，如果是：则需要判断是否代理主域名
	if "/" == urlPath {
		mainDomain := handler.Config.Maps["/"]

		if len(mainDomain) > 0 {
			doProxy(mainDomain, w, r)
		}

		return
	}

	//-- 传入域名和路径时，分割域名后第一层
	urlSplits := strings.SplitN(urlPath, "/", 3)
	urlDomain := urlSplits[1]

	//从配置映射中获取目标域名
	newDomain := handler.Config.Maps[urlDomain]

	//判断是否配置有目标域名，如果没有配置，则查找是否配置有主域名映射
	if len(newDomain) < 1 {
		mainDomain := handler.Config.Maps["/"]

		if len(mainDomain) > 0 {
			newUrl := mainDomain + r.URL.Path + ("?" + r.URL.RawQuery) + "#" + r.URL.RawFragment

			doProxy(newUrl, w, r)
		}
	} else {
		urlDomainPath := urlPath[1+len(urlDomain):]
		newUrl := newDomain + urlDomainPath + ("?" + r.URL.RawQuery) + "#" + r.URL.RawFragment

		doProxy(newUrl, w, r)
	}
}
