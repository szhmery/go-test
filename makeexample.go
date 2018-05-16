package main

import "fmt"

func main() {

	//分配片结构;* p==零
	var p *[]int = new([]int)
	*p = make([]int, 100, 100) //这样写有点复杂，很容易就搞乱了
	fmt.Print("Test1:")
	fmt.Println(p)

	//现在将V分配一个新的数组，100个整型
	//写法一
	//var v  []int = make([]int, 100)
	//写法二：非常常用的写法，简节明了
	v := make([]int, 100)
	fmt.Print("Test2:")
	fmt.Println(v)

	//创建切片也使用make函数，它被分配一个零数组和指向这个数组的切片。

	//创建一个初始元素个数为5的数组切片，元素初始值为0

	a := make([]int, 5)  // len(a)=5
	fmt.Print("Test5:")
	fmt.Println(a)

	//切片有长度和容量。切片的最大长度就是它的容量。
	//指定一个切片的容量，通过第三个参数。

	//创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间
	b := make([]int, 5, 10)   	// len(b)=5, cap(b)=10

	//通过重新切片，可使切片增加。
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
	fmt.Print("Test4:")
	fmt.Println(b)

	//直接创建并初始化包含5个元素的数组切片
	c := []int{1,2,3,4,5}
	fmt.Print("Test5:")
	fmt.Println(c)
}