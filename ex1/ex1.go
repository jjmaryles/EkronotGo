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
	//get filepath input
	var path string
	fmt.Println("Enter file path: ")
	fmt.Scanln(&path)

	//create list of ".vm" files
	vmFiles := GetVmFiles(path)

	//open vmFiles to read from
	var toWrite []string

	//for each vm file
	for _, file := range vmFiles {
		current, err2 := os.Open(file)
		if err2 != nil {
			log.Fatal(err2)
		}
		defer current.Close()

		//create a corresponding asm file for translation
		currentOut := filepath.Base(NoSuffix(file)) + ".asm"
		out, err1 := os.Create(currentOut)
		if err1 != nil {
			log.Fatal(err1)
		}

		defer out.Close()

		//read the vm file
		scanner := bufio.NewScanner(current)
		fileName := filepath.Base(NoSuffix(file))
		runNum := 1
		lineNum := 1

		//for each line, translate it
		for scanner.Scan() {
			line := strings.Split(scanner.Text(), " ")
			switch line[0] {
			case "push":
				toWrite = WritePush(line, fileName)
			case "pop":
				toWrite = WritePop(line, fileName, runNum)
				runNum = runNum + 1
			case "add","sub","and","or", "neg", "not":
				toWrite = writeArithmetic(line, 0)
			case "eq", "lt", "gt":
				toWrite = writeArithmetic(line, runNum)
				runNum = runNum + 1
			case "label", "goto", "if-goto":
				toWrite = WriteBranch(line, fileName)
			case "function", "return", "call":
				toWrite = WriteFunction(line, lineNum)
			case "//", "\n", "\t", " ", "":
				toWrite = strings.Split("","")
				lineNum--
			default:
				toWrite = strings.Split("ERROR\n","")
				lineNum--

			lineNum += 1
			}

			if err2 := scanner.Err(); err2 != nil {
				log.Fatal(err2)
			}

			//write the asm code
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
func WritePop(command []string, fileName string, runNum int) []string {
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
		res = append(res, "@3\n", "D=A\n", "@" + val + "\n", "D=D+A\n")
	case "static":
		res = append(res,"@" + fileName + "." + val + "\n","D=A\n")
	case "temp":
		res = append(res,"@5\n","D=A\n","@" + val +"\n","D=D+A\n")
	}
	return append(res,fmt.Sprintf("@addr_%d\n",runNum),"M=D\n","@SP\n","M=M-1\n","A=M\n","D=M\n",fmt.Sprintf("@addr_%d\n",runNum),"A=M\n","M=D\n")
}

// WritePush write function for push command
func WritePush(command []string, fileName string) []string {
	var res []string
	val := command[2]
	switch command[1] {
	case "constant":
		res = append(res,"@" + val + "\n","D=A\n")
	case "local":
		res = append(res,"@LCL\n","D=M\n","@" + val + "\n","D=D+A\n", "A=D\n","D=M\n")
	case "argument":
		res = append(res,"@ARG\n","D=M\n","@" + val + "\n","D=D+A\n", "A=D\n","D=M\n")
	case "this":
		res = append(res,"@THIS\n","D=M\n","@" + val + "\n","D=D+A\n", "A=D\n","D=M\n")
	case "that":
		res = append(res,"@THAT\n","D=M\n","@" + val + "\n","D=D+A\n", "A=D\n","D=M\n")
	case "pointer":
		res = append(res, "@3\n", "D=A\n", "@" + val + "\n", "D=D+A\n", "A=D\n", "D=M\n")
	case "static":
		res = append(res,"@" + fileName + "." + val + "\n","D=M\n")
	case "temp":
		res = append(res,"@5\n","D=A\n","@" + val +"\n","D=D+A\n", "A=D\n","D=M\n")
	}
	return append(res,"@SP\n","A=M\n","M=D\n","@SP\n","M=M+1\n")
}

// WriteArithmetic write function for arithmetic command
func writeArithmetic(line []string, count int) []string {
	firstWord := line[0]
	var res []string
	var strTemp string

	switch firstWord{
	case "add","sub":
		if firstWord == "add"{
			strTemp = "D=D+M"
		} else if firstWord == "sub" {
			strTemp = "D=M-D"
		}
		res = append(res, "@SP\n")
		res = append(res, "M=M-1\n")
		res = append(res, "A=M\n")
		res = append(res, "D=M\n")
		res = append(res, "A=A-1\n")
		res = append(res,  strTemp)
		res = append(res, "\n")
		res = append(res, "M=D\n")

	case "and", "or":
		if firstWord == "and"{
			strTemp = "M=D&M"
		} else if firstWord == "or"{
			strTemp = "M=D|M"
		}
		res = append(res, "@SP\n")
		res = append(res, "M=M-1\n")
		res = append(res, "A=M\n")
		res = append(res, "D=M\n")
		res = append(res, "A=A-1\n")
		res = append(res,  strTemp)
		res = append(res, "\n")

	case "neg", "not":
		if firstWord == "neg" {
			strTemp = "M=-M"
		}else{
			strTemp = "M=!M"
		}
		res = append(res, "@SP\n")
		res = append(res, "A=M-1\n")
		res = append(res,  strTemp + "\n")

	case "eq", "lt", "gt":
		var lable string
		if firstWord == "eq"{
			strTemp = "D;JEQ"
			lable = "EQ"
		} else if firstWord == "lt"{
			strTemp = "D;JLT"
			lable = "LT"
		} else if firstWord == "gt"{
			strTemp = "D;JGT"
			lable = "GT"
		}
		res = append(res, "@SP\n")
		res = append(res, "M=M-1\n")
		res = append(res, "A=M\n")
		res = append(res, "D=M\n")
		res = append(res, "A=A-1\n")
		res = append(res, "D=M-D\n")
		res = append(res, fmt.Sprintf("@%s_%d\n",lable, count))
		res = append(res, strTemp)
		res = append(res, "\n")
		res = append(res, "@SP\n")
		res = append(res, "A=M-1\n")
		res = append(res, "M=0\n")
		res = append(res, fmt.Sprintf("@END_%d\n", count))
		res = append(res, "0;JMP\n")
		res = append(res, fmt.Sprintf("(%s_%d)\n", lable, count))
		res = append(res, "@SP\n")
		res = append(res, "A=M-1\n")
		res = append(res, "M=-1\n")
		res = append(res, fmt.Sprintf("(END_%d)\n", count))

		count = count + 1
	}

	return res
}

func WriteBranch(command []string, fileName string,) []string {
	var res []string
	brType := command[0]
	label := fileName+command[1]

	switch brType {
	case "label":
		res  = append(res,"(" + label +")\n")
	case "goto":
		res = append(res, "@" + label + "\n", "0;JMP\n")
	case "if-goto":
		res = append(res, "@SP\n", "M=M-1\n","A=M\n", "D=M\n", "@" + label +  "\n", "D;JNE\n")
	}

	return res
}

func WriteFunction(VMline []string, line_number int) []string{
	var res []string
	callType := VMline[0]

	switch callType{
	case "function":
		funcName := VMline[1]
		vars := VMline[2]
		res = append(res,  "//function " + funcName + vars +"\n" ,"("+funcName+")\n")
		for i := 2; i < len(VMline); i++ {
			res = append(res, "@SP\n", "A=M\n", "M=0\n", "@SP\n", "M=M+1 // push 0\n")
		}

	case "return":
		res = append(res, "//return\n",
			"@LCL\n",
			"D=M\n",
			"@frame\n",
			"M=D //FRAME = LCL\n",
			"@5\n",
			"D=D-A\n",
			"A=D\n", "" +
			"D=M\n",
			"@return_address\n",
			"M=D //RET = *(FRAME - 5)\n",
			"@SP\n",
			"M=M-1\n",
			"A=M\n",
			"D=M\n",
			"@ARG\n",
			"A=M\n",
			"M=D // *ARG = pop()\n",
			"@ARG\n",
			"D=M+1\n",
			"@SP\n",
			"M=D // SP = ARG + 1\n",
			"@frame\n",
			"D=M-1\n",
			"A=D\n",
			"D=M\n",
			"@THAT\n",
			"M=D // THAT = *(FRAME - 1)\n",
			"@2\n",
			"D=A\n",
			"@frame\n",
			"D=M-D\n",
			"A=D\n",
			"D=M\n",
			"@THIS\n",
			"M=D // THIS = *(FRAME-2)\n",
			"@3\n",
			"D=A\n",
			"@frame\n",
			"D=M-D\n",
			"A=D\n",
			"D=M\n",
			"@ARG\n",
			"M=D // ARG = *(FRAME-3)\n",
			"@4\n",
			"D=A\n",
			"@frame\n" ,
			"D=M-D\n" ,
			"A=D\n" ,
			"D=M\n" ,
			"@LCL\n" ,
			"M=D // LCL = *(FRAME-4)\n" ,
			"@return_address\n" ,
			"A=M\n" ,
			"0;JMP // goto RET\n")

	case "call":
		funcName := VMline[1]
		args := VMline[2]
		res = append(res, "// call" + funcName + args + "\n",
		"// push return-address\n",
        "@" + funcName + "$ret." + fmt.Sprintf("%d",line_number) + "\n",
		"D=A\n",
		"@SP\n",
		"A=M\n",
		"M=D\n",
		"@SP\n",
		"M=M+1\n",
		"// push LCL\n",
		"@LCL\n",
		"D=M\n",
		"@SP\n",
		"A=M\n",
		"M=D\n",
		"@SP\n",
		"M=M+1\n",
		"// push ARG\n",
		"@ARG\n",
		"D=M\n",
		"@SP\n",
		"A=M\n",
		"M=D\n",
		"@SP\n",
		"M=M+1\n",
		"// push THIS\n",
		"@THIS\n",
		"D=M\n",
		"@SP\n",
		"A=M\n",
		"M=D\n",
		"@SP\n",
		"M=M+1\n",
		"// push THAT\n",
		"@THAT\n",
		"D=M\n",
		"@SP\n",
		"A=M\n",
		"M=D\n",
		"@SP\n",
		"M=M+1\n",
		"// ARG = SP-n-5\n",
		"@SP\n",
		"D=M\n",
		"@" + args + "\n",
		"D=D-A\n",
		"@5\n",
		"D=D-A\n",
		"@ARG\n",
		"M=D\n",
		"// LCL = SP\n",
		"@SP\n",
		"D=M\n",
		"@LCL\n",
		"M=D\n",
		"// goto f\n",
		"@" + funcName + "\n",
		"0;JMP\n",
		"// (return-address)\n",
		"(" + funcName + "$ret." + fmt.Sprintf("%d", line_number) + ")\n")
	}
	return res
}
