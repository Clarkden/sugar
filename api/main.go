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
	"sugar/middleware"
	"sugar/router"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"github.com/rs/cors"
)

// go:embed data/migrations/*.sql
var embedMigrations embed.FS

func main() {

	db := initDB()

	// goose.SetBaseFS(embedMigrations)

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

	q := sugar.New(db)

	routerConfig := types.RouterConfig{}
	h := handlers.NewHandler(q)
	m := middleware.NewMiddleware(q)

	router := c.Handler(router.NewRouter(&routerConfig, h, m))

	log.Println("Server is starting on port: ", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
