package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func reverseProxy(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Reverse proxy request received at", time.Now())
	originServer, err := url.Parse("http://localhost:8080")
	fmt.Println(err)
	req.Host = originServer.Host
	req.URL.Host = originServer.Host
	req.URL.Scheme = originServer.Scheme
	req.RequestURI = ""
	fmt.Println(req.Host, req.URL.Scheme, req.URL.Host, req.RequestURI)

	originServerResponse, err := http.DefaultClient.Do(req)

	w.WriteHeader(http.StatusOK)
	io.Copy(w, originServerResponse.Body)
	fmt.Println(err)
}

func main() {
	http.HandleFunc("/", reverseProxy)
	err := http.ListenAndServe(":8081", nil)
	fmt.Println(err)
}
