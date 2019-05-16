package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/forrana/grammar-trainer/api"
	"github.com/forrana/grammar-trainer/api/data"
    "github.com/go-chi/chi"
    "github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
    defer data.DB.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query",handler.GraphQL(gramma_trainer.NewExecutableSchema(gramma_trainer.Config{Resolvers: &gramma-trainer.Resolver{}})))


	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}


	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
