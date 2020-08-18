package domain

import (
	"errors"

	"github.com/JamieBShaw/golang-graphql-server/graph/models"
	"github.com/JamieBShaw/golang-graphql-server/postgres"
)

var (
	ErrBadCredentials      = errors.New("Email/password combintation is incorrect")
	ErrUserUnAuthenticated = errors.New("User not authenticated")
	ErrGeneric             = errors.New("Something went wrong")
	ErrForbidden           = errors.New("unauthorized")
)

type Domain struct {
	UsersRepo   postgres.UsersRepo
	MeetupsRepo postgres.MeetupsRepo
}

func New(usersRepo postgres.UsersRepo, meetupsRepo postgres.MeetupsRepo) *Domain {
	return &Domain{
		UsersRepo:   usersRepo,
		MeetupsRepo: meetupsRepo,
	}
}

type Ownable interface {
	IsOwner(user *models.User) bool
}

func checkOwnership(o Ownable, user *models.User) bool {
	return o.IsOwner(user)
}
