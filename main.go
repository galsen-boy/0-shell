package main

import (
    "bufio"
    "log"
    "os"
    "0-shell/utils"
)

func main() {
    // Obtient le répertoire initial
    initialDir, err := utils.GetInitialDir()
    if err != nil {
        log.Fatal(err)
    }

    // Crée un scanner pour lire l'entrée utilisateur
    scanner := bufio.NewScanner(os.Stdin)

    // Boucle principale du shell
    for {
        // Affiche le prompt
        utils.DisplayPrompt()

        // Lit une ligne de l'entrée utilisateur
        if !scanner.Scan() {
            break
        }

        line := scanner.Text()

        // Sort de la boucle si l'utilisateur entre "exit"
        if line == "exit" {
            break
        }

        // Traite la ligne de commande
        utils.ProcessLine(line, initialDir)
    }

    // Gère les erreurs potentielles du scanner
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}