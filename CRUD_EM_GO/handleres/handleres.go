package handleres

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/igorrm19/aprendendo_go/models"
)

type TaskHandle struct {
	DB *sql.DB
}

func NewTaskHandle(db *sql.DB) *TaskHandle {
	return &TaskHandle{DB: db}
}

func (taskHendle *TaskHandle) ReadTask(res http.ResponseWriter, req *http.Request) {

	rows, err := taskHendle.DB.Query("SELECT * FROM task")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task

		err := rows.Scan(&task.ID, &task.Description, &task.Status, &task.Title)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)
	}

	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(tasks); err != nil {
		http.error
	}

}
