package tasks

import (
	"database/sql"
	"fmt"
	"github.com/JasonSteinberg/timeTicker/database"
	"github.com/JasonSteinberg/timeTicker/middleware"
	"github.com/JasonSteinberg/timeTicker/structs"
	"github.com/JasonSteinberg/timeTicker/users"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func SetUpTaskRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", middleware.ProtectedMiddleWare(tasks)).Methods("GET")
	router.HandleFunc("/task/{id}", middleware.ProtectedMiddleWare(task)).Methods("GET")
}

// curl -X GET --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNvdW50X2Rvb2t1IiwiaXNzIjoiY291cnNlIn0.osrQe3VwnTGqjuhHg36R9DRDt5apXSqb5-5CltMdp6g" http://localhost:8808/tasks
func tasks(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, users.USERKEY).(structs.User)
	fmt.Fprintln(w, getOpenTasks(database.GetSqlReadDB(), user.ID))
}

// curl -X GET --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNvdW50X2Rvb2t1IiwiaXNzIjoiY291cnNlIn0.osrQe3VwnTGqjuhHg36R9DRDt5apXSqb5-5CltMdp6g" http://localhost:8808/task/2
func task(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	user := context.Get(r, users.USERKEY).(structs.User)
	fmt.Fprintln(w, getTask(database.GetSqlReadDB(), user.ID, id))
}

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
