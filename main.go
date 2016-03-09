package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

const linksFileName = "links.json"

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
	case "git":
		if len(os.Args) < 3 {
			fmt.Println("Missing arguments to 'git' subcommand")
		}
		out, err := runGit(os.Args[2:]...)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(out)
		return
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
	case "rm":
		if len(os.Args) < 3 {
			fmt.Println("No argument for rm")
			return
		}

		links, err = tryRemove(links, os.Args[2])
		if err != nil {
			fmt.Printf("Error removing link: %v\n", err)
			return
		}
	case "splice":
		if len(os.Args) < 3 {
			fmt.Println("No argument for splice")
			return
		}

		tryOpenN(links, 2)

		links, err = tryRemove(links, os.Args[2])
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
