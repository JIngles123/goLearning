package main

import (
	"fmt"
	"bufio"
	"os"
	"compress/gzip"
)

// compress 包提供了读取压缩文件的功能，支持的压缩文件格式为：bzip2、flate、gzip、lzw 和 zlib。
func main() {
	fName:="MyFile.gz"
	var r *bufio.Reader
	fi, err:=os.Open(fName)
	if err!=nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName, err)
		os.Exit(1)
	}
	fz, err:=gzip.NewReader(fi)
	if err!=nil {
		r=bufio.NewReader(fi)
	} else {
		r=bufio.NewReader(fz)
	}

	for {
		line, err:=r.ReadString('\n')
		if err!=nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}