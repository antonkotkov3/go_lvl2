package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	path *string
	file *string
	d    *bool
)

func init() {
	path = flag.String("path", ".", "dir for search")
	file = flag.String("file", "", "file name for search")
	d = flag.Bool("d", false, "Delete duplicated files?")
	flag.Parse()
}

func main() {
	p, _ := os.Getwd()
	fmt.Println(p)
	duplicateList, err := FindDuplicate(*path, *file)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	duplicateCount := len(duplicateList)

	if duplicateCount > 0 {
		fmt.Printf("Found duplicates: %d\n", duplicateCount)
		for i, duplicateName := range duplicateList {
			fmt.Printf("%d. %s\n", i+1, duplicateName)
		}
	} else {
		fmt.Printf("No copy of %q in path: %q\n", *file, *path)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	if *d {
		fmt.Print("\nAre you sure you want to delete the duplicates? ('Y' to confirm):\t")
		scanner.Scan()
		userAnswer := scanner.Text()
		if strings.ToLower(userAnswer) == "y" {
			for _, duplicateName := range duplicateList {
				fmt.Println("- deleted:", duplicateName)
			}
		}
	}
}
