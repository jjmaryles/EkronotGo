package ex4
/**
The following program imitates the syntax analyser part of the compiler
This is done in two steps:
1. Tokenizing
2. Parsing
 */
import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"bufio"
)

func main() {
	//get filepath input
	var path string
	fmt.Println("Enter file path: ")
	fmt.Scanln(&path)

	//create list of ".jack" files
	jackFiles := GetJackFiles(path)


	//for each jack file
	for _, file := range jackFiles {
		current, err2 := os.Open(file)
		if err2 != nil {
			log.Fatal(err2)
		}
		defer current.Close()

		//create a corresponding xml file for translation
		currentOut := filepath.Base(NoSuffix(file)) + "T.xml"
		out, err1 := os.Create(currentOut)
		if err1 != nil {
			log.Fatal(err1)
		}

		defer out.Close()

		scanner := bufio.NewScanner(current)

		for scanner.Scan(){
			temp := strings.Split(scanner.Text(), " ")
			line := RemoveSpaces(temp)
		}
	}

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

func RemoveSpaces(line []string) string {
	res := strings.Join(line,"")
	return res
}