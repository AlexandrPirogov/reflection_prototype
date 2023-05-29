package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/auth/user"
	"reflection_prototype/internal/core/contributes"
)

func (pg *pgConnection) contributeProcessCreation(u user.User) error {
	ctb := contributes.New(contributes.CREATE_PROCESS)

	//#TOFIX user id
	query := `insert into CONTRIBUTIONS values(default, 
		(select id from users where email = $1), $2, $3)`
	_, err := pg.conn.Exec(context.Background(), query, u.Email, contributes.Type(ctb), contributes.Time(ctb))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (pg *pgConnection) contributeThreadCreation(u user.User) error {
	ctb := contributes.New(contributes.CREATE_THREAD)

	//#TOFIX user id
	query := `insert into CONTRIBUTIONS values(default,  
		(select id from users where email = $1), $2, $3)`
	_, err := pg.conn.Exec(context.Background(), query, u.Email, contributes.Type(ctb), contributes.Time(ctb))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (pg *pgConnection) contributeQantCreation(u user.User) error {
	ctb := contributes.New(contributes.CREATE_QUANT)

	//#TOFIX user id
	query := `insert into CONTRIBUTIONS values(default,   
		(select id from users where email = $1), $2, $3)`
	_, err := pg.conn.Exec(context.Background(), query, u.Email, contributes.Type(ctb), contributes.Time(ctb))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
