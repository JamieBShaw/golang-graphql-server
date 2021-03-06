package resolvers

import (
	"context"
	"fmt"

	"github.com/JamieBShaw/golang-graphql-server/graph/generated"
	"github.com/JamieBShaw/golang-graphql-server/graph/models"
)

type userResolver struct{ *Resolver }

func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

func (r *userResolver) FirstName(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) LastName(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	return r.Domain.MeetupsRepo.GetMeetUpsForUser(obj)

}
