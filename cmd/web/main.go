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

	app := &application{
		logger: logger,
	}

	logger.Info("starting server","addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}