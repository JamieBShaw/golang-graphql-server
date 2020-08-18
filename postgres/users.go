package postgres

import (
	"fmt"

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

	return users, err
}

func (u *UsersRepo) GetUserByField(field, value string) (*models.User, error) {
	u.Log.Info("Getting user by field: ", field, " value: ", value)

	var user models.User
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()

	return &user, err
}

func (u *UsersRepo) GetByID(id string) (*models.User, error) {
	return u.GetUserByField("id", id)
}

func (u *UsersRepo) GetUserByEmail(email string) (*models.User, error) {
	return u.GetUserByField("email", email)
}
func (u *UsersRepo) GetUserByUsername(username string) (*models.User, error) {
	return u.GetUserByField("username", username)
}

func (u *UsersRepo) CreateUser(tx *pg.Tx, user *models.User) (*models.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()
	return user, err
}
