package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "args is not enough...")
		os.Exit(1)
	}
	var cmd *exec.Cmd
	if len(args) == 1 {
		cmd = exec.Command(args[0])
	} else {
		cmd = exec.Command(args[0], args[1:]...)
	}
	// cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	pid := make(chan int, 1)
	go func() {
		for {
			if cmd.Process != nil {
				pid <- cmd.Process.Pid
				break
			}
		}
	}()

	fmt.Println(<-pid)
}
