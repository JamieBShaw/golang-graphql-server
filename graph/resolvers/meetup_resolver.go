package resolvers

import (
	"context"
	"log"

	"github.com/JamieBShaw/golang-graphql-server/graph/generated"
	"github.com/JamieBShaw/golang-graphql-server/graph/models"
)

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	//u, err := r.UsersRepo.GetUserByID(obj.UserID)
	//if err != nil {
	//r.UsersRepo.Log.Error("Error receiving user", err)
	//return nil, err
	//}
	//return u, nil
	// Now we have set up our UserLoader for our meetupResolver we now longer have the call the above
	// Just need to enact the userloader here which will load in the user with the specific userkey
	user, err := getUserLoader(ctx).Load(obj.UserID)
	if err != nil {
		log.Printf("Error, userLoader", err)
		return nil, err
	}

	return user, nil
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

type meetupResolver struct{ *Resolver }
