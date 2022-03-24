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

	var toWrite []string
	for _, file := range vmFiles {
		current, err2 := os.Open(file)
		if err2 != nil {
			log.Fatal(err2)
		}
		defer current.Close()

		currentOut := filepath.Base(NoSuffix(file)) + ".asm"
		out, err1 := os.Create(currentOut)
		if err1 != nil {
			log.Fatal(err1)
		}
		defer out.Close()
		//read the file
		scanner := bufio.NewScanner(current)
		fileName := filepath.Base(NoSuffix(file))
		//for each line, check if it is "buy" or "cell" and call corresponding "Handle" functions
		count := 1

		for scanner.Scan() {
			line := strings.Split(scanner.Text(), " ")
			switch line[0] {
			case "push":
				toWrite = WritePush(line, fileName)
			case "pop":
				toWrite = WritePop(line, fileName)
			case "add","sub","and","or", "neg", "not":
				toWrite = writeArithmetic(line, 0)
			case "eq", "lt", "gt":
				toWrite = writeArithmetic(line, count)
				count = count + 1
			}

			if err2 := scanner.Err(); err2 != nil {
				log.Fatal(err2)
			}

			res := strings.Join(toWrite,"")
			out.WriteString(res)
		}
	}
}

// NoSuffix function returns file name without type
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

// WritePop write function for pop command
func WritePop(command []string, fileName string) []string {
	var res []string
	val := command[2]
	switch command[1] {
	case "local":
		res = append(res,"@LCL\n","D=M\n","@" + val + "\n","D=D+A\n")
	case "argument":
		res = append(res,"@ARG\n","D=M\n","@" + val + "\n","D=D+A\n")
	case "this":
		res = append(res,"@THIS\n","D=M\n","@" + val + "\n","D=D+A\n")
	case "that":
		res = append(res,"@THAT\n","D=M\n","@" + val + "\n","D=D+A\n")
	case "pointer":
		if val == "0" {
			res = append(res,"@THIS\n","D=A\n")
		} else {
			res = append(res,"@THAT\n","D=A\n")
		}
	case "static":
		res = append(res,"@" + fileName + "." + val + "\n","D=A\n")
	case "temp":
		res = append(res,"@R5\n","D=A\n","@" + val +"\n","D=D+A\n")
	}
	return append(res,"@R13\n","M=D\n","@SP\n","M=M-1\n","D=M\n","@R13\n","A=M\n","M=D\n")
}

// WritePush write function for push command
func WritePush(command []string, fileName string) []string {
	var res []string
	val := command[2]
	switch command[1] {
	case "constant":
		res = append(res,"@" + val + "\n","D=A\n")
	case "local":
		res = append(res,"@LCL\n","D=M\n","@" + val + "\n","A=D+A\n","D=M\n")
	case "argument":
		res = append(res,"@ARG\n","D=M\n","@" + val + "\n","A=D+A\n","D=M\n")
	case "this":
		res = append(res,"@THIS\n","D=M\n","@" + val + "\n","A=D+A\n","D=M\n")
	case "that":
		res = append(res,"@THAT\n","D=M\n","@" + val + "\n","A=D+A\n","D=M\n")
	case "pointer":
		if val == "0" {
			res = append(res,"@THIS\n","D=M\n")
		} else {
			res = append(res,"@THAT\n","D=M\n")
		}
	case "static":
		res = append(res,"@" + fileName + "." + val + "\n","D=M\n")
	case "temp":
		res = append(res,"@R5\n","D=A\n","@" + val +"\n","A=D+A\n","D=M\n")
	}
	return append(res,"@SP\n","A=M\n","M=D\n","@SP\n","M=M+1\n")
}

// WriteArithmetic write function for arithmetic command
func writeArithmetic(line []string, count int) []string {
	firstWord := line[0]
	var res []string
	var strTemp string

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
		res = append(res, fmt.Sprintf("@EQ%d\n", count))
		res = append(res, strTemp)
		res = append(res, "\n")
		res = append(res, "@0\n")
		res = append(res, "D=-A\n")
		res = append(res, "@SP\n")
		res = append(res, "A=M\n")
		res = append(res, "M=D\n")
		res = append(res, "@SP\n")
		res = append(res, "M=M+1\n")
		res = append(res, fmt.Sprintf("@doneEQ%d\n", count))
		res = append(res, "0;JMP\n")
		res = append(res, fmt.Sprintf("EQ%d\n", count))
		res = append(res, "@1\n")
		res = append(res, "D=A\n")
		res = append(res, "@SP\n")
		res = append(res, "A=M\n")
		res = append(res, "M=D\n")
		res = append(res, "@SP\n")
		res = append(res, "M=M+1\n")
		res = append(res, fmt.Sprintf("doneEQ%d\n", count))

		count = count + 1
	}

	return res
}