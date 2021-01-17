package cronnie_test

import (
	"syscall"
	"time"

	"github.com/tlwr/go-cronnie"
)

func Example() {
	c := cronnie.NewCronnie(250*time.Millisecond, syscall.SIGUSR1)

	c.Start()
	defer c.Done()

	for {
		<-c.Wait() // wait until SIGUSR1 or 250ms has elapsed
	}
}
