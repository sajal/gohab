package main

import (
	"flag"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

func l2ping(mac string) bool {
	log.Println("Checking ", mac)
	cmd := exec.Command("l2ping", "-c", "1", mac)
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return false
	}
	return true

}

func main() {
	var addr = flag.String("addr", ":8081", "address/port to listen on")
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mac := strings.Split(r.URL.Path, "/")[1]
		if mac == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if l2ping(mac) {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	s := &http.Server{
		Addr:           *addr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
