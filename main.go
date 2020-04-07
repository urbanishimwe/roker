package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host
	proxy.ServeHTTP(res, req)
}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	log.Println("from ", req.RemoteAddr)
	url := os.Getenv("URL")
	serveReverseProxy(url, res, req)
}

// Get env var or default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getListenAddress() string {
	port := getEnv("PORT", "3000")
	return ":" + port
}

func logSetup() {
	url := os.Getenv("URL")
	log.Printf("Server will run on %s\n", getListenAddress())
	log.Printf("Redirecting to a url: %s\n", url)
}

func main() {
	logSetup()

	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		log.Panicln("server error", err.Error())
	}
}
