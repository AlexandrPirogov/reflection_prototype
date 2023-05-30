package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/auth/jwt"
	"reflection_prototype/internal/core/auth/user"
)

const (
	QueryLogin = `select id from users
	where login = $1 and email = $2 and pwd = $3`

	QueryRegister = `insert into users values(default, $1, $2, $3, $4, $5, null, now())`
)

func (pg *pgConnection) Login(u user.User) (string, error) {
	var token string
	var id int
	err := pg.conn.QueryRow(context.Background(), QueryLogin, u.Login, u.Email, u.Pwd).Scan(&id)
	if err != nil {
		log.Println(err)
		return token, err
	}

	token = jwt.GenerateJWT(u)
	return token, nil
}

func (pg *pgConnection) Register(u user.User) error {
	_, err := pg.conn.Exec(context.Background(), QueryRegister, u.Login, u.Name, u.Surname, u.Email, u.Pwd)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
