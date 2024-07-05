package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home")
	fmt.Println("Enpoint hit: homePage")
}

type Response struct {
	Data string
}

func returnPics(w http.ResponseWriter, r *http.Request) {
	var apiURL string //APi key goes here
	fmt.Println("Endpoint hit: returnPics")
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
	return
}
func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/request", returnPics)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Mux router")
	handleRequest()
}
