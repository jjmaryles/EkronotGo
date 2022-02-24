package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Enter file path: ")
	var path string
	fmt.Scanln(&path)

	var vmFiles []string = GetVmFiles(path)

	outfile, err1 := os.Create("ex0.asm")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer outfile.Close()

	var name string
	var line string
	for _, file := range vmFiles {
		name = filepath.Base(file)
		outfile.WriteString(name[:len(name)-3] + "\n")

		current, err2 := os.Open(file)
		if err2 != nil {
			log.Fatal(err2)
		}
		defer current.Close()

		scanner := bufio.NewScanner(current)
		for scanner.Scan() {
			line = scanner.Text()
			line := strings.Fields(line)
			println(line)
		}

		if err2 := scanner.Err(); err2 != nil {
			log.Fatal(err2)
		}
	}
}

func GetVmFiles(folder string) []string {
	var files []string

	root := folder
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	var res []string
	for _, file := range files {
		if file[len(file)-3:] == ".vm" {
			res = append(res, file)
		}
	}
	return res
}

func HandleBuy(a string, b string, c string) {

}

func HandleSell(a string, b string, c string) {

}
