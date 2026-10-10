package even

import "testing"

// 测试文件名必须以 _test.go 结尾，测试用例都是自己写的，不需要 import 其他包

// 测试偶数
func TestEven(t *testing.T) {
	if !Even(10) {
		t.Log(" 10 must be even!")
		t.Fail()
	}
	if Even(7) {
		t.Log(" 7 is not even!")
		t.Fail()
	}
}

// 测试奇数
func TestOdd(t *testing.T) {
	if !Odd(11) {
		t.Log(" 11 must be odd!")
		t.Fail()
	}
	if Odd(10) {
		t.Log(" 10 is not odd!")
		t.Fail()
	}
}
