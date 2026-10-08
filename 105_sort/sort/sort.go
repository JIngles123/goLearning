package sort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// data Sorter：只要实现了 Len、Less、Swap，就能传进来。函数体里只通过这三个方法访问数据，不关心里面是 []int、[]string 还是别的。
// 所以 IntArray、StringArray 都能交给同一个 Sort (interface 的多态性)
// sort.Sort(sort.IntArray(data))
// sort.Sort(sort.StringArray(data))
func Sort(data Sorter) {
	for pass:=1; pass<data.Len(); pass++ { // data.Len(),其中 data 就是函数接受者的类型
		for i:=0; i<data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n:=data.Len()
	for i:=n-1; i>0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

type IntArray []int
func (p IntArray) Len() int {return len(p)} // (p IntArray)是方法接收者。Len 等方法属于 IntArray，调用时是IntArray.Len()
func (p IntArray) Less(i, j int) bool {return p[i]<p[j]}
func (p IntArray) Swap(i, j int) {p[i], p[j]=p[j], p[i]}

type StringArray []string
func (p StringArray) Len() int {return len(p)}
func (p StringArray) Less(i,j int) bool {return p[i]<p[j]}
func (p StringArray) Swap(i,j int) {p[i], p[j]=p[j], p[i]}

func SortInts(a []int) { Sort(IntArray(a))}
func SortStrings(a []string) { Sort(StringArray(a))}

func IntsAreSorted(a []int) bool {return IsSorted(IntArray(a))}
func StringsAreSorted(a []string) bool {return IsSorted(StringArray(a))}

