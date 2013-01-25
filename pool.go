//author gonghh
//Copyright 2013 gonghh(screscent). All rights reserved.
package coroutinepool

type coroutine struct {
	recv   chan func()
	idlech chan *coroutine
	exitch chan int
}

type Pool struct {
	idlech   chan *coroutine
	poolSize int
}

var default_pool *Pool

func NewPool() *Pool {
	return &Pool{make(chan *coroutine), 0}
}

func (p *Pool) Create(size int) {
	p.poolSize = size
	for i := 0; i < p.poolSize; i++ {
		c := &coroutine{make(chan func()), p.idlech, make(chan int)}
		go c.run()
	}
}

func Create(size int) {
	default_pool := NewPool()
	default_pool.Create(size)
}

func (p *Pool) Run(f func()) {
	c := <-p.idlech
	c.recv <- f
}

func Run(f func()) {
	c := <-default_pool.idlech
	c.recv <- f
}

func (p *Pool) Exit() {
	for i := 0; i < p.poolSize; i++ {
		c := <-p.idlech
		c.exitch <- 1
	}
}

func Exit() {
	default_pool.Exit()
}

func (c *coroutine) run() {

	for {
		c.idlech <- c
		select {
		case f := <-c.recv:
			f()
		case <-c.exitch:
			return
		}
	}

}
