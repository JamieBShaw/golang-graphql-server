package resolvers

import "github.com/JamieBShaw/golang-graphql-server/postgres"

//go:generate go run github.com/99designs

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UsersRepo   postgres.UsersRepo
	MeetupsRepo postgres.MeetupsRepo
}
