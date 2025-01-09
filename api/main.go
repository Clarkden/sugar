package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	sugar "sugar/data"
	"sugar/globals/types"
	"sugar/handlers"
	"sugar/router"

	"github.com/pressly/goose/v3"
	"github.com/rs/cors"
)

// go:embed data/migrations/*
var embedMigrations embed.FS

func main() {

	db := initDB()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "data/migrations"); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	routerConfig := types.RouterConfig{}

	q := sugar.New(db)

	h := handlers.NewHandler(q)

	router := c.Handler(router.NewRouter(&routerConfig, h))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
