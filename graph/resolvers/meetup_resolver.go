package resolvers

import (
	"context"
	"errors"

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
		r.MeetupsRepo.Log.Error("Could not retrieve user associated with meetup", err)
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input *models.UpdateMeetup) (*models.Meetup, error) {
	meetup, err := r.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return nil, err
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

	meetup, err = r.MeetupsRepo.Update(meetup)
	if err != nil {
		return nil, err
	}

	return meetup, nil
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

type meetupResolver struct{ *Resolver }
