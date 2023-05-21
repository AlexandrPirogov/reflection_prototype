package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/contributes"
)

func (pg *pgConnection) contributeProcessCreation() error {
	ctb := contributes.New(contributes.CREATE_PROCESS)

	//#TOFIX user id
	query := `insert into CONTRIBUTIONS values(default, 1, $1, $2)`
	_, err := pg.conn.Exec(context.Background(), query, contributes.Type(ctb), contributes.Time(ctb))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (pg *pgConnection) contributeThreadCreation() error {
	ctb := contributes.New(contributes.CREATE_THREAD)

	//#TOFIX user id
	query := `insert into CONTRIBUTIONS values(default, 1, $1, $2)`
	_, err := pg.conn.Exec(context.Background(), query, contributes.Type(ctb), contributes.Time(ctb))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (pg *pgConnection) contributeQantCreation() error {
	ctb := contributes.New(contributes.CREATE_QUANT)

	//#TOFIX user id
	query := `insert into CONTRIBUTIONS values(default, 1, $1, $2)`
	_, err := pg.conn.Exec(context.Background(), query, contributes.Type(ctb), contributes.Time(ctb))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
