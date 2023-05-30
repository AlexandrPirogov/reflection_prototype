package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/auth/user"
	"reflection_prototype/internal/core/contributes"
)

const (
	QueryCtbProcess = `insert into CONTRIBUTIONS values(default, 
		(select id from users where email = $1), $2, $3)`

	QueryCtbThread = `insert into CONTRIBUTIONS values(default,  
		(select id from users where email = $1), $2, $3)`

	QueryCtbQuant = `insert into CONTRIBUTIONS values(default,   
		(select id from users where email = $1), $2, $3)`
)

func (pg *pgConnection) contributeProcessCreation(u user.User) error {
	ctb := contributes.New(contributes.CREATE_PROCESS)

	_, err := pg.conn.Exec(context.Background(), QueryCtbProcess, u.Email, contributes.Type(ctb), contributes.Time(ctb))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (pg *pgConnection) contributeThreadCreation(u user.User) error {
	ctb := contributes.New(contributes.CREATE_THREAD)

	_, err := pg.conn.Exec(context.Background(), QueryCtbThread, u.Email, contributes.Type(ctb), contributes.Time(ctb))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (pg *pgConnection) contributeQantCreation(u user.User) error {
	ctb := contributes.New(contributes.CREATE_QUANT)

	_, err := pg.conn.Exec(context.Background(), QueryCtbQuant, u.Email, contributes.Type(ctb), contributes.Time(ctb))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
