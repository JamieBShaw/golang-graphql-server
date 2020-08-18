package resolvers

import (
	"context"

	"github.com/JamieBShaw/golang-graphql-server/graph/generated"
	"github.com/JamieBShaw/golang-graphql-server/graph/models"
)

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

func (q *queryResolver) Meetups(ctx context.Context, filter *models.MeetupFilter, limit *int, offset *int) ([]*models.Meetup, error) {
	return q.Domain.MeetupsRepo.GetMeetups(filter, limit, offset)
}
func (q *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return q.Domain.UsersRepo.GetByID(id)
}

// Query returns generated.QueryResolver implementation.
