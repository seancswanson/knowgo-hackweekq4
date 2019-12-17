// Package declration, indicates you'll compile an executable.
package main

// Imports libs
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// What is "w"? Window?
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello, World!")
}

// This prints the HTTP request headers.
func headers(w http.ResponseWriter, req *http.Request) {
	// for loop
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func fetchPokemon(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(responseData))
}

// Entry point for executables
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/pokemon", fetchPokemon)
	http.ListenAndServe(":8090", nil)
}
