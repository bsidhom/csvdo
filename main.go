package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <command>\n", os.Args[0])
		return 2
	}
	command := os.Args[1]
	reader := csv.NewReader(os.Stdin)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			return 0
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
			return 1
		}
		cmd := exec.Command(command, record...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not execute %q: %v", record, err)
			return 1
		}
	}
}
