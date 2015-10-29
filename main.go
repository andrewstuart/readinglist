package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/toqueteos/webbrowser"
)

const linksFileName = "links.json"

var usage = fmt.Sprintf(`
usage: 
 - %s [(ls|show)] - Show all links
 - %s push <link> - Add a link
 - %s pop - Print and remove the most recent link
 - %s shift - Print and remove the oldest link
 - %s [open] <number> - Open the link at number <number>
 `, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])

var rlHome, rlFileName string

func init() {
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal("Could not determine home directory for reading list storage.")
	}

	rlHome = fmt.Sprintf("%s/.local/readinglinks/", dir)
	rlFileName = rlHome + linksFileName

	if _, err := os.Stat(rlHome); err != nil {
		err := os.MkdirAll(rlHome, 0770)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func main() {
	f, err := getFile()
	if err != nil {
		log.Fatal(err)
	}
	var links []string
	if err := json.NewDecoder(f).Decode(&links); err != nil {
		links = make([]string, 0, 1)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		printLinks(links)
		return
	}

	switch os.Args[1] {
	case "push":
		if len(os.Args) < 3 {
			warnEmptylinks()
			return
		}

		links = append(links, os.Args[2])
	case "pop":
		if len(links) < 1 {
		}

		tryOpen(links[len(links)-1])
		links = links[:len(links)-1]
	case "show", "ls":
		printLinks(links)
	case "shift":
		if len(links) < 1 {
			fmt.Println("No items in links")
			return
		}

		tryOpen(links[0])
		links = links[1:]
	case "open":
		tryOpenN(links, 2)
	default:
		tryOpenN(links, 1)
	}

	f, err = getFile()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(f).Encode(links)
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func tryOpenN(links []string, arg int) {
	if len(os.Args) < arg+1 {
		fmt.Println(usage)
		return
	}

	i, err := strconv.Atoi(os.Args[arg])
	if err != nil {
		fmt.Println("Argument was not a number")
		os.Exit(1)
	}

	if i < 1 || i > len(links) {
		fmt.Printf("Invalid number. Acceptable range: %d-%d\n", 1, len(links))
		printLinks(links)
		os.Exit(1)
	}

	tryOpen(links[i-1])
}

func getFile() (*os.File, error) {
	_, err := os.Stat(rlFileName)
	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		return os.Create(rlFileName)
	}
	return os.OpenFile(rlFileName, os.O_RDWR, 0550)
}

func warnEmptylinks() {
	fmt.Println("No items in links")
}

func tryOpen(link string) {
	err := webbrowser.Open(link)
	if err != nil {
		fmt.Println(link)
	}
}

func printLinks(links []string) {
	if len(links) < 1 {
		fmt.Println("No items in reading list.")
	}
	for i := range links {
		fmt.Printf("%d.\t%s\n", i+1, links[i])
	}
}
