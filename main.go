package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, rw *http.Request) {
	id, err := strconv.Atoi(rw.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, rw)
		return
	}
	msg := fmt.Sprintf("Display widget with id %v", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Display a form for creating snippets"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4001")

	err := http.ListenAndServe("127.0.0.1:4001", mux)

	if err != nil {
		log.Fatal(err)
	}
}
