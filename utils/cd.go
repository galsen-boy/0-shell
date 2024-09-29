package utility

import (
	"fmt"
	"os"
)

func ChangeDirectory(args []string) {
	var dir string
	if len(args) > 1 {
		dir = args[1]
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			return
		}
		dir = homeDir
	}
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println(err)
	}
}
