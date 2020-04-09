package timeLog

import (
	"database/sql"
	"github.com/JasonSteinberg/timeTicker/database"
	"github.com/JasonSteinberg/timeTicker/structs"
)

func createTime(db *sql.DB, user structs.User, timeLog structs.TimeRequest) string {

	_, err := db.Exec("insert into time (user_id, time_log, happened_date) values (?,?,?);",
		user.ID, timeLog.LoggedTime, database.NewNullDate(timeLog.HappenedDate))

	if err != nil {
		return err.Error()
	}

	return ""
}

func CreateTable() []string {
	var cmds = []string{`
create table time
(
	id int auto_increment,
	user_id int null,
	time_log float(4,2) not null,
	happened_date date null,
	constraint time_pk
		primary key (id),
	constraint User
		foreign key (user_id) references user (id)
			on delete cascade
);`}
	return cmds
}
