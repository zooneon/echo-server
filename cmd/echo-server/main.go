package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func echo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("echo"))
}

func getHostname(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write([]byte(hostname))
}

func getIP(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	w.Write([]byte(localAddr.IP.String()))
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%s %s %s\n", r.Method, r.URL, r.Proto)))
	w.Write([]byte(fmt.Sprintf("Host: %s\n", r.Host)))
	w.Write([]byte(fmt.Sprintf("RemoteAddr: %s\n", r.RemoteAddr)))
	for k, v := range r.Header {
		w.Write([]byte(fmt.Sprintf("%s: %s\n", k, v)))
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("echo server listening on port %s\n", port)

	http.HandleFunc("/", echo)
	http.HandleFunc("/hostname", getHostname)
	http.HandleFunc("/ip", getIP)
	http.HandleFunc("/request", getRequest)

	err := http.ListenAndServe(":"+port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal("server closed")
	} else if err != nil {
		log.Fatalf("server error: %s\n", err)
	}
}