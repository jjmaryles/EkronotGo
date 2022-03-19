package main

import (
	"log"
	"os"
	"strconv"
)

//region locArgThisThat
//region local
func PushLocalX(command []string, file string) []string {
	writeFile, err1 := os.Open(file)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer writeFile.Close()

	var res []string
	res = append(res, "@"+command[2])
	res = append(res, "D=A")
	res = append(res, "@LCL")
	res = append(res, "A=M+D")
	res = append(res, "D=M")
	res = append(res, "@SP")
	res = append(res, "A=M")
	res = append(res, "M=D")
	res = append(res, "@SP")
	res = append(res, "M=M+1")

	return res
}

//endregion local

//region argument
func PushArgumentX(command []string, file string) []string {
	writeFile, err1 := os.Open(file)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer writeFile.Close()

	var res []string
	res = append(res, "@ARG")
	res = append(res, "D=A")
	res = append(res, "@"+command[2])
	res = append(res, "D=D+A")
	res = append(res, "A=D")
	res = append(res, "@SP")
	res = append(res, "M=D")
	res = append(res, "@SP")
	res = append(res, "M=M+1")

	return res
}

//endregion argument

//region this
func PopThisX(command []string, file string) []string {
	writeFile, err1 := os.Open(file)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer writeFile.Close()

	var res []string
	res = append(res, "@SP")
	res = append(res, "A=M-1")
	res = append(res, "D=M")
	res = append(res, "@THIS")
	res = append(res, "A=M")

	num, _ := strconv.Atoi(command[2])
	for i := 0; i < num; i++ {
		res = append(res, "A=A+1")
	}

	res = append(res, "M=D")
	res = append(res, "@SP")
	res = append(res, "M=M-1")

	return res
}

//endregion this

//endregion local,argument,this,that

//region temp
//endregion temp

//region static
//endregion static

//region pointer
//endregion pointer

//region constant
func PushConstantX(command []string, file string) []string {
	writeFile, err1 := os.Open(file)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer writeFile.Close()

	var res []string
	res = append(res, "@"+command[2])
	res = append(res, "D=A")
	res = append(res, "@SP")
	res = append(res, "A=M")
	res = append(res, "M=D")
	res = append(res, "@SP")
	res = append(res, "M=M+1")

	return res
}

//endregion constant
