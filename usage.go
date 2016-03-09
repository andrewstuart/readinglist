package main

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

var usage string

var rlHome, rlFileName string

func init() {
	tpl := template.Must(template.New("usage").Parse(`
usage: 
- {{.name}} [(ls|show)] - Show all links
- {{.name}} (push|add) <url> - Add a link
- {{.name}} pop - Open and remove the most recent link
- {{.name}} rm <number> - Remove a link (at <number>) from the list
- {{.name}} shift - Open and remove the first link
- {{.name}} splice <number> - Open and remove link at number <number>
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

}
