package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.TrimSpace(command)
		fields := strings.Fields(command)

		builtins := []string{"exit", "type", "echo"}

		if fields[0] == "exit" {
			break
		} else if fields[0] == "type" {
			if slices.Contains(builtins, fields[1]) {
				fmt.Printf("%s is a shell builtin\n", fields[1])
			} else if path, err := exec.LookPath(fields[1]); err == nil {
				fmt.Printf("%s is %s\n", fields[1], path)
			} else {
				fmt.Printf("%s: not found\n", fields[1])
			}
		} else if fields[0] == "echo" {
			fmt.Println(strings.Join(fields[1:], " "))
		} else if _, err := exec.LookPath(fields[0]); err == nil {
			out, _ := exec.Command(fields[0], fields[1:]...).Output()
			fmt.Print(string(out))
		} else {
			fmt.Printf("%s: command not found\n", fields[0])
		}
	}
}
