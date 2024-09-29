package utility

import (
	"fmt"
	"os"
)

func PrintWorkingDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dir)
}

func PrintWorkingDirectoryWithDolllar() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	final:= fmt.Sprintf("%s$", dir)
	fmt.Print(final)
}