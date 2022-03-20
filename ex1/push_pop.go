package main

import (
	"strconv"
)

//region group1-locArgThisThat
//region local
func PushLocalX(command []string) []string {
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
func PushArgumentX(command []string) []string {

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
func PopThisX(command []string) []string {

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

//region that
func PushThatX(command []string) []string {


	var res []string
	res = append(res, "@"+command[2])
	res = append(res, "D=A")
	res = append(res, "@THAT")
	res = append(res, "A=M+D")
	res = append(res, "D=M")
	res = append(res, "@SP")
	res = append(res, "A=M")
	res = append(res, "M=D")
	res = append(res, "@SP")
	res = append(res, "M=M+1")

	return res
}
//endregion that
//endregion group1-locArgThisThat

//region group2-temp
func PopTempX(command []string) []string {

	var res []string
	res = append(res, "@SP")
	res = append(res, "A=M-1")
	res = append(res, "D=M")

	num, _ := strconv.Atoi(command[2])
	res = append(res, "@"+strconv.Itoa(5+num))

	res = append(res, "M=D")
	res = append(res, "@SP")
	res = append(res, "M=M-1")

	return res
}
//endregion group2-temp

//region group3-static
func PopStaticX(command []string, fileName string) []string {

	var res []string
	res = append(res, "@SP")
	res = append(res, "A=M-1")
	res = append(res, "D=M")
	res = append(res, "@" + fileName + "." + command[2])
	res = append(res, "M=D")
	res = append(res, "@SP")
	res = append(res, "M=M-1")

	return res
}

func PushStaticX(command []string, fileName string) []string {

	var res []string
	res = append(res, "@" + fileName + "." + command[2])
	res = append(res, "D=M")
	res = append(res, "@SP")
	res = append(res, "A=M")
	res = append(res, "M=D")
	res = append(res, "@SP")
	res = append(res, "M=M+1")

	return res
}
//endregion group3-static

//region group4-pointer0,1
func PopPointer0(command []string) []string {

	var res []string
	res = append(res, "@SP")
	res = append(res, "A=M-1")
	res = append(res, "D=M")
	res = append(res, "@THIS")
	res = append(res, "M=D")
	res = append(res, "@SP")
	res = append(res, "M=M-1")

	return res
}

func PopPointer1(command []string) []string {

	var res []string
	res = append(res, "@SP")
	res = append(res, "A=M-1")
	res = append(res, "D=M")
	res = append(res, "@THAT")
	res = append(res, "M=D")
	res = append(res, "@SP")
	res = append(res, "M=M-1")

	return res
}

func PushPointer0(command []string) []string {

	var res []string
	res = append(res, "@THIS")
	res = append(res, "D=M")
	res = append(res, "@SP")
	res = append(res, "A=M")
	res = append(res, "M=D")
	res = append(res, "@SP")
	res = append(res, "M=M+1")

	return res
}

func PushPointer1(command []string) []string {

	var res []string
	res = append(res, "@THAT")
	res = append(res, "D=M")
	res = append(res, "@SP")
	res = append(res, "A=M")
	res = append(res, "M=D")
	res = append(res, "@SP")
	res = append(res, "M=M+1")

	return res
}
//endregion group4-pointer0,1

//region group5-constant
func PushConstantX(command []string) []string {

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

//endregion group5-constant
