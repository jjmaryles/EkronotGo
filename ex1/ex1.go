package main

//C:\Users\jjmar\GolandProjects\Go-Compiler\ex0
//C:\Users\yehos\OneDrive\Desktop\School\SemB\Ekronot\nand2tetris\projects\07C:\Users\yehos\OneDrive\Desktop\School\SemB\Ekronot\nand2tetris\projects\07
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

	/*
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
	*/

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
				toWrite = WritePush(line, fileName)
			case "pop":
				toWrite = WritePop(line, fileName)
			case "add","sub","and","or", "neg", "not", "eq", "lt", "gt":
				toWrite = writeArithmetic(line)
			}
		}

		if err2 := scanner.Err(); err2 != nil {
			log.Fatal(err2)
		}

		currentOut := filepath.Base(NoSuffix(file)) + ".asm"
		out, err1 := os.Create(currentOut)
		if err1 != nil {
			log.Fatal(err1)
		}
		defer out.Close()
		res := strings.Join(toWrite,"")
		out.WriteString(res)
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

func WritePop(command []string, fileName string) []string {
	var res []string
	val := command[2]
	switch command[1] {
	case "local":
		res = append(res,"@LCL","D=M","@" + val,"D=D+A")
	case "argument":
		res = append(res,"@ARG","D=M","@" + val,"D=D+A")
	case "this":
		res = append(res,"@THIS","D=M","@" + val,"D=D+A")
	case "that":
		res = append(res,"@THAT","D=M","@" + val,"D=D+A")
	case "pointer":
		if val == "0" {
			res = append(res,"@THIS","D=A")
		} else {
			res = append(res,"@THAT","D=A")
		}
	case "static":
		res = append(res,"@" + fileName + "." + val,"D=A")
	case "temp":
		res = append(res,"@R5","D=A","@" + val,"D=D+A",)
	}
	return append(res,"@R13","M=D","@SP","AM=M-1","D=M","@R13","A=M","M=D")
}

func WritePush(command []string, fileName string) []string {
	var res []string
	val := command[2]
	switch command[1] {
	case "constant":
		res = append(res,"@" + val,"D=A")
	case "local":
		res = append(res,"@LCL","D=M","@" + val,"A=D+A","D=M")
	case "argument":
		res = append(res,"@ARG","D=M","@" + val,"A=D+A","D=M")
	case "this":
		res = append(res,"@THIS","D=M","@" + val,"A=D+A","D=M")
	case "that":
		res = append(res,"@THAT","D=M","@" + val,"A=D+A","D=M")
	case "pointer":
		if val == "0" {
			res = append(res,"@THIS","D=M")
		} else {
			res = append(res,"@THAT","D=M")
		}
	case "static":
		res = append(res,"@" + fileName + "." + val,"D=M")
	case "temp":
		res = append(res,"@R5","D=A","@" + val,"A=D+A","D=M")
	}
	return append(res,"@SP","A=M","M=D","@SP","M=M+1")
}

func writeArithmetic(line []string) []string {
	firstWord := line[0]
	var res []string
	var strTemp string
	count := 0
	switch firstWord{
	case "add","sub","and","or":
		if firstWord == "add"{
			strTemp = "D=D+M"
		} else if firstWord == "sub"{
			strTemp = "D=M-D"
		} else if firstWord == "and"{
			strTemp = "D=D&M"
		} else if firstWord == "or"{
			strTemp = "D=D|M"
		}
		res = append(res, "@SP\n")
		res = append(res, "M=M-1\n")
		res = append(res, "A=M\n")
		res = append(res, "D=M\n")
		res = append(res, "A=A-1\n")
		res = append(res,  strTemp)
		res = append(res, "\n")
		res = append(res, "M=D\n")
		res = append(res, "D=A+1\n")
		res = append(res, "@SP\n")
		res = append(res, "M=D\n")

	case "neg", "not":
		if firstWord == "neg" {
			strTemp = "M=-M"
		}else{
			strTemp = "M!=M"
		}
		res = append(res, "@SP\n")
		res = append(res, "M=M-1\n")
		res = append(res, "A=M\n")
		res = append(res,  strTemp)
		res = append(res, "\n")
		res = append(res, "D=A+1\n")
		res = append(res, "@SP\n")
		res = append(res, "M=D\n")

	case "eq", "lt", "gt":
		if firstWord == "eq"{
			strTemp = "D;JEQ"
		} else if firstWord == "lt"{
			strTemp = "D;JLT"
		} else if firstWord == "gt"{
			strTemp = "D;JGT"
		}
		res = append(res, "@SP\n")
		res = append(res, "M=M-1\n")
		res = append(res, "A=M\n")
		res = append(res, "D=M\n")
		res = append(res, "A=A-1\n")
		res = append(res, "D=M-D\n")
		res = append(res, fmt.Sprint("@EQ%d\n", count))
		res = append(res, strTemp)
		res = append(res, "\n")
		res = append(res, "@0\n")
		res = append(res, "D=-A\n")
		res = append(res, "@SP\n")
		res = append(res, "A=M\n")
		res = append(res, "M=D\n")
		res = append(res, "@SP\n")
		res = append(res, "M=M+1\n")
		res = append(res, fmt.Sprint("@doneEQ%d\n", count))
		res = append(res, "0;JMP\n")
		res = append(res, fmt.Sprint("EQ%d\n", count))
		res = append(res, "@1\n")
		res = append(res, "D=A\n")
		res = append(res, "@SP\n")
		res = append(res, "A=M\n")
		res = append(res, "M=D\n")
		res = append(res, "@SP\n")
		res = append(res, "M=M+1\n")
		res = append(res, fmt.Sprint("doneEQ%d\n", count))

		count = count + 1
	}

	return res
}