package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// URL stores your ngrok dynamic address
var (
	URL    = os.Getenv("URL")
	secret = os.Getenv("SECRET")
)

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
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

func handleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	log.Println("from ", r.RemoteAddr)
	serveReverseProxy(URL, w, r)
}

func changeURL(w http.ResponseWriter, r *http.Request) {
	log.Println("from ", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal error while Parsing")
		return
	}

	key := r.URL.Query().Get("secret")
	url := r.URL.Query().Get("url")

	if key != secret {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Forbidden")
		return
	}

	URL = url
	fmt.Printf("now URL forward to: %s\n", url)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success")
}

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

func main() {
	log.Printf("Server will run on %s\n", getListenAddress())
	log.Printf("Redirecting to a url: %s\n", URL)

	http.HandleFunc("/update_ngrok", changeURL)
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		log.Panicln("server error", err.Error())
	}
}
