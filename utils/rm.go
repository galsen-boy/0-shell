package utility

import (
	"fmt"
	"os"
)

func RemoveFile(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: rm <filename> [-r]")
		return
	}

	recursive := false
	var filenames []string

	// Vérifie les arguments pour l'option -r
	for _, arg := range args {
		if arg == "-r" {
			recursive = true
		} else {
			filenames = append(filenames, arg)
		}
	}

	for _, filename := range filenames {
		if recursive {
			err := os.RemoveAll(filename)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := os.Remove(filename)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
