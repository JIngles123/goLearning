package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

func main() { // simple example
    pa := &Address{"private", "Aartselaar", "Belgium"}
    wa := &Address{"work", "Boom", "Belgium"}
    vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
    // fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
    // JSON format:
    js, _ := json.Marshal(vc)
    fmt.Printf("JSON format: %s", js)
    // using an encoder:
    file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
    defer file.Close()
    enc := json.NewEncoder(file)
    err := enc.Encode(vc)
    if err != nil {
        log.Println("Error in encoding json")
    }
}

func main1() { // more complex example
	pa:=&Address{"private", "Aartselaar", "Belgium"}
	wa:=&Address{"work", "Boom", "Belgium"}
	vc:=VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	// JSON format:
	js,_:=json.Marshal(vc) // 将内容序列化到内存中，方便打印输出
	fmt.Printf("JSON format: %s", js)

	// using an encoder:
	file,_:=os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	enc:=json.NewEncoder(file) // 创建一个编码器，将vc编码为JSON格式并写入文件
	err:=enc.Encode(vc) // 编码vc并写入文件
	if err!=nil {
		log.Println("Error in encoding json")
	}

	// using Unmarshal:（可删除，这里只是为了测试反序列化 Unmarshal 的功能）
	// 把 JSON 解码为数据结构
	var m VCard
	err = json.Unmarshal(js, &m) // 解析 [] byte 中的 JSON 数据 js 并将结果存入指针 &m 指向的值
	if err != nil {
		log.Println("Error in decoding json")
	}
	fmt.Printf("\nUnmarshal: %+v\n", m)

	// using a decoder:
	fileDec, err := os.Open("vcard.json")
	if err != nil {
		log.Println("Error opening vcard.json")
		return
	}
	defer fileDec.Close()

	var m2 VCard
	dec := json.NewDecoder(fileDec)
	err = dec.Decode(&m2)
	if err != nil {
		log.Println("Error in decoding json")
	}
	fmt.Printf("Decoder: %+v\n", m2)
}
