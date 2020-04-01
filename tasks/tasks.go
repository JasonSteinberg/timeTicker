package tasks

import (
	"fmt"
	"github.com/JasonSteinberg/timeTicker/database"
	"github.com/JasonSteinberg/timeTicker/middleware"
	"github.com/JasonSteinberg/timeTicker/structs"
	"github.com/JasonSteinberg/timeTicker/users"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
)

func SetUpTaskRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", middleware.ProtectedMiddleWare(tasks)).Methods("GET")
}

// curl -X GET --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNvdW50X2Rvb2t1IiwiaXNzIjoiY291cnNlIn0.osrQe3VwnTGqjuhHg36R9DRDt5apXSqb5-5CltMdp6g" http://localhost:8808/tasks
func tasks(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, users.USERKEY).(structs.User)
	fmt.Fprintln(w, getOpenTasks(user.ID))
}

// curl -X GET --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNvdW50X2Rvb2t1IiwiaXNzIjoiY291cnNlIn0.osrQe3VwnTGqjuhHg36R9DRDt5apXSqb5-5CltMdp6g" http://localhost:8808/tasks
func getOpenTasks(userID int64) string {
	db := database.GetSqlReadDB()

	rows, err := db.Query("select name, due_date from task where user_id=? and is_completed != 1;", userID)

	if err != nil {
		return err.Error()
	}

	results, err := database.ReturnJson(rows)
	return results
}
