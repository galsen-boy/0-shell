package utility

import (
	"fmt"
	"os"
)

func CatFile(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: cat <filename>")
		return
	}
	content, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}
