package main

//C:\Users\jjmar\GolandProjects\Go-Compiler\ex0
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//get filepath input
	var path string
	fmt.Println("Enter file path: ")
	fmt.Scanln(&path)

	//create list of ".vm" files
	vmFiles := GetVmFiles(path)

	//create output files to write to
	var outFiles []string
	for _, file := range vmFiles {
		outName := filepath.Base(NoSuffix(file)) + ".asm"
		out, err1 := os.Create(outName)
		if err1 != nil {
			log.Fatal(err1)
		}
		defer out.Close()
		outFiles = append(outFiles, out.Name())
	}

	var toWrite []string
	var fileName string
	for _, file := range vmFiles {
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
			switch line[0] {
			case "push":
				switch line[1] {
				case "local":
					toWrite = PopLocalX(line, file)
				case "argument":
				case "this":
				case "that":
				case "temp":
				case "static":
				case "pointer":
				case "constant":
				}
			case "pop":
				switch line[1] {
				case "local":
				case "argument":
				case "this":
				case "that":
				case "temp":
				case "static":
				case "pointer":
				case "constant":
				}
			}
		}

		if err2 := scanner.Err(); err2 != nil {
			log.Fatal(err2)
		}
	}
}

func NoSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
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

func PopLocalX(command []string, writeFile string) []string {

}
