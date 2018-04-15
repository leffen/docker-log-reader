package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	d, err := NewDlog()
	if err != nil {
		panic(err)
	}

	d.Run()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

}
