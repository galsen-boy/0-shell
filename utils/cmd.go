// Ce fichier contient les fonctions principales pour traiter etexécuter les commandes shell.
// Il gère spécifiquement la commande "cd" et permet l'utilisation de variables d'environnement temporaires.


package utils

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "regexp"
    "strings"
)

// ProcessLine traite la ligne de commande entrée par l'utilisateur
func ProcessLine(line string, initialDir string) {
    // Vérifie si la commande commence par "cd "
    if strings.HasPrefix(line, "cd ") {
        handleCdCommand(line, initialDir)
        return
    }
    // Sinon, exécute la commande
    executeCommand(line)
}

// handleCdCommand gère la commande "cd" spécifiquement
func handleCdCommand(line string, initialDir string) {
    // Extrait le nouveau répertoire de la commande
    newDir := strings.TrimSpace(strings.TrimPrefix(line, "cd "))
    // Si aucun répertoire n'est spécifié, utilise le répertoire initial
    if newDir == "" {
        newDir = initialDir
    }
    // Change le répertoire courant
    if err := os.Chdir(newDir); err != nil {
        log.Printf("cd: %s\n", err)
    }
}

// executeCommand exécute une commande shell
func executeCommand(line string) {
    // Regex pour détecter les variables d'environnement temporaires
    pattern := regexp.MustCompile(`^(\w+)="?(.*?)"?\s+(.*)`)
    matches := pattern.FindStringSubmatch(line)

    var cmd *exec.Cmd
    if len(matches) > 0 {
        // Si une variable d'environnement temporaire est détectée
        varName := matches[1]
        varValue := strings.Trim(matches[2], `"`)
        command := matches[3]
        // Crée une commande avec la variable d'environnement
        cmd = exec.Command("bash", "-c", command)
        cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", varName, varValue))
    } else {
        // Sinon, crée une commande normale
        cmd = exec.Command("bash", "-c", line)
    }

    // Configure les entrées/sorties de la commande
    cmd.Stdin = os.Stdin
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout

    // Exécute la commande
    if err := cmd.Run(); err != nil {
        log.Printf("failed to execute process: %s\n", err)
    }
}