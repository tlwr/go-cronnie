# `go-cronnie`

[![Go Reference](https://pkg.go.dev/badge/github.com/tlwr/go-cronnie.svg)](https://pkg.go.dev/github.com/tlwr/go-cronnie)

Very basic package for doing things periodically, or when signalled

```
// Example
c := cronnie.NewCronnie(250*time.Millisecond, syscall.SIGUSR1)

c.Start()
defer c.Done()

for {
	<-c.Wait() // wait until SIGUSR1 or 250ms has elapsed
}
```
