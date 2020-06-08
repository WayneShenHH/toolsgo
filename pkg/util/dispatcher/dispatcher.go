package dispatcher

import (
	"sync"
)

// A Dispatcher dispatches tasks to workers.
type Dispatcher struct {
	pool    chan *worker
	queue   chan Tasker
	workers []*worker
	wg      sync.WaitGroup
	quit    chan struct{}
}

// New creates and initializes a new Dispatcher.
// The Dispatcher has count Workers and size queue.
func New(count, size int) *Dispatcher {
	d := &Dispatcher{
		pool:  make(chan *worker, count),
		queue: make(chan Tasker, size),
		quit:  make(chan struct{}),
	}

	d.workers = make([]*worker, cap(d.pool))
	for i := 0; i < len(d.workers); i++ {
		d.workers[i] = &worker{
			dispatcher: d,
			task:       make(chan Tasker),
			quit:       make(chan struct{}),
		}
	}

	return d
}

// Start starts to dispatch tasks.
func (d *Dispatcher) Start() {
	for _, w := range d.workers {
		w.start()
	}

	go func() {
		for {
			select {
			case t := <-d.queue:
				(<-d.pool).task <- t
			case <-d.quit:
				return
			}
		}
	}()
}

// Enqueue enqueues a new task to the dispatcher.
func (d *Dispatcher) Enqueue(t Tasker) {
	d.wg.Add(1)
	d.queue <- t
}

// Wait waits until all tasks are completed.
func (d *Dispatcher) Wait() {
	d.wg.Wait()
}

type worker struct {
	dispatcher *Dispatcher
	task       chan Tasker
	quit       chan struct{}
}

func (w *worker) start() {
	go func() {
		for {
			w.dispatcher.pool <- w

			select {
			case t := <-w.task:
				t.Run()

				w.dispatcher.wg.Done()
			}
		}
	}()
}

// Tasker is the interface that wraps the Run method.
type Tasker interface {
	Run()
}

// The TaskerFunc type is an adapter to allow the use of ordinary functions as task runners.
// If f is a function with the appropriate signature, TaskerFunc(f) is a Tasker that calls f.
type TaskerFunc func()

// Run calls f().
func (f TaskerFunc) Run() {
	f()
}
