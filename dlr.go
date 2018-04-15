package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {
	d, err := NewDlog()
	if err != nil {
		panic(err)
	}

	d.list()
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

}

// Dlog collection of docker handling
type Dlog struct {
	cli        *client.Client
	containers []types.Container
}

// NewDlog creates a new Docker container log object
func NewDlog() (*Dlog, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &Dlog{cli: cli}, nil
}

func (d *Dlog) list() error {
	cc, err := d.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return err
	}
	d.containers = cc

	for _, c := range cc {
		fmt.Printf("%-3.3s - %-30.30v %s\n", c.ID, c.Names, c.State)
		go d.dumpLog(c.ID)
	}
	return nil
}

func (d *Dlog) dumpLog(id string) {
	ctx := context.Background()

	options := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: true,
	}

	// Replace this ID with a container that really exists
	out, err := d.cli.ContainerLogs(ctx, id, options)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(out)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%4.4s STOP reading log on container %v\n", id, err)
			out.Close()
			return
		}
		fmt.Printf("%4.4s %s\n", id, line)
	}

}
