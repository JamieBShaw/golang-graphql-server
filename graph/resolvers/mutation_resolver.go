package resolvers

import (
	"context"
	"errors"

	"github.com/JamieBShaw/golang-graphql-server/graph/generated"
	"github.com/JamieBShaw/golang-graphql-server/graph/models"
)

func (r *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3 || len(input.Description) < 10 {
		return nil, errors.New("Name and or description to short")
	}
	newMeetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      "2",
	}

	res, err := r.MeetupsRepo.Create(newMeetup)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := r.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("Meetup not found")
	}
	err = r.MeetupsRepo.Delete(meetup)
	if err != nil {
		return false, errors.New("Meet up could not be deleted")
	}
	return true, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
