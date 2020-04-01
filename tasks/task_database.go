package tasks

import (
	"database/sql"
	"github.com/JasonSteinberg/timeTicker/database"
	"github.com/JasonSteinberg/timeTicker/structs"
)

func getOpenTasks(db *sql.DB, userID int64) string {
	rows, err := db.Query("select id, name, due_date from task where user_id=? and is_completed != 1;", userID)
	defer rows.Close()

	if err != nil {
		return err.Error()
	}

	results, err := database.ReturnJson(rows)
	return results
}

func getTask(db *sql.DB, userID int64, taskID int) string {
	rows, err := db.Query("select id, name, due_date from task where user_id=? and id=?;", userID, taskID)
	defer rows.Close()

	if err != nil {
		return err.Error()
	}

	results, err := database.ReturnJson(rows)
	return results
}

func getCompletedTasks(db *sql.DB, userID int64) string {
	rows, err := db.Query("select id, name, due_date from task where user_id=? and is_completed = 1;", userID)
	defer rows.Close()

	if err != nil {
		return err.Error()
	}

	results, err := database.ReturnJson(rows)
	if err != nil {
		return err.Error()
	}
	return results
}

func createTask(db *sql.DB, user structs.User, task structs.Task) string {

	var completed = 0
	if task.IsCompleted {
		completed = 1
	}

	_, err := db.Exec("insert into task (user_id, name, due_date, is_completed) values (?,?,?,?);",
		user.ID, database.NewNullString(task.Name), database.NewNullDate(task.DueDate), completed)

	if err != nil {
		return err.Error()
	}

	return ""
}

func updateTask(db *sql.DB, user structs.User, task structs.Task) string {

	var completed = 0
	if task.IsCompleted {
		completed = 1
	}

	_, err := db.Exec(`update task 
    						  set name=?, due_date=?, is_completed=?
							  where user_id=? and id=?;`,
		database.NewNullString(task.Name), database.NewNullDate(task.DueDate), completed, user.ID, task.ID)

	if err != nil {
		return err.Error()
	}

	return ""
}

func deleteTask(db *sql.DB, userID int64, taskID int) string {
	_, err := db.Exec("delete from task where user_id=? and id=?;", userID, taskID)

	if err != nil {
		return err.Error()
	}

	return ""
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
