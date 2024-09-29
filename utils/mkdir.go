package utility

import (
	"fmt"
	"os"
)

func MakeDirectory(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: mkdir <directory1> [directory2] ...")
		return
	}
	for _, dir := range args {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
		}
	}
}
