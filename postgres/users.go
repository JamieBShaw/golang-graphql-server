package postgres

import (
	"github.com/JamieBShaw/golang-graphql-server/graph/models"
	"github.com/go-pg/pg/v10"
	"github.com/hashicorp/go-hclog"
)

type UsersRepo struct {
	DB  *pg.DB
	Log hclog.Logger
}

func (u *UsersRepo) GetAll() ([]*models.User, error) {
	u.Log.Info("Getting users")

	var users []*models.User

	err := u.DB.Model(&users).Select()
	if err != nil {
		u.Log.Error("Could not retrieve users", err)
		return nil, err
	}

	return users, nil
}

func (u *UsersRepo) GetByID(id string) (*models.User, error) {
	u.Log.Info("Getting single user")

	var user models.User

	err := u.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		u.Log.Error("Could not retrieve users", err)
		return nil, err
	}
	return &user, nil
}
