package trans

import "math"

var Pi float64

func init() {
  Pi = 4 * math.Atan(1) // 默认包里的init函数会在包被导入时按import顺序，单线程自动执行
}