package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputError:=os.Open("input.dat") // 以只读模式打开文件
	if inputError!=nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
	        "Does the file exist?\n" + 
	        "Have you got access to it?\n")
		return
	}
	defer inputFile.Close() // 如果文件正常打开，使用这个语句确保在程序退出前关闭此文件

	inputReader:=bufio.NewReader(inputFile)
	for { // 将文件的内容逐行（行结束符为'\n'）读取，直到文件结束
		inputString, readerError:=inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError==io.EOF {
			return
		}
	}
}
