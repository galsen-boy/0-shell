package main

import (
	"bufio"
	"fmt"
	"io"
	// "io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
	"sort"
)

func main() {
	supportedCommands := "echo, cd, ls, pwd, cat, cp, rm, mv, mkdir, exit"
	allowedCommands := []string{"ls", "pwd", "cat", "cp", "rm", "mv", "mkdir", "exit"}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println()
				return
			}
			fmt.Println(err)
			continue
		}
		command = strings.TrimSpace(command)
		if command == "exit" {
			break
		}
		args := strings.Fields(command)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "cd":
			changeDirectory(args)
		case "echo":
			echo(args[1:])
		case "pwd":
			printWorkingDirectory()
		case "ls":
			listDirectory(args)
		case "cat":
			catFile(args[1:])
		case "cp":
			copyFile(args[1:])
		case "rm":
			removeFile(args[1:])
		case "mv":
			moveFile(args[1:])
		case "mkdir":
			makeDirectory(args[1:])
		default:
			if !isAllowed(args[0], allowedCommands) && args[0] != "" {
				fmt.Printf("Command [%s] not allowed.\n", args[0])
				fmt.Printf("Supported commands: %s.\n", supportedCommands)
			}
		}
	}
}

func isAllowed(cmd string, allowedCommands []string) bool {
	for _, allowed := range allowedCommands {
		if cmd == allowed {
			return true
		}
	}
	return false
}

func changeDirectory(args []string) {
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

func echo(args []string) {
	str := strings.Join(args, " ")
	if len(str) > 1 && ((str[0] == '"' && str[len(str)-1] == '"') || (str[0] == '\'' && str[len(str)-1] == '\'')) {
		str = str[1 : len(str)-1]
	}
	fmt.Println(str)
}

func printWorkingDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dir)
}
func listDirectory(args []string) {
	var showAll, longFormat, showIndicators bool
	dir := "."
	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, "-") {
			for _, flag := range arg[1:] {
				switch flag {
				case 'a':
					showAll = true
				case 'l':
					longFormat = true
				case 'F':
					showIndicators = true
				}
			}
		} else {
			dir = arg
		}
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	var files []os.FileInfo
	if showAll {
		curDirInfo, _ := os.Stat(dir)
		files = append(files, curDirInfo)
		parentDirInfo, _ := os.Stat(filepath.Dir(dir))
		files = append(files, parentDirInfo)
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if info.IsDir() || !strings.HasPrefix(entry.Name(), ".") {
			files = append(files, info)
		}
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	if longFormat {
		totalBlocks := int64(0)
		for _, file := range files {
			stat := file.Sys().(*syscall.Stat_t)
			totalBlocks += stat.Blocks
		}
		fmt.Printf("total %d\n", totalBlocks/2)
	}

	for _, file := range files {
		name := file.Name()
		if name == "." || name == ".." {
			printFileInfo(file, name, longFormat, false)
		} else {
			printFileInfo(file, name, longFormat, showIndicators)
		}
		fmt.Println()
	}
}
func printFileInfo(info os.FileInfo, name string, longFormat bool, showIndicators bool) {
	if longFormat {
		printLongFormat(info, name)
	} else {
		fmt.Print(name)
	}
}


	func printLongFormat(info os.FileInfo, name string) {
		stat := info.Sys().(*syscall.Stat_t)
		
		fmt.Printf("%s ", info.Mode().String())
		fmt.Printf("%3d ", stat.Nlink)
		
		owner, err := user.LookupId(strconv.Itoa(int(stat.Uid)))
		if err != nil {
			fmt.Print("? ")
		} else {
			fmt.Printf("%-8s ", owner.Username)
		}
		
		group, err := user.LookupGroupId(strconv.Itoa(int(stat.Gid)))
		if err != nil {
			fmt.Print("? ")
		} else {
			fmt.Printf("%-8s ", group.Name)
		}
		
		fmt.Printf("%8d ", info.Size())
		fmt.Printf("%s ", formatTime(info.ModTime()))
		fmt.Print(name)
	}
	
	func getFileIndicator(info os.FileInfo) string {
		if info.IsDir() {
			return "/"
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return "@"
		}
		if info.Mode()&os.ModeNamedPipe != 0 {
			return "|"
		}
		if info.Mode()&os.ModeSocket != 0 {
			return "="
		}
		if info.Mode()&os.ModeDevice != 0 {
			return "%"
		}
		if info.Mode().IsRegular() && info.Mode()&0111 != 0 {
			return "*"
		}
		return ""
	}
	

func catFile(args []string) {
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

func copyFile(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: cp <source> <destination>")
		return
	}

	source := args[0]
	destination := args[1]

	// Ouvrir le fichier source
	inputFile, err := os.Open(source)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	// Vérifier si la destination est un dossier
	info, err := os.Stat(destination)
	if err == nil && info.IsDir() {
		// Créer le chemin complet pour le fichier dans le dossier de destination
		destination = filepath.Join(destination, filepath.Base(source))
	}

	// Créer le fichier destination
	outputFile, err := os.Create(destination)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	// Copier le contenu du fichier source vers le fichier destination
	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		fmt.Println(err)
	}
}

func removeFile(args []string) {
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
			} else {
				// fmt.Printf("Removed directory recursively: %s\n", filename)
			}
		} else {
			err := os.Remove(filename)
			if err != nil {
				fmt.Println(err)
			} else {
				// fmt.Printf("Removed file: %s\n", filename)
			}
		}
	}
}


func moveFile(args []string) {
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

func makeDirectory(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: mkdir <directory1> [directory2] ...")
		return
	}
	for _, dir := range args {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
		} else {
			// fmt.Printf("Created directory: %s\n", dir)
		}
	}
}
func formatTime(t time.Time) string {
	now := time.Now()
	if t.Year() == now.Year() {
		return t.Format("Jan _2 15:04")
	}
	return t.Format("Jan _2  2006")
}