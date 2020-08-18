package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/JamieBShaw/golang-graphql-server/graph/models"
	"github.com/JamieBShaw/golang-graphql-server/postgres"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/pkg/errors"
)

const CurrentUserKey = "currentUser"

func AuthMiddleware(repo postgres.UsersRepo) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token, err := parseToken(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				user, err := repo.GetByID(claims["jti"].(string))
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}
				ctx := context.WithValue(r.Context(), CurrentUserKey, user)
				next.ServeHTTP(w, r.WithContext(ctx))

			} else {
				next.ServeHTTP(w, r)
				return
			}
		})
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromToken,
}

// Bearer tokenstring121323421432542
func stripBearerPrefixFromToken(token string) (string, error) {
	bearer := "BEARER"
	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}
	return token, nil
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(os.Getenv("JWT_SECRET"))
		return t, nil
	})
	return jwtToken, errors.Wrap(err, "parseToken error: ")
}

func GetCurrentUserFromCTX(ctx context.Context) (*models.User, error) {

	if ctx.Value(CurrentUserKey) == nil {
		return nil, errors.New("No user in context")
	}
	user, ok := ctx.Value(CurrentUserKey).(*models.User)

	if !ok || user.ID == "" {
		return nil, errors.New("No user in context")
	}
	return user, nil
}
