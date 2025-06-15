package dispatcher

import (
	"log"

	"github.com/vegitobluefan/task-manager/domain"
)

type Dispatcher struct {
	queue   chan *domain.Task
	workers int
}

type HandlerFunc func(task *domain.Task)

func NewDispatcher(workers int, handler HandlerFunc) *Dispatcher {
	d := &Dispatcher{
		queue:   make(chan *domain.Task, 100),
		workers: workers,
	}

	for i := 0; i < workers; i++ {
		go func(id int) {
			for task := range d.queue {
				log.Printf("[Worker %d] Processing task %s", id, task.ID)
				handler(task)
			}
		}(i)
	}

	return d
}

func (d *Dispatcher) Dispatch(task *domain.Task) error {
	d.queue <- task
	return nil
}
