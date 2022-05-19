package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	"github.com/kcraley/howtographql/graph"
	"github.com/kcraley/howtographql/graph/generated"
	"github.com/kcraley/howtographql/internal/auth"
	database "github.com/kcraley/howtographql/internal/pkg/db/migrations/mysql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Create HTTP router mux and add auth middleware.
	router := chi.NewRouter()
	router.Use(auth.Middleware())

	// Initialize database connect and perform migrations.
	database.InitDB()
	database.Migrate()

	// Create
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
