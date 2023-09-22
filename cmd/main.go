package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	log.Print("Simple Http: Enter main()")
	http.HandleFunc("/", handler)

	server := &http.Server{Addr: ":8088"}

	server.ListenAndServe()
	log.Printf("main: serving on 8080")

	log.Printf("main: serving for 30 seconds")

	time.Sleep(30 * time.Second)

	server.Shutdown(context.Background())
	log.Print("main: server Shutdown")
}

// printing request headers/params
func handler(w http.ResponseWriter, r *http.Request) {

	log.Print("request from address: %q \n", r.RemoteAddr)

	fmt.Fprintf(w, "%s %s %s\n\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "\n\nHost = %q\n", r.Host)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "\n===> local IP: %q\n\n", GetOutboundIP())

	fmt.Fprintf(w, "Good job!")
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
