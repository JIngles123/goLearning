package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

var t, token xml.Token
var err1 error

func main() {
	input:="<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputReader:=strings.NewReader(input)
	p:=xml.NewDecoder(inputReader)

	// Token() returns the next XML token in the input stream
	for t,err1=p.Token(); err1==nil; t,err1=p.Token() { // 读到输入流的末尾时，Token()返回io.EOF错误
		switch token:=t.(type) {
		case xml.StartElement:
			name:=token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _,attr:=range token.Attr {
				attrName:=attr.Name.Local
				attrValue:=attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
				// ...
			}
		case xml.EndElement:
			fmt.Println("End of token, token name: ", token.Name.Local)
		case xml.CharData:
			content:=string([]byte(token)) // 将xml里的文本节点转成可打印的String
			fmt.Printf("This is the content: %v\n", content)
			// ...
		default:
			// ...
		}
	}
}
