package main

import (
	"gqlgen-postgres-demo/graph"
	"gqlgen-postgres-demo/graph/generated"
	"gqlgen-postgres-demo/infra/postgres"
	"gqlgen-postgres-demo/usecase"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	db, err := postgres.NewDB()
	if err != nil {
		panic(err)
	}
	getter := postgres.NewUserGetter(db)
	usecase := usecase.NewUser(getter)
	resolver := graph.NewResolver(usecase)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
