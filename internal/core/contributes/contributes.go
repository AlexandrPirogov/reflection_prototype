package contributes

import "time"

/*
CREATE TABLE CONTRIBUTIONS(
    ID SERIAL PRIMARY KEY,
    USER_ID INT,
    TYPE CONTRIBUTES_TYPE,
    CREATED_AT TIMESTAMP,
    CONSTRAINT fk_contr_user FOREIGN KEY(USER_ID) REFERENCES USERS(ID)
);
*/

type constraintType int

const (
	CREATE_PROCESS constraintType = iota
	CREATE_THREAD
	CREATE_QUANT
)

var contributeTypeToString = map[constraintType]string{
	CREATE_PROCESS: "create_process",
	CREATE_THREAD:  "create_thread",
	CREATE_QUANT:   "create_quant",
}

type contribute struct {
	ctype     constraintType
	createdAt time.Time
}

func New(ctype constraintType) contribute {
	return contribute{
		ctype:     ctype,
		createdAt: time.Now(),
	}
}

func Time(c contribute) time.Time {
	return c.createdAt
}

func Type(c contribute) string {
	return contributeTypeToString[c.ctype]
}
