package main

import (
	"encoding/gob"
	"log"
	"os"
	"fmt"
)

type Address struct {
	Type string
	City string
	Country string
}

type VCard struct {
	FirstName string
	LastName string
	Addresses []*Address
	Remark string
}

var content string

// 编码到文件
func main() {
	pa:=&Address{"private", "Aartselaar", "Belgium"}
	wa:=&Address{"work", "Boom", "Belgium"}
	vc:=VCard{"Jan", "Kersschot", []*Address{pa,wa}, "none"}
	fmt.Printf("%v: \n", vc)

	file,_:=os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	enc:=gob.NewEncoder(file)// 创建一个编码器，将vc编码为gob格式并写入文件
	err:=enc.Encode(vc)// 编码vc并写入文件
	if err!=nil {
		log.Println("Error in encoding gob")
	}
}
