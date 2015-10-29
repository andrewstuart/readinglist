package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/toqueteos/webbrowser"
)
import "os/user"

const listFileName = "list.json"

var usage = fmt.Sprintf(`
usage: 
 - %s push <link> - Add a link
 - %s pop - Print and remove the most recent link
 - %s shift - Print and remove the oldest link
 - %s (ls|show) - Show all links
 - %s open <number> - Open the link at number <number>
 `, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])

var (
	rlHome     = "~/.local/readinglist/"
	rlFileName = rlHome + listFileName
)

func init() {
	u, err := user.Current()
	if err == nil {
		rlHome = fmt.Sprintf("%s/.local/readinglist/", u.HomeDir)
		rlFileName = rlHome + listFileName
	}

	if _, err := os.Stat(rlHome); err != nil {
		err := os.MkdirAll(rlHome, 0770)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	f, err := getFile()
	if err != nil {
		log.Fatal(err)
	}
	var list []string
	if err := json.NewDecoder(f).Decode(&list); err != nil {
		list = make([]string, 0, 1)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "push":
		if len(os.Args) < 3 {
			warnEmptyList()
		}

		list = append(list, os.Args[2])
	case "pop":
		if len(list) < 1 {
		}

		tryOpen(list[len(list)-1])
		list = list[:len(list)-1]
	case "show", "ls":
		for i := range list {
			fmt.Printf("%d.\t%s\n", i+1, list[i])
		}
	case "shift":
		if len(list) < 1 {
			fmt.Println("No items in list")
			os.Exit(0)
		}

		tryOpen(list[0])
		list = list[1:]
	case "open":
		if len(os.Args) < 3 {
			fmt.Println(usage)
			os.Exit(0)
		}

		i, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Argument was not a number")
			os.Exit(1)
		}

		if i < 1 || i > len(list)+1 {
			fmt.Printf("Invalid number. Acceptable range: %d-%d\n", 1, len(list)+1)
			os.Exit(1)
		}

		tryOpen(list[i-1])
	default:
		fmt.Println(usage)
		os.Exit(1)
	}

	f, err = getFile()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(f).Encode(list)
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func getFile() (*os.File, error) {
	_, err := os.Stat(rlFileName)
	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		return os.Create(rlFileName)
	}
	return os.OpenFile(rlFileName, os.O_RDWR, 0550)
}

func warnEmptyList() {
	fmt.Println("No items in list")
	os.Exit(0)
}

func tryOpen(link string) {
	err := webbrowser.Open(link)
	if err != nil {
		fmt.Println(link)
	}
}
