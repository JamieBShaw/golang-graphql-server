package domain

import (
	"context"
	"errors"
	"log"

	"github.com/JamieBShaw/golang-graphql-server/graph/models"
)

func (d *Domain) LoginUser(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	if *input.Username == "" && *input.Email == "" {
		return nil, ErrBadCredentials
	}

	user, err := d.UsersRepo.GetUserByEmail(*input.Email)
	if err != nil {
		return nil, ErrBadCredentials
	}

	err = user.ValidatePassword(input.Password)
	if err != nil {
		return nil, ErrBadCredentials
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, ErrGeneric
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil

}

func (d *Domain) RegisterUser(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	_, err := d.UsersRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("Email already in use")
	}
	_, err = d.UsersRepo.GetUserByUsername(input.Username)
	if err == nil {
		return nil, errors.New("Username already in use")
	}

	user := &models.User{
		Email:     input.Email,
		Username:  input.Username,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("Erorr while hashing password %v", err)
		return nil, errors.New("something went wrong")
	}

	// TODO: Create verification code

	tx, err := d.UsersRepo.DB.Begin()
	if err != nil {
		log.Printf("error creating transaction: %v", err)
		return nil, errors.New("something went wrong")
	}
	defer tx.Rollback()

	if _, err := d.UsersRepo.CreateUser(tx, user); err != nil {
		log.Printf("error creating user: %v", err)
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		log.Panicf("error committing transaction: %v", err)
		return nil, err
	}

	token, err := user.GenerateToken()
	if err != nil {
		log.Panicf("error generating token: %v", err)
		return nil, errors.New("something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil

}
