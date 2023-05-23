package jwt

import (
	"fmt"
	"reflection_prototype/internal/core/auth/user"

	"github.com/go-chi/jwtauth/v5"
)

var authedUsers = make(map[string]user.User, 0)
var TokenAuth *jwtauth.JWTAuth = jwtauth.New("HS256", []byte("secret"), nil)

func GenerateJWT(u user.User) string {

	key := u.Login + u.Pwd + u.Email
	claims := map[string]interface{}{key: u}
	//jwtauth.SetExpiry(claims, time.Now().Add(time.Second*30))
	_, tokenString, _ := TokenAuth.Encode(claims)
	authedUsers[tokenString] = u
	return tokenString
}

func UserFromToken(token string) (user.User, error) {
	if u, ok := authedUsers[token]; ok {
		return u, nil
	}

	return user.User{}, fmt.Errorf("not authorized")
}
