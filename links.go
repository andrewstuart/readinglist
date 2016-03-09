package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/toqueteos/webbrowser"
)

func tryOpenN(links []string, arg int) {
	if len(os.Args) != arg+1 {
		fmt.Println(usage)
		return
	}

	i, err := strconv.Atoi(os.Args[arg])
	if err != nil {
		fmt.Println("Argument was not a number")
		fmt.Println(usage)
		os.Exit(1)
	}

	if i < 1 || i > len(links) {
		fmt.Printf("Invalid number. Acceptable range: %d-%d\n", 1, len(links))
		printLinks(links)
		os.Exit(1)
	}

	tryOpen(links[i-1])
}

func tryOpen(link string) {
	err := webbrowser.Open(link)
	if err != nil {
		fmt.Println(link)
	}
}

func tryRemove(links []string, number string) ([]string, error) {
	i, err := strconv.Atoi(number)

	if err != nil {
		return links, fmt.Errorf("invalid number for rm: %s", os.Args[2])
	}

	if len(links) < i {
		return links, fmt.Errorf("invalid number, longer than list: %d", i)
	}

	if i < 1 {
		return links, fmt.Errorf("invalid link number less than zero")
	}

	links = append(links[0:i-1], links[i:]...)

	return links, nil
}

func printLinks(links []string) {
	if len(links) < 1 {
		fmt.Println("No items in reading list.")
	}
	for i := range links {
		fmt.Printf("%d.\t%s\n", i+1, links[i])
	}
}
