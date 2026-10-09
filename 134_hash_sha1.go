package main

import (
	"fmt"
	"crypto/sha1"
	"io"
	"log"
)

// 通过 io.WriteString 或 hasher.Write 将给定的 [] byte 附加到当前的 hash.Hash 对象中。

func main() {
	// 1) 创建一个 SHA1 计算器（实现了 hash.Hash 接口）
	hasher := sha1.New()

	// 2) 往计算器里写入要哈希的数据（可以多次 Write，效果是拼接后再算）
	io.WriteString(hasher, "test") // 将字符串写入hash计算器

	// 3) Sum 用来取出最终的哈希结果（20 字节）
	//    参数 b 是“前缀”：结果 = append(b, 哈希值...)
	//    这里传空切片，表示只要纯哈希值
	b := []byte{}
	fmt.Printf("Result: %x\n", hasher.Sum(b)) // %x：用十六进制打印，最常见的哈希展示方式
	fmt.Printf("Result: %d\n", hasher.Sum(b)) // %d：按每个字节的十进制数打印，一般只用来对比格式

	// 4) Reset 清空内部状态，同一个 hasher 可以重新用来算另一段数据
	hasher.Reset()

	// 5) 换一种写入方式：直接 Write []byte（和上面 WriteString 作用类似）
	data := []byte("We shall overcome!")
	n, err := hasher.Write(data) // 将字节切片写入hash计算器
	// Write 返回“实际写入字节数”和错误；正常时应写完整段 data
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}

	// 6) 再次取出哈希值（校验和 / checksum）
	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
}
