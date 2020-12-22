package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/JHaig343/graphql-server/graph"
	"github.com/JHaig343/graphql-server/graph/generated"
	database "github.com/JHaig343/graphql-server/internal/pkg/db/migrations/mysql"
	// "github.com/go-chi/chi"
	// "github.com/golang-migrate/migrate/v4/database"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// router := chi.NewRouter()

	database.InitDB()
	database.Migrate()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
