package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Proxy to Services")

	w.Write([]byte("Services Running!"))
}
