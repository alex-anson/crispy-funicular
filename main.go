package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	// Using the gorilla/mux 3rd party router package instead of the standard library
	// net/http router. Allows you to more easily perform tasks such as parsing path
	// or query params.
	"github.com/gorilla/mux"
)

const PORT = ":10000"

type Movie struct {
	Id          string `json:"Id"`
	Title       string `json:"Title"`
	Desc        string `json:"Desc"`
	ReleaseYear int    `json:"ReleaseYear"`
}

// NOTE: updating a global variable (MovieList) in order to keep this simple. Not
// doing any checks to ensure race conditions don't happen. This code isn't "thread-safe".

// Global Movies array. Can populate in the `main` function to simulate a db
var MovieList []Movie

// Descriptions
const allAtOnceDesc = "When an interdimensional rupture unravels reality, an unlikely hero must channel her newfound powers to fight bizarre and..."
const troopersDesc = "Five Vermont state troopers, avid pranksters with a knack for screwing up, try to save their jobs and out-do the local police..."

// IMPORTANT:
// Matches the URL path hit with a defined function
func registerHandlers() {
	// Create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	// Add "/movies" endpoint & map it to the getMovieList ƒn
	myRouter.HandleFunc("/movies", getMovieList)
	// NOTE: Order matters. Must be before the other "/movie" endpoint
	myRouter.HandleFunc("/movie", addMovie).Methods("POST")
	// Order still matters.
	myRouter.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	myRouter.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	// {id} = path variable
	myRouter.HandleFunc("/movie/{id}", getMovie)

	log.Fatal(http.ListenAndServe(PORT, myRouter))
}

// Obvi most important ƒn ✨
func main() {
	// Will execute when you `go run` this file
	fmt.Println("Mux Routers 🦊")
	MovieList = []Movie{
		{Id: "1", Title: "Everything Everywhere All at Once", Desc: allAtOnceDesc, ReleaseYear: 2022},
		{Id: "2", Title: "Super Troopers", Desc: troopersDesc, ReleaseYear: 2001},
	}
	registerHandlers()
}

// SECTION: Route Handlers

// Handles requests to the root URL.
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage of my first Go repo 🂡")
	fmt.Println("Endpoint hit: homepage")
}

// "R" of CRUD
func getMovieList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: /movies")
	// Encodes the movies into a JSON string
	json.NewEncoder(w).Encode(MovieList)
}

// "R" of CRUD
func getMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: /movie/{id}")

	routeVariables := mux.Vars(r)
	key := routeVariables["id"]

	// Loop over MovieList
	for _, movie := range MovieList {
		if movie.Id == key {
			// Return the movie encoded as JSON
			json.NewEncoder(w).Encode(movie)
		}
	}
}

// the "C" of CRUD
// ... doesn't take validation into consideration.
func addMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit (POST request received): /movie")

	postBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err.Error(), "\nproblem reading request body")
		return
	}
	// Unmarshal the request body JSON into a new `Movie` struct.
	var movie Movie
	err = json.Unmarshal(postBody, &movie)

	if err != nil {
		fmt.Println(err.Error(), "\nproblem unmarshalling")
	}

	// Update the global MovieList to include the new movie.
	MovieList = append(MovieList, movie)

	// %+v  -> value in a default format, with field name.
	// use when printing structs
	fmt.Fprintf(w, "%+v", string(postBody))
}

// the "D" of CRUD
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit (DELETE request received): /movie/{id}")

	// Extract the ID
	routeVariables := mux.Vars(r)
	deleteId := routeVariables["id"]

	// Loop the movies, remove any entry whose Id property matches deleteId
	for index, movie := range MovieList {
		if movie.Id == deleteId {
			MovieList = append(MovieList[:index], MovieList[index+1:]...)
		}
	}
}

// the "U" of CRUD
func updateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit (PUT request received): /movie/{id}")

	// Extract the id from the route
	routeVariables := mux.Vars(r)
	updateId := routeVariables["id"]

	// Get the request body
	putBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err.Error(), "\nproblem reading request body")
		return
	}

	// Unmarshal the request body JSON into a new `Movie` struct.
	var movie Movie
	err = json.Unmarshal(putBody, &movie)

	if err != nil {
		fmt.Println(err.Error(), "\nproblem unmarshalling")
	}

	for index, xmovie := range MovieList {
		if xmovie.Id == updateId {
			// Remove the movie we're trying to update
			MovieList = append(MovieList[:index], MovieList[index+1:]...)

			// Manually construct the movie because I don't know how else to do it
			updatedMovie := Movie{Id: updateId, Title: movie.Title, Desc: movie.Desc, ReleaseYear: movie.ReleaseYear}

			// Update the global MovieList to include the "new" (updated) movie.
			MovieList = append(MovieList, updatedMovie)
		}
	}

	fmt.Fprintf(w, "%+v", string(putBody))
}
