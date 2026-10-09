package main

import (
	"bytes"
	"fmt"
	"encoding/gob"
	"log"
)

type P struct {
	X,Y,Z int
	Name string
}

type Q struct {
	X,Y *int32
	Name string
}

// gob 编解码的例子，并以 字节缓冲 模拟 网络传输
func main() {
	var network bytes.Buffer
	enc:=gob.NewEncoder(&network)
	dec:=gob.NewDecoder(&network)

	// 编码 P 结构体
	err:=enc.Encode(P{3,4,5,"Pythagoras"})
	if err!=nil {
		log.Fatal("encode error:", err)
	}

	// 解码 Q 结构体
	var q Q
	err=dec.Decode(&q)
	if err!=nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
}

