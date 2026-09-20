package main

import "fmt"

func main() {
	var num1 int = 100

	switch num1 {
	case 98,99:
		fmt.Println("It's equal to 98 or 99")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98, 99 or 100")
	}
}

// switch i {
//     case 0: // 空分支，只有当 i == 0 时才会进入分支
//     case 1:
//         f() // 当 i == 0 时函数不会被调用!!
// }

// switch i {
//     case 0: fallthrough
//     case 1:
//         f() // 当 i == 0 时函数也会被调用!!
// }