package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/JamieBShaw/golang-graphql-server/graph/generated"
	"github.com/JamieBShaw/golang-graphql-server/graph/resolvers"
	"github.com/JamieBShaw/golang-graphql-server/postgres"
	"github.com/go-pg/pg/v10"
	"github.com/hashicorp/go-hclog"
)

const defaultPort = "8080"

func main() {

	l := hclog.Default()

	db := postgres.NewDB(&pg.Options{

		User:     "postgres",
		Password: "postgres",
		Database: "meetmeup_dev",
	})

	defer db.Close()

	db.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {

		port = defaultPort
	}

	conf := generated.Config{
		Resolvers: &resolvers.Resolver{
			UsersRepo:   postgres.UsersRepo{DB: db, Log: l},
			MeetupsRepo: postgres.MeetupsRepo{DB: db, Log: l},
		}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(conf))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", resolvers.DataLoaderMiddleware(db, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
