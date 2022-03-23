package main

/*
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

 */
/*
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
*/