package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/TimRobillard/handicap_tracker/store"
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

type ContextKey string

const (
	ContextId ContextKey = "userId"
)

func GetUserFromRequest(r *http.Request, us store.UserStore) (*store.UIUser, error) {
	userId, ok := r.Context().Value(ContextId).(int)
	if !ok {
		return nil, fmt.Errorf("no user id found in context")
	}
	return us.GetUIUserById(r.Context(), userId)
}

func AddUserToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("_q")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		token, err := jwt.ParseWithClaims(
			cookie.Value,
			&JwtCustomClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			next.ServeHTTP(w, r)
			return
		}

		claims := token.Claims.(*JwtCustomClaims)

		c := context.WithValue(r.Context(), ContextId, claims.Id)
		next.ServeHTTP(w, r.WithContext(c))
	})
}

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(ContextId)
		if userId == nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GenerateToken(id int) (string, error) {
	claims := &JwtCustomClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
