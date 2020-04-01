package tasks

import (
	"database/sql"
	"github.com/JasonSteinberg/timeTicker/database"
)

func getOpenTasks(db *sql.DB, userID int64) string {
	rows, err := db.Query("select id, name, due_date from task where user_id=? and is_completed != 1;", userID)

	if err != nil {
		return err.Error()
	}

	results, err := database.ReturnJson(rows)
	return results
}

func getTask(db *sql.DB, userID int64, taskID int) string {
	rows, err := db.Query("select id, name, due_date from task where user_id=? and id=?;", userID, taskID)

	if err != nil {
		return err.Error()
	}

	results, err := database.ReturnJson(rows)
	return results
}

func CreateTable() []string {
	var cmds = []string{`
create table task
(
	id int auto_increment,
	user_id int null,
	name varchar(6000) null,
	due_date date null,
	is_completed int(1) default 0 not null,
	constraint task_pk
		primary key (id),
	constraint task_user_id_fk
		foreign key (user_id) references user (id)
			on delete cascade
);`,
		`create index task_due_date_index
	on task (due_date);`,
		`create index task_is_completed_index
	on task (is_completed);`,
		`create index task_user_id_index
	on task (user_id);`}
	return cmds
}
