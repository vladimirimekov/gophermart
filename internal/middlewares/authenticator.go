package middlewares

import (
	"context"
	"net/http"

	"github.com/vladimirimekov/gophermart/internal/storage"

	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwt"
)

type UserCookies struct {
	Storage storage.Repositories
	UserKey interface{}
}

func (h UserCookies) CheckUserCookies(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token, data, err := jwtauth.FromContext(r.Context())
		if err != nil || token == nil || jwt.Validate(token) != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		userID, ok := data["userID"].(float64)
		if !ok {
			http.Error(w, "wrong data in cookie", http.StatusInternalServerError)
			return
		}

		if ok = h.Storage.IsUserExistByUserID(int(userID), r.Context()); !ok {
			http.Error(w, "user doesn't exist", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), h.UserKey, int(userID))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})

}
