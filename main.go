package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){

    fmt.Fprintf(w, "Welcome Golang test!")

    fmt.Println("Endpoint Hit: F03")

    fmt.Println("F002")


}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}
