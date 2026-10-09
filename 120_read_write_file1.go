package main

import (
	"fmt"
	"os"
)

// 将整个文件的内容读到一个字符串里
func main() {
	inputFile:="input.dat"
	outputFile:="input_copy.txt"
	buf, err:=os.ReadFile(inputFile)
	if err!=nil {
		fmt.Fprintf(os.Stderr, "File error: %s\n", err)
	}

	fmt.Printf("%s\n", string(buf))
	err=os.WriteFile(outputFile, buf, 0644)
	if err!=nil {
		panic(err.Error())
	}
}
