package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gluebooze/movies_crud/functions"
	"github.com/gluebooze/movies_crud/types"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	functions.Movies = append(functions.Movies, types.Movie{ID: "1", Isbn: "45656", Title: "Movie One", Director: &types.Director{Firstname: "Akash", Lastname: "Anand"}})
	functions.Movies = append(functions.Movies, types.Movie{ID: "2", Isbn: "456209", Title: "Movie Two", Director: &types.Director{Firstname: "Harsh", Lastname: "Name"}})
	r.HandleFunc("/movies", functions.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", functions.GetMovie).Methods("GET")
	r.HandleFunc("/movies", functions.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", functions.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", functions.DeleteMovie).Methods("DELETE")

	fmt.Printf("startting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
