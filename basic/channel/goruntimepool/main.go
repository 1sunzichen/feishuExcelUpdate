package main

import (
	"errors"
	"sync/atomic"
)

// 协程池
// 任务的定义
// 任务要包含需要执行的函数、以及函数要传的参数, 因为参数类型、个数不确定, 这里使用可变参数和空接口的形式
type Task struct {
	Handler func(v ...interface{})
	Params  []interface{}
}

// 协程池的定义
type Pool struct {
	capacity       uint64
	runningWorkers uint64
	state          int64
	taskC          chan *Task
	closeC         chan bool
}

var ErrInvalidPoolCap = errors.New("invalid pool cap")

const (
	RUNNING = 1
	STOP    = 0
)

func NewPool(capacity uint64) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrInvalidPoolCap
	}
	return &Pool{
		capacity: capacity,
		state:    RUNNING,
		taskC:    make(chan *Task, capacity),
		closeC:   make(chan bool),
	}, nil
}

func (p *Pool) Run() {

	p.incRunning()
	go func() {

		defer func() {
			p.decRunning()
		}()
		for {
			select {
			case task, ok := <-p.taskC:
				if !ok {
					return
				}
				task.Handler(task.Params...)
			case <-p.closeC:
				return
			}
		}
	}()
}

func (p *Pool) incRunning() {
	atomic.AddUint64(&p.runningWorkers, 1)
}

func (p *Pool) decRunning() {
	atomic.AddUint64(&p.runningWorkers, ^uint64(0))
}

func (p *Pool) GetRunningWorkers() uint64 {
	return atomic.LoadUint64(&p.runningWorkers)
}

func (p *Pool) GetCap() uint64 {
	return atomic.LoadUint64(&p.capacity)
}

func (p *Pool) Put(task *Task) {
	if p.GetRunningWorkers() >= p.GetCap() {
		return
	}
	p.taskC <- task
}
