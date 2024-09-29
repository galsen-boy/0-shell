package utility

import (
	"fmt"
	"os"
	"path/filepath"
)

func MoveFile(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: mv <source> <destination>")
		return
	}

	source := args[0]
	destination := args[1]

	// Vérifie si le fichier source existe
	if _, err := os.Stat(source); os.IsNotExist(err) {
		fmt.Printf("Source file or directory does not exist: %s\n", source)
		return
	}

	// Assure-toi que le chemin de destination est valide
	destinationDir := filepath.Dir(destination)
	if _, err := os.Stat(destinationDir); os.IsNotExist(err) {
		fmt.Printf("Destination directory does not exist: %s\n", destinationDir)
		return
	}

	// Déplace le fichier ou le dossier
	err := os.Rename(source, destination)
	if err != nil {
		fmt.Println("Error moving file:", err)
	} else {
		fmt.Printf("Moved %s to %s\n", source, destination)
	}
}
