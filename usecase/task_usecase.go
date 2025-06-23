package usecase

import (
	"errors"

	"github.com/vegitobluefan/task-manager/dispatcher"
	"github.com/vegitobluefan/task-manager/domain"
)

type TaskUseCase interface {
	Enqueue(task *domain.Task) (string, error)
	GetTask(id string) (*domain.Task, error)
	ListTasks() ([]*domain.Task, error)
}

type taskUseCase struct {
	repo       domain.TaskRepository
	dispatcher *dispatcher.Dispatcher
}

func NewTaskUseCase(repo domain.TaskRepository, d *dispatcher.Dispatcher) TaskUseCase {
	return &taskUseCase{
		repo:       repo,
		dispatcher: d,
	}
}

func (uc *taskUseCase) Enqueue(task *domain.Task) (string, error) {
	err := uc.repo.Save(task)
	if err != nil {
		return "", err
	}
	uc.dispatcher.Enqueue(task)
	return task.ID, nil
}

func (uc *taskUseCase) GetTask(id string) (*domain.Task, error) {
	task, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, errors.New("задача не найдена")
	}
	return task, nil
}

func (uc *taskUseCase) ListTasks() ([]*domain.Task, error) {
	return uc.repo.ListTasks()
}
