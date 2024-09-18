// Ce fichier contient des fonctions utilitaires pour obtenir le répertoire initial et afficher le prompt.

package utils

import (
    "fmt"
    "log"
    "os"
)

// GetInitialDir récupère le répertoire de travail initial
func GetInitialDir() (string, error) {
    return os.Getwd()
}

// DisplayPrompt affiche le prompt avec le répertoire courant
func DisplayPrompt() {
    currentDir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    // Affiche le répertoire courant en vert suivi de "$"
    fmt.Printf("\033[32m%s\033[0m $ ", currentDir)
}