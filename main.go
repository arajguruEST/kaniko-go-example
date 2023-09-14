package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Println("API is running on port 8080")

	handler := http.NewServeMux()

	handler.HandleFunc("/", SayHello)
	http.ListenAndServe("0.0.0.0:8080", handler)

}

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Hello there!, You have just hit the Go container`)
}
