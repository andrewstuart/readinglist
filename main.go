package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var linksFileName string

var fn = flag.String("list", "links", "which list to manage")
var args []string

func init() {
	flag.Parse()
	linksFileName = fmt.Sprintf("%s.json", *fn)
	args = flag.Args()
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

	if len(args) < 2 {
		printLinks(links)
		return
	}

	switch args[1] {
	case "git":
		if len(args) < 3 {
			fmt.Println("Missing arguments to 'git' subcommand")
		}
		out, err := runGit(args[2:]...)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(out)
		return
	case "push", "add":
		if len(args) < 3 {
			warnEmptylinks()
			return
		}

		links = append(links, args[2])
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
	case "rm":
		if len(args) < 3 {
			fmt.Println("No argument for rm")
			return
		}

		links, err = tryRemove(links, args[2])
		if err != nil {
			fmt.Printf("Error removing link: %v\n", err)
			return
		}
	case "splice":
		if len(args) < 3 {
			fmt.Println("No argument for splice")
			return
		}

		tryOpenN(links, 2)

		links, err = tryRemove(links, args[2])
		if err != nil {
			log.Printf("error splicing link: %v\n", err)
			return
		}
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

	if checkGit() != nil {
		return
	}

	commitGit()
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
