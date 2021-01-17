package cronnie

import (
	"os"
	"os/signal"
	"time"
)

// Cronnie ignals a channel returned by the Done method
// It will signal after a duration or when the process has been signalled
type Cronnie interface {
	Start()
	Done()
	Wait() chan (struct{})
}

type cronnie struct {
	duration time.Duration

	waiter chan (os.Signal)
	work   chan (struct{})
	done   chan (struct{})
}

func (c *cronnie) Start() {
	go func() {
		for {
			select {
			case <-c.waiter:
				break
			case <-time.After(c.duration):
				break
			case <-c.done:
				return
			}

			c.work <- struct{}{}
		}
	}()
}

func (c *cronnie) Wait() chan (struct{}) {
	return c.work
}

func (c *cronnie) Done() {
	c.done <- struct{}{}
	close(c.work)
}

// NewCronnie constructs a Cronnie
// Optionally pass signals to listen for, in addition to the duration
//
// For example:
//
// 	c := cronnie.NewCronnie(250*time.Millisecond, syscall.SIGUSR1)
// 	c.Start()
// 	defer c.Done()
// 	for {
// 		<-c.Wait() // wait until SIGUSR1 or 250ms has elapsed
// 	}
//
func NewCronnie(duration time.Duration, signals ...os.Signal) Cronnie {
	c := &cronnie{
		duration: duration,
	}

	c.waiter = make(chan os.Signal)
	c.done = make(chan struct{})
	c.work = make(chan struct{})

	signal.Notify(c.waiter, signals...)

	return c
}
