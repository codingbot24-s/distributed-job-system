package worker

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

type Worker struct {
	Rc *redis.Client
	// procceser interface
	PI     *Processor
	Stopch chan string
	Wg     sync.WaitGroup
	Flag   bool
}

func (w *Worker) Start()     {}
func (w *Worker) Stop()      {}
func (w *Worker) Poll()      {}
func (w *Worker) HandleJob() {}
func (w *Worker) Register()  {}
