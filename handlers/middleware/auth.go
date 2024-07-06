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
	return us.GetUIUserById(userId)
}

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("_q")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		token, err := jwt.ParseWithClaims(
			cookie.Value,
			&JwtCustomClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)
		fmt.Printf("Valid Token %t", token.Valid)
		if err != nil || !token.Valid {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		claims := token.Claims.(*JwtCustomClaims)

		c := context.WithValue(r.Context(), ContextId, claims.Id)
		next.ServeHTTP(w, r.WithContext(c))
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
	fmt.Println(token)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
