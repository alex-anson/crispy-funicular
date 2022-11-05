package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Using the gorilla/mux 3rd party router package instead of the standard library
// net/http router. Allows you to more easily perform tasks such as parsing path
// or query params.

type Movie struct {
	Id          string `json:"Id"`
	Title       string `json:"Title"`
	Desc        string `json:"Desc"`
	ReleaseYear int    `json:"ReleaseYear"`
}

// Global Movies array. Can populate in the `main` function to simulate a db
var MovieList []Movie

const allAtOnceDesc = "When an interdimensional rupture unravels reality, an unlikely hero must channel her newfound powers to fight bizarre and bewildering dangers from the multiverse as the fate of the world hangs in the balance."
const troopersDesc = "Five Vermont state troopers, avid pranksters with a knack for screwing up, try to save their jobs and out-do the local police department by solving a crime."

// Handles requests to the root URL.
func homePage(w http.ResponseWriter, r *http.Request) {
	// func Fprintf(w io.Writer, format string, a ...any) (n int, err error)
	// Fprintf formats according to a format specifier and writes to w. It returns
	// the number of bytes written and any write error encountered.
	fmt.Fprintf(w, "Welcome to the homepage of my first Go repo 🂡")
	fmt.Println("Endpoint hit: homepage")
}

// Matches the URL path hit with a defined function
func handleRequests() {
	// Create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	// Add "/movies" endpoint & map it to the getMovieList ƒn
	myRouter.HandleFunc("/movies", getMovieList)
	// NOTE: Order matters. Must be before the other "/movie" endpoint
	myRouter.HandleFunc("/movie", addMovie).Methods("POST")
	// {id} = path variable
	myRouter.HandleFunc("/movie/{id}", getMovie)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

// Obvi most important ƒn ✨
func main() {
	// Will execute when you `go run` this file
	fmt.Println("Mux Routers 🦊")
	MovieList = []Movie{
		{Id: "1", Title: "Everything Everywhere All at Once", Desc: allAtOnceDesc, ReleaseYear: 2022},
		{Id: "2", Title: "Super Troopers", Desc: troopersDesc, ReleaseYear: 2001},
	}
	handleRequests()
}

func getMovieList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: /movies")
	// Encodes the movies into a JSON string
	json.NewEncoder(w).Encode(MovieList)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit. /movie/{id}")

	routeVariables := mux.Vars(r)
	key := routeVariables["id"]

	// Loop over moviesList
	for _, movie := range MovieList {
		if movie.Id == key {
			// Return the movie encoded as JSON
			json.NewEncoder(w).Encode(movie)
		}
	}
}

// the "C" of CRUD
func addMovie(w http.ResponseWriter, r *http.Request) {
	postBody, _ := io.ReadAll(r.Body)

	// %+v   value in a default format, with field name.
	// use when printing structs
	fmt.Fprintf(w, "%+v", string(postBody))
}
