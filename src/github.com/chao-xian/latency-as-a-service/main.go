package main

import (
	"fmt"
	"net/http"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "Yo I'm Batman")
}

func latency(w http.ResponseWriter, r *http.Request) {
	durationParam := r.FormValue("duration")
	youWait, err := time.ParseDuration(durationParam)

	if err != nil {
		http.Error(w, "Nope "+err.Error(), http.StatusBadRequest)
		return
	} else {
		time.Sleep(youWait)
		fmt.Fprintf(w, "Batman says OK after %v", youWait)
	}
}

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/latency", latency)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
