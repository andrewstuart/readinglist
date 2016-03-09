package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/toqueteos/webbrowser"
)

const rangeSplit = ".."

func tryOpenN(links []string, arg int) {
	if len(args) != arg+1 {
		fmt.Println(usage)
		return
	}

	from, to := args[arg], args[arg]

	if strings.Contains(from, rangeSplit) {
		rng := strings.Split(from, rangeSplit)

		if len(rng) < 2 {
			fmt.Println("illegal range given; format is <number>..<number>")
			os.Exit(1)
		}

		from, to = rng[0], rng[1]
	}

	fromI, err := strconv.Atoi(from)
	if err != nil {
		fmt.Println("Argument was not a number")
		fmt.Println(usage)
		os.Exit(1)
	}

	toI, err := strconv.Atoi(to)
	if err != nil {
		fmt.Println("Argument was not a number")
		fmt.Println(usage)
		os.Exit(1)
	}

	if fromI > toI {
		fromI, toI = toI, fromI
	}

	if toI < 1 || fromI > len(links) {
		fmt.Printf("Invalid number. Acceptable range: %d-%d\n", 1, len(links))
		printLinks(links)
		os.Exit(1)
	}

	for i := fromI; i <= toI && i < len(links) && i > 0; i++ {

		tryOpen(links[i-1])
	}
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
		return links, fmt.Errorf("invalid number for rm: %s", args[2])
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
