package main
import (
	"strings"
)
func writeArithmetic(line []string) {
	firstWord := line[0]
	var res []string
	var strTemp string
	switch firstWord{
	case "add","sub","and","or":
		if firstWord == "add"{
			res[0] = "D=D+M"
		} else if firstWord == "sub"{
			res[0] = "D=M-D"
		} else if firstWord == "and"{
			res[0] = "D=D&M"
		} else if firstWord == "or"{
			res[0] = "D=D|M"
		}
		res = append(res, "@SP\n")
		res = append(res, "M=M-1\n")
		res = append(res, "A=M\n")
		res = append(res, "D=M\n")
		res = append(res, "A=A-1\n")
		strTemp = strings.Join(res, "\n")
		res = append(res,  strTemp)
		res = append(res, "M=D\n")
		res = append(res, "D=A+1\n")
		res = append(res, "@SP\n")
		res = append(res, "M=D\n")

	case "neg", "not":
		if firstWord == "neg" {
			res = append(res, "M=-M")
		}else{
			res = append(res,"M!=M")
		}
		res = append(res, "@SP\n")
		res = append(res, "M=M-1\n")
		res = append(res, "A=M\n")
		strTemp = strings.Join(res, "\n")
		res = append(res,  strTemp)
		res = append(res, "D=A+1\n")
		res = append(res, "@SP\n")
		res = append(res, "M=D\n")
	}
}

