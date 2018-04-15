package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

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

// Run lists containers and starts listening on logfile changes
func (d *Dlog) Run() error {
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

func (d *Dlog) dumpLog(id string) error {
	ctx := context.Background()

	options := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: true,
		Details:    true,
	}

	info, err := d.cli.ContainerInspect(ctx, id)

	spew.Dump(info)

	out, err := d.cli.ContainerLogs(ctx, id, options)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(out)

	for {
		if scanner.Scan() {
			b := logRecord(scanner.Bytes())

			fmt.Printf("%4.4s %v\n", id, b)

		}
		err = scanner.Err()
		if err != nil {

			if err != io.EOF {
				out.Close()
				return err
			}
			// EOF. Sleep a bit and check again
			<-time.After(500 * time.Millisecond)
		}
	}

}

type DockerLog struct {
	Type      string            `json:"Type"`
	Labels    map[string]string `json:"Labels"`
	Timestamp string            `json:"Timestamp"`
	IP        string            `json:"IP"`
	Message   string            `json:"Message"`
}

func logRecord(b []byte) *DockerLog {
	// It is encoded on the first 8 bytes like this:
	//
	// header := [8]byte{STREAM_TYPE, 0, 0, 0, SIZE1, SIZE2, SIZE3, SIZE4}
	//
	// `STREAM_TYPE` can be:
	//
	// -   0: stdin (will be written on stdout)
	// -   1: stdout
	// -   2: stderr
	//
	// `SIZE1, SIZE2, SIZE3, SIZE4` are the 4 bytes of
	// the uint32 size encoded as big endian.

	h := make([]byte, 8)
	buf := bytes.NewBuffer(b)
	buf.Read(h)

	var t string

	switch h[0] {
	case 0:
		t = "stdin"
	case 1:
		t = "stdout"
	case 2:
		t = "stderr"
	}

	logmsg := bytes.SplitN(buf.Bytes(), []byte(" - - "), 2)

	part := bytes.SplitN(logmsg[0], []byte(" "), 3)

	labels := make(map[string]string)

	l := string(part[1])

	if l != "" {
		items := strings.Split(l, ",")

		for _, val := range items {
			pair := strings.SplitN(val, "=", 2)

			labels[pair[0]] = pair[1]
		}
	}

	fmt.Printf("--> logmsg = %v\n", logmsg)

	return &DockerLog{
		Type:      t,
		Labels:    labels,
		Timestamp: string(part[0]),
		IP:        string(part[2]),
		Message:   string(logmsg[1]),
	}
}
