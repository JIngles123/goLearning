package main
import "fmt"

func main() {
	array:=[3]float64{7.0, 8.5, 9.1}
	x:=Sum(&array)
	fmt.Printf("The sum of the array is: %f\n", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v:=range *a {
		sum+=v
	}
	return
}

// 把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：
// 1、传递数组的指针
// 2、使用数组的切片