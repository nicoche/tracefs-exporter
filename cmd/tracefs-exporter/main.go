package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	for {
		err := tracePipeToStdout()
		if err != nil {
			log.Printf("Failed to dump trace_pipe to stdout, retrying in 1s: %s", err)
		}

		time.Sleep(1 * time.Second)
	}
}

func tracePipeToStdout() error {
	f, err := os.Open("/sys/kernel/debug/tracing/trace_pipe")
	if err != nil {
		return fmt.Errorf("failed to open /sys/kernel/debug/tracing/trace_pipe: %s", err)
	}

	r := io.TeeReader(f, os.Stdout)
	for {
		//var data []byte
		_, err := io.ReadAll(r)
		if err != nil {
			return fmt.Errorf("failed to read from trace_pipe fd: %s", err)
		}

	}
}
