package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("KLSH simple HTTP server...")
	fmt.Println("Connect to http://localhost:8080/")

	http.HandleFunc("/", getRoot)

	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v\n", r)
	w.Write([]byte("Root page"))
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%#v\n", r)
	w.Write([]byte("Hello page"))
}
