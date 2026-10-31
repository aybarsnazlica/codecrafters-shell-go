package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var builtins = []string{"exit", "type", "echo", "pwd"}

func isBuiltin(command string) bool {
	return slices.Contains(builtins, command)
}

func executeBuiltin(fields []string) error {
	switch fields[0] {
	case "exit":
		os.Exit(0)
	case "type":
		if len(fields) < 2 {
			return fmt.Errorf("type: missing argument")
		}

		if isBuiltin(fields[1]) {
			fmt.Printf("%s is a shell builtin\n", fields[1])
		} else if path, err := exec.LookPath(fields[1]); err == nil {
			fmt.Printf("%s is %s\n", fields[1], path)
		} else {
			return fmt.Errorf("%s: not found", fields[1])
		}
	case "echo":
		fmt.Println(strings.Join(fields[1:], " "))
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("error getting current directory: %w", err)
		} else {
			fmt.Println(dir)
		}
	}
	return nil
}

func executeExternal(fields []string) error {
	if _, err := exec.LookPath(fields[0]); err == nil {
		out, err := exec.Command(fields[0], fields[1:]...).Output()
		if err != nil {
			return fmt.Errorf("error executing command: %w", err)
		} else {
			fmt.Print(string(out))
		}
	} else {
		return fmt.Errorf("%s: command not found", fields[0])
	}
	return nil
}

func main() {
	buf := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		command, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		command = strings.TrimSpace(command)
		fields := strings.Fields(command)

		if len(fields) == 0 {
			continue
		}

		if isBuiltin(fields[0]) {
			if err := executeBuiltin(fields); err != nil {
				fmt.Println(err)
			}
		} else {
			if err := executeExternal(fields); err != nil {
				fmt.Println(err)
			}
		}
	}
}
