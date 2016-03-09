package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/mitchellh/go-homedir"
)

var usage string

var rlHome, rlFileName string

func init() {
	tpl := template.Must(template.New("usage").Parse(`
usage: 
- {{.name}} [(ls|show)] - Show all links
- {{.name}} push <url> - Add a link
- {{.name}} pop - Print and remove the most recent link
- {{.name}} rm <number> - Remove a link from the list
- {{.name}} shift - Print and remove the oldest link
- {{.name}} [open] <number> - Open the link at number <number>
`))

	m := map[string]string{
		"name": os.Args[0],
	}

	bf := &bytes.Buffer{}

	err := tpl.Execute(bf, m)
	if err != nil {
		log.Fatal("error executing template", err)
	}

	usage = bf.String()

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
