package main

import (
	"log"
	"time"

	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/vegitobluefan/task-manager/repository"

	"github.com/gin-gonic/gin"

	"github.com/vegitobluefan/task-manager/api"
	"github.com/vegitobluefan/task-manager/dispatcher"
	"github.com/vegitobluefan/task-manager/domain"
	"github.com/vegitobluefan/task-manager/grpc"
	"github.com/vegitobluefan/task-manager/usecase"
)

func main() {
	db, err := connectToDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	repo := repository.NewPostgresRepo(db)

	handler := func(task *domain.Task) {
		time.Sleep(5 * time.Second)
		_ = repo.UpdateStatus(task.ID, "выполнен", "сон 5с")
	}

	d := dispatcher.NewDispatcher(4, handler)
	uc := usecase.NewTaskUseCase(repo, d)

	r := gin.Default()
	api.SetupRoutes(r, uc, repo)

	go func() {
		if err := grpc.RunGRPCServer(uc, ":50051"); err != nil {
			log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
		}
	}()

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

func connectToDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return sql.Open("postgres", dsn)
}
