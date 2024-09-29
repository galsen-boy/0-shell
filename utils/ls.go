package utility

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func ListDirectory(args []string) {
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
		if info.IsDir() && showIndicators {
			fmt.Print(name + "/")
		} else {
			fmt.Print(name)
		}
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
func formatTime(t time.Time) string {
	now := time.Now()
	if t.Year() == now.Year() {
		return t.Format("Jan _2 15:04")
	}
	return t.Format("Jan _2  2006")
}
