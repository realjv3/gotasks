package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/realjv3/gotasks/domain"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	taskService domain.TaskService
}

func NewTaskHandler(taskService domain.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if task.UserID == 0 || task.Title == "" {
		http.Error(w, "user ID and title are required", http.StatusBadRequest)
	}

	t, err := h.taskService.CreateTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

func (h *TaskHandler) getTasksByUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tasks, err := h.taskService.GetTasksByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

func (h *TaskHandler) finishTask(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(chi.URLParam(r, "taskID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	task, err := h.taskService.FinishTask(taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}
