package jwt

import (
	"reflection_prototype/internal/core/auth/user"

	"github.com/go-chi/jwtauth/v5"
)

var claims = make(map[string]interface{}, 0)
var TokenAuth *jwtauth.JWTAuth = jwtauth.New("HS256", []byte("secret"), nil)

func GenerateJWT(u user.User) string {

	key := u.Login + u.Pwd + u.Email
	claims[key] = u
	//jwtauth.SetExpiry(claims, time.Now().Add(time.Second*30))
	_, tokenString, _ := TokenAuth.Encode(claims)

	return tokenString
}
