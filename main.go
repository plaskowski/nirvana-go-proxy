package main

import (
    "flag"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func main() {
    port := flag.String("port", "3000", "Port to listen on")
    flag.Parse()

    targetURL, err := url.Parse("https://focus.nirvanahq.com")
    if err != nil {
        log.Fatalf("Invalid target URL: %v", err)
    }

    proxy := httputil.NewSingleHostReverseProxy(targetURL)

    proxy.Director = func(req *http.Request) {
        req.URL.Scheme = targetURL.Scheme
        req.URL.Host = targetURL.Host
        req.Host = targetURL.Host
    }

    proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
        log.Printf("Proxy error: %v", err)
        http.Error(rw, "Proxy error", http.StatusBadGateway)
    }

    server := &http.Server{
        Addr:    ":" + *port,
        Handler: proxy,
    }

    log.Printf("🟢 Proxy running at http://localhost:%s
", *port)
    log.Fatal(server.ListenAndServe())
}
