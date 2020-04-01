package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/JasonSteinberg/timeTicker/database"
	"github.com/JasonSteinberg/timeTicker/responses"
	"github.com/JasonSteinberg/timeTicker/structs"
	"github.com/JasonSteinberg/timeTicker/users"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

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

// curl -X GET --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNvdW50X2Rvb2t1IiwiaXNzIjoiY291cnNlIn0.osrQe3VwnTGqjuhHg36R9DRDt5apXSqb5-5CltMdp6g" -X POST http://localhost:8808/task/completed
func taskCompleted(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, users.USERKEY).(structs.User)
	fmt.Fprintln(w, getCompletedTasks(database.GetSqlReadDB(), user.ID))
}

// curl -X GET --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNvdW50X2Rvb2t1IiwiaXNzIjoiY291cnNlIn0.osrQe3VwnTGqjuhHg36R9DRDt5apXSqb5-5CltMdp6g" -X POST -d '{"name":"I am the 2nd greatest", "due_date":"2020-04-01T12:42:31Z"}' http://localhost:8808/task/
func taskNew(w http.ResponseWriter, r *http.Request) {
	var task structs.Task
	var error structs.ServerMessage

	json.NewDecoder(r.Body).Decode(&task)

	if task.Name == "" {
		error.Message = "Name is missing."
		responses.ErrorResponder(w, http.StatusBadRequest, error)
		return
	}

	user := context.Get(r, users.USERKEY).(structs.User)
	fmt.Fprintln(w, createTask(database.GetSqlWriteDB(), user, task))
}
