package usecase

import "github.com/vegitobluefan/task-manager/domain"

// Интерфейс юзкейса
type TaskUseCase interface {
	Enqueue(task *domain.Task) error
}

// Интерфейс диспетчера
type TaskDispatcher interface {
	Dispatch(task *domain.Task) error
}

// Реализация юзкейса
type taskUseCase struct {
	repo       domain.TaskRepository
	dispatcher TaskDispatcher
}

// Конструктор
func NewTaskUseCase(r domain.TaskRepository, d TaskDispatcher) TaskUseCase {
	return &taskUseCase{repo: r, dispatcher: d}
}

// Метод постановки задачи в очередь
func (u *taskUseCase) Enqueue(task *domain.Task) error {
	if err := u.repo.Save(task); err != nil {
		return err
	}
	return u.dispatcher.Dispatch(task)
}
