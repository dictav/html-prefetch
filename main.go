package main

import (
	"io"
	"log"
	"net/http"
	"syscall"
)

func main() {
	mux := http.NewServeMux()
	pwd, err := syscall.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		moz := r.Header.Get("X-moz")
		purpose := r.Header.Get("X-Purpose")
		log.Printf("%s %s moz=%s, purpose=%s\n", r.Method, r.URL.Path, moz, purpose)

		// Ignore cache
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		if r.URL.Path == "/header_next.html" {
			w.Header().Set("Link", "</prefetch.html>; rel=next")
		}

		if r.URL.Path == "/header_prefetch.html" {
			w.Header().Set("Link", "</prefetch.html>; rel=prefetch")
		}

		if r.URL.Path == "/header_prerender.html" {
			w.Header().Set("Link", "</prefetch.html>; rel=prerender")
		}

		h := http.FileServer(http.Dir(pwd))
		h.ServeHTTP(w, r)
	})

	mux.HandleFunc("/log", func(rw http.ResponseWriter, r *http.Request) {
		buf := make([]byte, 256)
		if _, err := r.Body.Read(buf); err != nil && err != io.EOF {
			http.Error(rw, "no body:"+err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("log: %s\n", buf)
	})

	mux.HandleFunc("/echo", func(rw http.ResponseWriter, r *http.Request) {
		buf := make([]byte, 256)
		if _, err := r.Body.Read(buf); err != nil && err != io.EOF {
			http.Error(rw, "no body:"+err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("echo: %s\n", buf)
		rw.Write(buf)
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
