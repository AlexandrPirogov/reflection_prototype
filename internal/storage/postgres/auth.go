package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/config/env"
	"reflection_prototype/internal/core/auth/jwt"
	"reflection_prototype/internal/core/auth/user"

	"github.com/jackc/pgx/v5"
)

func (pg *PgConnection) Login(u user.User) (string, error) {
	var jwt string
	tmpUrl := env.ReadPgUrl()
	conn, err := pgx.Connect(context.Background(), tmpUrl)
	if err != nil {
		return jwt, err
	}
	defer conn.Close(context.Background())

	query := `select jwt from users
	where login = $1 and email = $2 and pwd = $3`
	err = conn.QueryRow(context.Background(), query, u.Login, u.Email, u.Pwd).Scan(&jwt)
	if err != nil {
		log.Println(err)
		return jwt, err
	}

	return jwt, nil
}

func (pg *PgConnection) Register(u user.User) error {
	tmpUrl := env.ReadPgUrl()
	conn, err := pgx.Connect(context.Background(), tmpUrl)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	/*
		ID SERIAL PRIMARY KEY,
		LOGIN VARCHAR(255),
		NAME VARCHAR(255),
		SURNAME VARCHAR(255),
		EMAIL VARCHAR(300),
		PWD VARCHAR(300),
		PHOTO VARCHAR(300),
		JWT VARCHAR(300),
		REGISTERED_AT TIMESTAMP
	*/
	query := "insert into users values(default, $1, $2, $3, $4, $5, null, $6,now())"
	jwt := jwt.GenerateJWT(u)
	_, err = conn.Exec(context.Background(), query, u.Login, u.Name, u.Surname, u.Email, u.Pwd, jwt)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
