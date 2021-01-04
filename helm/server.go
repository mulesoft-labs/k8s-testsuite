package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	server := &http.Server{
		Addr: ":80",
	}
	count := 0

	http.HandleFunc("/cat", func(w http.ResponseWriter, r *http.Request) {
		print(w, r, "Cool Cat")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		print(w, r, "forward slash")
	})

	http.HandleFunc("/dog", func(w http.ResponseWriter, r *http.Request) {
		print(w, r, "Dapper Dog")
	})

	http.HandleFunc("/500", func(w http.ResponseWriter, r *http.Request) {
		count++
		log.Printf("Error until 5: %d", count%5)

		if count%5 == 0 {
			print(w, r, "Erroneous Elephant")
			return
		}
		time.Sleep(1 * time.Second)
		http.Error(w, "Some 500 Error", 500)
	})

	start(server)
}

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

func start(server *http.Server) {
	log.Print("Server listening NOW")
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal(err.Error())
	}
}
