package models

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type TaskRequest struct {
	Name string `json:"name"`
}
type TaskResponse struct {
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	LeadTimeMin int       `json:"lead_time_min"`
	EndTime     time.Time `json:"end_time"`
}

//type Task struct {
//	Name          string    `json:"name"`
//	Lead_time_min int       `json:"lead___time___min"`
//	Beg_time      time.Time `json:"beg_time"`
//	End_time      time.Time `json:"end_time"`
//}

type Status struct {
	Status string `json:"status"`
}
type TaskValue struct {
	LeadTimeMin int       `json:"lead_time"`
	BegTime     time.Time `json:"beg_time"`
	EndTime     time.Time `json:"end_time"`
}
type DataModel map[string]TaskValue

func NewDataModel() DataModel {
	data := DataModel{}
	return data
}

func (data DataModel) Add(task TaskRequest, mu *sync.Mutex) error {
	randomTime := 3 + rand.Intn(3)

	mu.Lock()
	defer mu.Unlock()

	if _, ok := data[task.Name]; ok {
		return errors.New("task already exists")
	}
	data[task.Name] = TaskValue{LeadTimeMin: randomTime,
		BegTime: time.Now(),
		EndTime: time.Now().Add(time.Duration(randomTime) * time.Minute)}

	fmt.Println(data)
	return nil
}

func (data DataModel) Get(name string, mu *sync.Mutex) (TaskValue, error) {
	mu.Lock()
	defer mu.Unlock()

	if val, ok := data[name]; ok {

		return val, nil
	}

	return TaskValue{}, errors.New("task not found")
}

func (data DataModel) ValuesResp(mu *sync.Mutex) []TaskResponse {
	mu.Lock()
	defer mu.Unlock()

	m := make([]TaskResponse, 0, len(data))
	stat := ""

	for nameTask, valTask := range data {
		stat = "in progress"
		if time.Now().After(valTask.EndTime) {
			stat = "completed"
		}
		m = append(m, TaskResponse{Name: nameTask,
			Status:      stat,
			LeadTimeMin: valTask.LeadTimeMin,
			EndTime:     valTask.EndTime})
	}

	return m
}

func (data DataModel) Delete(name string, mu *sync.Mutex) error {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := data[name]; !ok {
		return errors.New("task not found")
	}

	delete(data, name)

	return nil
}
