package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	fmt.Println("Go Docker")

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello World")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}