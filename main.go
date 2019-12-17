// Package declration, indicates you'll compile an executable.
package main

// Imports libs
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./tpl/*.html",
	}

	rnd = renderer.New(opts)
}

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

// When we access "/pokemon", I want to render my template named "pokemon".
func pokemonPage(w http.ResponseWriter, r *http.Request) {
	// Set the content type to HTML so it doesn't print to console.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Making a request to the PokeAPI
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Saving the body of the response from the PokeAPI to responseData.
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Printing to the window the responseData JSON.
	fmt.Fprint(w, string(responseData))

	rnd.HTML(w, http.StatusOK, "pokemon", nil)
}

// When we access "/", I want to render my template named "index".
func indexPage(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "index", nil)
}

// Entry point for executables
func main() {
	// mux is an HTTP request multiplexer... matches the URL of incoming requests against a list of registered patterns
	// My templates wouldn't render to the page with http.
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexPage)
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/headers", headers)
	mux.HandleFunc("/pokemon", pokemonPage)
	// With http.HandleFunc I would be passing nil into the headers param, but here I need mux.
	http.ListenAndServe(":8090", mux)
}
