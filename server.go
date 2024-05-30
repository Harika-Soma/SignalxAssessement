package main

import (
	"log"
	"net/http"
	"os"
	"supplychain/db"
	"supplychain/graph"
	"supplychain/store"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	d := db.New()
	supply := store.NewSupplyChainStore(d)
	c := graph.Config{Resolvers: &graph.Resolver{
		SupplyChainStore: supply,
	}}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
