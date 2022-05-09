package main

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