package main
/**
The following program imitates the syntax analyser part of the compiler
This is done in two steps:
1. Tokenizing
2. Parsing
jj is phat
 */
import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"regexp"
)

func main() {
	//get filepath input
	var path string
	fmt.Println("Enter file path: ")
	fmt.Scanln(&path)

	//create list of ".jack" files
	jackFiles := GetJackFiles(path)

	//open jackFiles to read from
	var toWrite []string

	//for each jack file
	for _, file := range jackFiles {
		current, err2 := os.Open(file)
		if err2 != nil {
			log.Fatal(err2)
		}
		defer current.Close()

		//create a corresponding xml file for translation
		currentOut := filepath.Base(NoSuffix(file)) + ".xml"
		out, err1 := os.Create(currentOut)
		if err1 != nil {
			log.Fatal(err1)
		}

		defer out.Close()

	}
	
	/*
	rKeyword := regexp.MustCompile("class|constructor|function|method|field|static|var|int|char|boolean|void|true|false|null|this|let|do|if|else|while|return")
	rSymbol := regexp.MustCompile("\\{ | \\} | \\( | \\) | \\[ | \\] | \\. | , | ; | \\+ | \\- | \\* | / | \\| |<|>|=|~")
	rIntConst, _ :=  regexp.MustCompile("[0-9][0-9]")
	rStrConst, _ :=  regexp.MustCompile("\"\P{}\"")
	 */

}

// NoSuffix function returns file name without type
func NoSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// GetJackFiles function returns of list of ".jack" files in a specified directory
func GetJackFiles(folder string) []string {
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
		if file[len(file)-3:] == ".jack" {
			res = append(res, file)
		}
	}
	return res
}
