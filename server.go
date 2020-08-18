package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/JamieBShaw/golang-graphql-server/domain"
	"github.com/JamieBShaw/golang-graphql-server/graph/generated"
	"github.com/JamieBShaw/golang-graphql-server/graph/resolvers"
	customMiddleware "github.com/JamieBShaw/golang-graphql-server/middleware"
	"github.com/JamieBShaw/golang-graphql-server/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v10"
	"github.com/hashicorp/go-hclog"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {

	l := hclog.Default()

	db := postgres.NewDB(&pg.Options{

		User:     "postgres",
		Password: "postgres",
		Database: "meetmeup_dev",
	})

	userRepo := postgres.UsersRepo{DB: db, Log: l}
	meetupRepo := postgres.MeetupsRepo{DB: db, Log: l}

	defer db.Close()

	db.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	opts := cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true}
	// Middlwares
	router.Use(cors.New(opts).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(userRepo))

	d := domain.New(userRepo, meetupRepo)

	conf := generated.Config{
		Resolvers: &resolvers.Resolver{Domain: d}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(conf))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", resolvers.DataLoaderMiddleware(db, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
