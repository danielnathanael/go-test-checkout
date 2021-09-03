package main

import (
	"go-test-checkout/database"
	"go-test-checkout/graph"
	"go-test-checkout/graph/generated"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = ":4000"

func main() {
	db := database.GetDatabase()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			DB: db,
		},
	}))

	http.Handle("/", playground.Handler("Backend Candidate Test", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(defaultPort, nil))
}
