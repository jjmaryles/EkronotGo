package ex1
import (
	"fmt"
)
/*
func writeArithmetic(line []string) {
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
		} else if firstWord == "and"{m
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


}
 */

