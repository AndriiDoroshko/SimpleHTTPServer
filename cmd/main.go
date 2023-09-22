package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	log.Print("simplehttp: Enter main()")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
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
