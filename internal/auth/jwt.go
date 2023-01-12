package auth

import (
	"time"

	"github.com/adminsemy/URLShorting/internal/model"
	"github.com/golang-jwt/jwt/v4"
)

func MakeJWT(user model.User) (string, error) {
	claims := model.UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "adminsemy-org",
			IssuedAt: jwt.NewNumericDate(time.Now().UTC()),
		},
		User: user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("adf45b6969c54308d14818f8eb973eafc731522e"))
}
