package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Print("k-shell > ")
		} else {
			fmt.Print("k@" + cwd + " $ ")
		}
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		execInput(input)

	}
}

func execInput(input string) {
	input = strings.TrimSpace(input)

	if input == "" {
		return
	}

	args := strings.Split(input, " ")

	switch args[0] {
	case "exit":
		os.Exit(0)
	case "cd":
		if len(args) < 2 {
			fmt.Println("cd: missing argument")
			return
		}
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("error finding directory :", err)
			return
		}
		fmt.Println(dir)
		return
	case "ls":
		list, err := os.ReadDir("./")
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		for _, file := range list {
			fmt.Println(file.Name())
		}
		return
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("command error", err)
	}
}
