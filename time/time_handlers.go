package timeLog

import (
	"encoding/json"
	"fmt"
	"github.com/JasonSteinberg/timeTicker/database"
	"github.com/JasonSteinberg/timeTicker/responses"
	"github.com/JasonSteinberg/timeTicker/structs"
	"net/http"
)

func timeAdd(w http.ResponseWriter, r *http.Request) {
	var timeLog structs.TimeRequest
	var error structs.ServerMessage

	json.NewDecoder(r.Body).Decode(&timeLog)

	if timeLog.LoggedTime == "" {
		error.Message = "Time to log is missing."
		responses.ErrorResponder(w, http.StatusBadRequest, error)
		return
	}

	user := r.Context().Value("User").(structs.User)
	fmt.Fprintln(w, createTime(database.GetSqlWriteDB(), user, timeLog))
}
