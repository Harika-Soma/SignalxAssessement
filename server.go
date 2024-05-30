package main

import (
	"log"
	"net/http"
	"os"
	"supplychain/db"
	"supplychain/graph"
	"supplychain/store"

	"supplychain/pkg/auth"
	Direct "supplychain/pkg/directives"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	router := chi.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	d := db.New()
	router.Use(auth.Middleware(d))
	supply := store.NewSupplyChainStore(d)
	c := graph.Config{Resolvers: &graph.Resolver{
		SupplyChainStore: supply,
	}}

	c.Directives.Auth = Direct.Auth
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	handler := cors.AllowAll().Handler(router)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
