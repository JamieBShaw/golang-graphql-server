package resolvers

import (
	"context"
	"errors"

	val "github.com/JamieBShaw/golang-graphql-server/graph"
	"github.com/JamieBShaw/golang-graphql-server/graph/generated"
	"github.com/JamieBShaw/golang-graphql-server/graph/models"
)

var (
	ErrInput = errors.New("Input error")
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

func (m *mutationResolver) LoginUser(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	isValid := val.Validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}
	return m.Domain.LoginUser(ctx, input)
}

func (m *mutationResolver) RegisterUser(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	isValid := val.Validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}
	return m.Domain.RegisterUser(ctx, input)
}

func (m *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	return m.Domain.CreateMeetup(ctx, input)
}

func (m *mutationResolver) UpdateMeetup(ctx context.Context, id string, input *models.UpdateMeetup) (*models.Meetup, error) {
	return m.Domain.UpdateMeetup(ctx, id, input)
}
func (m *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	return m.Domain.DeleteMeetup(ctx, id)
}
