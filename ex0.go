package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var outfileName string //global variable as outfile name

func main() {
	//get filepath input
	var path string
	fmt.Println("Enter file path: ")
	fmt.Scanln(&path)

	//create list of ".vm" files
	vmFiles := GetVmFiles(path)

	//create output file to write to
	outfileName = filepath.Base(path) + ".asm"
	outfile, err1 := os.Create(outfileName)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer outfile.Close()

	var name string
	for _, file := range vmFiles {
		//write filename without .vm
		name = filepath.Base(file)
		outfile.WriteString(name[:len(name)-3] + "\n")

		current, err2 := os.Open(file)
		if err2 != nil {
			log.Fatal(err2)
		}
		defer current.Close()

		//read the file
		scanner := bufio.NewScanner(current)
		//for each line, check if it is "buy" or "cell" and call corresponding "Handle" functions
		for scanner.Scan() {
			line := strings.Split(scanner.Text(), " ")
			amount, _ := strconv.Atoi(line[2])
			price, _ := strconv.ParseFloat(line[3], 32)
			if line[0] == "buy" {
				HandleBuy(line[1], amount, price)
			}
			if line[0] == "cell" {
				HandleCell(line[1], amount, price)
			}
		}

		if err2 := scanner.Err(); err2 != nil {
			log.Fatal(err2)
		}
	}
}

// GetVmFiles function returns of list of ".vm" files in a specified directory
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

// HandleBuy Handle function for "buy"
func HandleBuy(ProductName string, Amount int, Price float64) {
	writefile, err1 := os.Open(outfileName)
	if err1 != nil {
		log.Fatal(err1)
	}
	writefile.WriteString("### BUY " + ProductName + " ###\n")
	writefile.WriteString(strconv.FormatFloat(float64(Amount)*Price, 'E', -1, 64) + "/n")
}

//Handle function for "cell"
func HandleCell(ProductName string, Amount int, Price float64) {
	writefile, err1 := os.Open(outfileName)
	if err1 != nil {
		log.Fatal(err1)
	}
	writefile.WriteString("### CELL " + ProductName + " ###\n")
	writefile.WriteString(strconv.FormatFloat(float64(Amount)*Price, 'E', -1, 64) + "/n")
}
