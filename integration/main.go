package main

import (
	"fmt"
	"syscall"
	"time"

	"github.com/tlwr/go-cronnie"
)

func main() {
	c := cronnie.NewCronnie(250*time.Millisecond, syscall.SIGUSR1)

	c.Start()

	<-c.Wait()

	c.Done()
	fmt.Printf("done")
}
