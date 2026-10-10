package even // 红标只是提示，go 规定，一个目录名=一个包

func Even(i int) bool {
	return i%2==0
}

func Odd(i int) bool {
	return i%2!=0
}
