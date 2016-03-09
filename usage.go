package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

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
