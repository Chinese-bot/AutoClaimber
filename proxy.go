package main

import (
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func Rand(n int) string {
	const letterBytes = "0123456789"

	bytes := make([]byte, n)
	for i := range bytes {
		bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(bytes)
}

func Premium() *url.URL {
	for {
		proxy, _ := url.Parse("http://Chinese2-cc-any-sid-" + Rand(8) + ":RovDwVZ@gw.rainproxy.io:5959")
		client := http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy), DisableKeepAlives: true}, Timeout: time.Second * 1}

		resp, err := client.Get("https://ipapi.co/json/")
		if err != nil {
			continue
		}
		_ = resp.Body.Close()

		return proxy
	}
}
