package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"database/sql"
 	"github.com/anxious-aurelius/snippetbox/internal/models"

	_ "github.com/go-sql-driver/mysql"

)

type application struct {
	logger *slog.Logger
	snippetModel *models.SnippetModel
}

func main() {

	addr := flag.String("addr", "127.0.0.1:4003", "Network interface and port for the service")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	dsn := "web:0711@tcp(127.0.0.1:3306)/snippetbox"
	db := getDB(dsn, logger)

	app := &application{
		logger: logger,
		snippetModel: &models.SnippetModel{DB: db},
	}

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func getDB(dsn string, logger *slog.Logger) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	if err := db.Ping(); err != nil{
		logger.Error(err.Error())
		os.Exit(1)
	} 

	logger.Info("Successfully connected to the DB")

	return db
}