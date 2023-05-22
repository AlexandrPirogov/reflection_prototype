package work

import "time"

/*
CREATE TABLE WORK_SESSIONS(
    ID SERIAL PRIMARY KEY,
    USER_ID INT NOT NULL,
    DT_START TIMESTAMP NOT NULL,
    DT_END TIMESTAMP NOT NULL,
    CONSTRAINT fk_worksession_user FOREIGN KEY(USER_ID) REFERENCES USERS(ID)
);
*/

type work struct {
	dateTimeStart time.Time
	dateTimeEnd   time.Time
}

func New() work {
	return work{
		dateTimeStart: time.Now(),
	}
}

func Done(w work) work {
	w.dateTimeEnd = time.Now()
	return w
}
