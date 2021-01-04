package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	proxyproto "github.com/pires/go-proxyproto"
)

type H struct{}

func print(w http.ResponseWriter, r *http.Request, kind string) {
	fmt.Println("\n==============================\n")
	fmt.Fprintf(w, "*I am a %s*\n", kind)
	fmt.Fprintf(w, "* Method=%s\n", r.Method)
	log.Printf("- Method=%s\n", r.Method)
	fmt.Fprintf(w, "* URL=%s\n", r.URL)
	log.Printf("- URL=%s\n", r.URL)
	fmt.Fprintf(w, "* Proto=%s\n", r.Proto)
	log.Printf("- Proto=%s\n", r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "* Header=%s:%s\n", k, strings.Join(v, ","))
		log.Printf("- Header=%s:%s\n", k, strings.Join(v, ","))
	}

	fmt.Fprintf(w, "* Host=%s\n", r.Host)
	log.Printf("- Host=%s\n", r.Host)
	fmt.Fprintf(w, "* RemoteAddr=%s\n", r.RemoteAddr)
	fmt.Fprint(w, "** END BODY **\n")
	log.Printf("- RemoteAddr=%s\n", r.RemoteAddr)
	log.Print("Request served\n\n")
}

func (h *H) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	print(w, r, "Cool Cat")
}

func main() {
	// Create a listener
	addr := "0.0.0.0:80"
	list, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("couldn't listen to %q: %q\n", addr, err.Error())
	}

	// Wrap listener in a proxyproto listener
	proxyListener := &proxyproto.Listener{Listener: list}
	defer proxyListener.Close()
	for {
		// Wait for a connection and accept it
		conn, err := proxyListener.Accept()
		if err != nil {
			log.Printf("ERROR:      accepting %+v", err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// Print connection details
	if conn.LocalAddr() == nil {
		log.Printf("ERROR:      couldn't retrieve local address")
	}
	log.Printf("local address: %q", conn.LocalAddr().String())

	if conn.RemoteAddr() == nil {
		log.Printf("ERROR:      couldn't retrieve remote address")
	}
	log.Printf("remote address: %q", conn.RemoteAddr().String())

	conn.Write([]byte("XX"))
}
