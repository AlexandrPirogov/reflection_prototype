package jwt

import (
	"reflection_prototype/internal/core/auth/user"

	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth = jwtauth.New("HS256", []byte("secret"), nil)

func GenerateJWT(u user.User) string {
	key := u.Login + u.Pwd + u.Email
	_, tokenString, _ := TokenAuth.Encode(map[string]interface{}{key: u})
	return tokenString
}
