package main

import (
	"gqlgen-postgres-demo/graph"
	"gqlgen-postgres-demo/graph/generated"
	"gqlgen-postgres-demo/infra/postgres"
	"gqlgen-postgres-demo/usecase"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	db, err := postgres.NewDB()
	if err != nil {
		panic(err)
	}
	userGetter := postgres.NewUserGetter(db)
	todoGetter := postgres.NewTodoGetter(db)
	userUsecase := usecase.NewUser(userGetter, todoGetter)
	todoUsecase := usecase.NewTodo(todoGetter)
	u := struct {
		usecase.UserGetter
		usecase.TodoGetter
	}{
		userUsecase,
		todoUsecase,
	}
	resolver := graph.NewResolver(u)

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
