package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

var content string
var vc VCard

func main() {
	// using a decoder:
	file, _ := os.Open("vcard.gob")
	defer file.Close()

	inReader := bufio.NewReader(file)// 创建一个读取器，从文件中读取gob编码的数据
	dec := gob.NewDecoder(inReader)// 创建一个解码器
	err := dec.Decode(&vc)// 解码gob编码的数据到vc结构体
	if err != nil {
		log.Println("Error in decoding gob")
	}
	fmt.Println(vc)
}