//author gonghh
//Copyright 2013 gonghh(screscent). All rights reserved.
package coroutinepool

type coroutine struct {
  recv   chan func()
	idlech chan *coroutine
	exitch chan int
}

var idlech chan *coroutine
var poolSize int = 0

func Create(size int) {
	poolSize = size
	idlech = make(chan *coroutine, poolSize)
	for i := 0; i < poolSize; i++ {
		c := &coroutine{make(chan func()), idlech, make(chan int)}
		go c.run()
	}
}

func Run(f func()) {
	c := <-idlech
	c.recv <- f
}

func Exit() {
	for i := 0; i < poolSize; i++ {
		c := <-idlech
		c.exitch <- 1
	}
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
