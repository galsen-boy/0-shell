package main

import (
	"bufio"
	"fmt"
	"io"
	utility "main/utils"
	"os"
	"strings"
)

func main() {
	supportedCommands := "echo, cd, ls, pwd, cat, cp, rm, mv, mkdir, exit"
	allowedCommands := []string{"ls", "pwd", "cat", "cp", "rm", "mv", "mkdir", "exit"}

	reader := bufio.NewReader(os.Stdin)
	for {
		utility.PrintWorkingDirectoryWithDolllar()

		command, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println()
				return
			}
			fmt.Println(err)
			continue
		}
		command = strings.TrimSpace(command)
		if command == "exit" {
			break
		}
		args := strings.Fields(command)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "cd":
			utility.ChangeDirectory(args)

		case "echo":
			utility.Echo(args[1:])
		case "pwd":
			utility.PrintWorkingDirectory()
		case "ls":
			utility.ListDirectory(args)
		case "cat":
			utility.CatFile(args[1:])
		case "cp":
			utility.CopyFile(args[1:])
		case "rm":
			utility.RemoveFile(args[1:])
		case "mv":
			utility.MoveFile(args[1:])
		case "mkdir":
			utility.MakeDirectory(args[1:])
		default:
			if !isAllowed(args[0], allowedCommands) && args[0] != "" {
				fmt.Printf("Command [%s] not allowed.\n", args[0])
				fmt.Printf("Supported commands: %s.\n", supportedCommands)
			}
		}
	}
}

func isAllowed(cmd string, allowedCommands []string) bool {
	for _, allowed := range allowedCommands {
		if cmd == allowed {
			return true
		}
	}
	return false
}
