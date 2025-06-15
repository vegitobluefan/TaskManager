package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/vegitobluefan/task-manager/api"
	"github.com/vegitobluefan/task-manager/dispatcher"
	"github.com/vegitobluefan/task-manager/domain"
	"github.com/vegitobluefan/task-manager/usecase"
)

func main() {
	memRepo := &InMemoryRepo{tasks: map[string]*domain.Task{}}

	handler := func(task *domain.Task) {
		time.Sleep(time.Second * 5)
		_ = memRepo.UpdateStatus(task.ID, "done", "slept 5s")
	}

	d := dispatcher.NewDispatcher(4, handler)
	uc := usecase.NewTaskUseCase(memRepo, d)
	r := gin.Default()
	api.SetupRoutes(r, uc, memRepo)

	log.Println("Запуск сервера на порту :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

type InMemoryRepo struct {
	tasks map[string]*domain.Task
}

func (r *InMemoryRepo) Save(t *domain.Task) error {
	r.tasks[t.ID] = t
	return nil
}

func (r *InMemoryRepo) UpdateStatus(id, status, result string) error {
	t, ok := r.tasks[id]
	if !ok {
		return nil
	}
	t.Status = status
	t.Result = result
	return nil
}

func (r *InMemoryRepo) GetByID(id string) (*domain.Task, error) {
	return r.tasks[id], nil
}

func (r *InMemoryRepo) List() ([]*domain.Task, error) {
	tasks := make([]*domain.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}
