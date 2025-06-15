package services

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log/slog"
	"my_tastAPI/internal/models"
	"net/http"
	"sync"
	"time"
)

func HomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(models.Status{Status: "connect"})
	}
}

func GetTaskInfo(data models.DataModel, mu *sync.Mutex, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nameTask := chi.URLParam(r, "nameTask")
		value, err := data.Get(nameTask, mu)
		if err != nil {
			log.Error(fmt.Sprintf("GetTaskInfo err: %s", err.Error()))
			_ = json.NewEncoder(w).Encode(models.Status{Status: err.Error()})
			return
		}
		stat := "in progress"
		if time.Now().After(value.EndTime) {
			stat = "completed"
		}
		log.Info(fmt.Sprintf("GetTaskInfo value: %v", value))
		_ = json.NewEncoder(w).Encode(models.TaskResponse{Name: nameTask, LeadTimeMin: value.LeadTimeMin, Status: stat, EndTime: value.EndTime})
	}
}

func AddTask(data models.DataModel, mu *sync.Mutex, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.TaskRequest
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			log.Error(fmt.Sprintf("AddTask err: %s", err.Error()))
			return
		}
		err = data.Add(task, mu)
		if err != nil {
			log.Error(fmt.Sprintf("AddTask err: %s", err.Error()))
			_ = json.NewEncoder(w).Encode(models.Status{Status: err.Error()})
			return
		}

		log.Info(fmt.Sprintf("AddTask value: %s", task))
		_ = json.NewEncoder(w).Encode(models.Status{Status: "add task"})

	}
}

func DeleteTask(data models.DataModel, mu *sync.Mutex, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.TaskRequest
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			log.Error(fmt.Sprintf("DeleteTask err: %s", err.Error()))
		}
		err = data.Delete(task.Name, mu)
		if err != nil {
			log.Error(fmt.Sprintf("DeleteTask err: %s", err.Error()))
			_ = json.NewEncoder(w).Encode(models.Status{Status: err.Error()})
			return
		}
		log.Info(fmt.Sprintf("DeleteTask value: %s", task))
		_ = json.NewEncoder(w).Encode(models.Status{Status: "delete task"})
	}
}

func GetAllTasks(data models.DataModel, mu *sync.Mutex, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info(fmt.Sprintf("GetAllTasks value: %v", data))
		_ = json.NewEncoder(w).Encode(data.ValuesResp(mu))
	}
}
