package parse

import (
	"fmt"
	"strings"
	"strconv"
)

type ParseError struct {
	Index int
	Word string
	Err error
}

func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r:=recover(); r!=nil { // 恢复panic异常,不影响main函数的执行，且defer语句会保证执行
			var ok bool
			err, ok=r.(error)
			if !ok {
				err=fmt.Errorf("pkg: %v", r)
			}
		}
	} ()
	
	fields:=strings.Fields(input)// 把字符串 input 按「空白字符」切分成多个子串，返回一个字符串切片
	numbers=fields2numbers(fields)
	return // 返回numbers和err，如果panic异常被恢复，err为nil，numbers为空
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields)==0 {
		panic("no words to parse") // 扔出 panic 异常
	}

	for idx, field:=range fields {
		num, err:=strconv.Atoi(field)
		if err!=nil {
			panic(&ParseError{idx, field, err}) // 扔出 panic 异常
		}
		numbers=append(numbers, num)
	}
	return
}
