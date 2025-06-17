package dispatcher

import (
	"log"
	"sync"

	"github.com/vegitobluefan/task-manager/domain"
)

type Dispatcher struct {
	workerCount int
	queue       chan *domain.Task
	handler     func(task *domain.Task)
	wg          sync.WaitGroup
}

func NewDispatcher(workerCount int, handler func(task *domain.Task)) *Dispatcher {
	d := &Dispatcher{
		workerCount: workerCount,
		queue:       make(chan *domain.Task, 100),
		handler:     handler,
	}

	d.start()
	return d
}

func (d *Dispatcher) start() {
	for i := 0; i < d.workerCount; i++ {
		d.wg.Add(1)
		go func(workerID int) {
			defer d.wg.Done()
			for task := range d.queue {
				log.Printf("[worker-%d] processing task: %s\n", workerID, task.ID)
				d.handler(task)
			}
		}(i)
	}
}

// 🔧 Вот этот метод тебе и нужен:
func (d *Dispatcher) Enqueue(task *domain.Task) {
	d.queue <- task
}

func (d *Dispatcher) Stop() {
	close(d.queue)
	d.wg.Wait()
}
