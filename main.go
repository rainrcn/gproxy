package main

import (
	"fmt"
	"github.com/rainrcn/gproxy/config"
	"net/http"
)

func main() {
	conf := config.LoadAppConfig()

	handler := GProxyHandler{
		Config: conf,
	}

	if conf.Http.Enabled {
		errHttp := http.ListenAndServe(conf.Http.ListenAddr, &handler)

		if errHttp != nil {
			fmt.Println("http listen error.", errHttp)
		}
	}

	if conf.Https.Enabled {
		errHttp := http.ListenAndServeTLS(conf.Https.ListenAddr, conf.Https.CertFile, conf.Https.KeyFile, &handler)

		if errHttp != nil {
			fmt.Println("https listen error.", errHttp)
		}
	}
}
