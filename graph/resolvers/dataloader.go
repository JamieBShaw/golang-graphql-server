package resolvers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/JamieBShaw/golang-graphql-server/graph/models"
	"github.com/go-pg/pg/v10"
)

const userLoaderKey = "userloader"

// Dataloader middleware for n+1 issue
func DataLoaderMiddleware(db *pg.DB, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userLoader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*models.User, []error) {
				var users []*models.User
				// problem is slices dont key their order,
				// rerunning the dataloader can cause issues

				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
				if err != nil {
					return nil, []error{err}
				}
				// Therefore make an empty map of users
				u := make(map[string]*models.User, len(users))
				// Now loop over map and place user where index == userID
				for i, user := range users {
					fmt.Println("LOOP", "Index", i, "Userid", user.ID)
					u[user.ID] = user
				}

				// Now we need to make a slice of length ids as our return is a pointer user slice
				res := make([]*models.User, len(ids))
				// loop over range of ids, at each index i insert user
				for i, id := range ids {
					res[i] = u[id]
				}

				return res, nil
			},
		}
		ctx := context.WithValue(r.Context(), userLoaderKey, &userLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}
