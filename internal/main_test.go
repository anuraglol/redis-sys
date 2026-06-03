package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	go main()
	time.Sleep(1 * time.Second)

	commands := [][]string{
		{"SET", "foo", "bar"},
		{"GET", "foo"},
		{"PING"},
		{"EXPIRE", "foo", "5"},
	}

	for range 1000 {
		go func() {
			cmdArgs := commands[rand.Intn(len(commands))]
			args := append([]string{"-p", "8000"}, cmdArgs...)
			cmd := exec.Command("redis-cli", args...)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Failed to run redis-cli command:", err)
			}
			time.Sleep(1 * time.Millisecond)
		}()
	}
}
