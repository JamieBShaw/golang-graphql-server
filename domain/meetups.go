package domain

import (
	"context"
	"errors"

	"github.com/JamieBShaw/golang-graphql-server/graph/models"
	"github.com/JamieBShaw/golang-graphql-server/middleware"
)

func (d *Domain) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUserUnAuthenticated
	}

	if len(input.Name) < 3 || len(input.Description) < 10 {
		return nil, errors.New("Name and or description to short")
	}

	newMeetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      currentUser.ID,
	}

	res, err := d.MeetupsRepo.Create(newMeetup)

	if err != nil {
		return nil, err
	}

	return res, nil
}
func (d *Domain) UpdateMeetup(ctx context.Context, id string, input *models.UpdateMeetup) (*models.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUserUnAuthenticated
	}

	meetup, err := d.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return nil, err
	}

	if !checkOwnership(meetup, currentUser) {
		return false, ErrForbidden
	}

	didUpdate := false

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("Name is not long enough")
		}
		meetup.Name = *input.Name
		didUpdate = true
	}

	if input.Description != nil {
		if len(*input.Description) < 10 {
			return nil, errors.New("Description is not long enough")
		}
		meetup.Description = *input.Description
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("No valid input given to update")
	}

	meetup, err = d.MeetupsRepo.Update(meetup)
	if err != nil {
		return nil, err
	}

	return meetup, nil
}
func (d *Domain) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return false, ErrUserUnAuthenticated
	}
	meetup, err := d.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("Meetup not found")
	}

	if !checkOwnership(meetup, currentUser) {
		return false, ErrForbidden
	}

	err = d.MeetupsRepo.Delete(meetup)
	if err != nil {
		return false, errors.New("Meet up could not be deleted")
	}
	return true, nil
}
