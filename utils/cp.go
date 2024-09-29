package utility

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFile(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: cp <source> <destination>")
		return
	}

	source := args[0]
	destination := args[1]

	// Ouvrir le fichier source
	inputFile, err := os.Open(source)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier source:", err)
		return
	}
	defer inputFile.Close()

	// Obtenir les informations du fichier ou dossier de destination
	info, err := os.Stat(destination)
	if err == nil && info.IsDir() {
		// Si la destination est un dossier, construire le chemin du fichier
		destination = filepath.Join(destination, filepath.Base(source))
	}

	// Créer le fichier de destination
	outputFile, err := os.Create(destination)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier destination:", err)
		return
	}
	defer outputFile.Close()

	// Copier le contenu du fichier source vers le fichier destination
	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		fmt.Println("Erreur lors de la copie du contenu:", err)
		return
	}

	// Copier les permissions du fichier source vers le fichier destination
	sourceInfo, err := os.Stat(source)
	if err != nil {
		fmt.Println("Erreur lors de l'obtention des informations du fichier source:", err)
		return
	}
	err = os.Chmod(destination, sourceInfo.Mode())
	if err != nil {
		fmt.Println("Erreur lors de la modification des permissions du fichier:", err)
		return
	}

	fmt.Println("Fichier copié avec succès.")
}
