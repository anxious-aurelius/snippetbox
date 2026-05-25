package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	
	addr := flag.String("addr", "127.0.0.1:4003", "Network interface and port for the service")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level : slog.LevelDebug,
		AddSource : true, 
	}))


	fileserver := http.FileServer(http.Dir("./ui/static"))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	logger.Info("starting server","addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}