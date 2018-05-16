/**
 * Created with IntelliJ IDEA.
 * User: szhmery
 * Date: 18-5-16
 * Time: 下午15:45
 * To change this template use File | Settings | File Templates.
 * Name:New()
 */
package main

import (
	"fmt"
)

//定义类型 Test为一个结构
//这个结构为0值
//这里相当如一个new(Test)分配一块清零的存储空间给类型Test
//他的作用将接收一个类型 *Test的值
type Test struct {
	fd      int
	name    string
	nepipe  int
}

//这里将返回一个值＊Test（暂理解为指针值吧）
func NewFile(fd int, name string) *Test {
	if fd < 0 {
		return nil
	}
	f := Test{fd, name, 2}
	return &f
}


//有时零值是不够好的，初始化构造函数是必要的
//这个时候我们就在里面新构
func NewFile2(fd int, name string) *Test {
	if fd < 0 {
		return nil
	}
	f := new(Test)
	f.fd = fd
	f.name = name
	f.nepipe = 0
	return f
}

//注意：以上例子中很多注模。它是个每次求值即生成新实例的表达式。
//如NewFile中 f := Test{fd, name, 2} 和return &f 这样就产生了两次新实例（暂理解为使用了两个内存空间吧）
//变量对应的存储空间在函数返回后仍然存在。每取一个组合字面的地址求值时都生成一个新实例，因此我们可以把最后两行合起来。
func NewFile3(fd int, name string) *Test {
	if fd < 0 {
		return nil
	}
	//这样我们就只产生了一次新实例
	//组合字面的Key必须按顺序给出并全部出现
	return &Test{fd, name,3}
}

func NewFile4(fd int, name string) *Test {
	if fd < 0 {
		return nil
	}

	//如果明确使用   Key:value 对应元素，初始化可用任意顺序，未出现的Key对应着零值或空
	//此例中把name放前面了，也没有定义nepipe
	return &Test{name:name,fd:fd}
}

func NewFile5(fd int, name string) *Test {
	if fd < 0 {
		return nil
	}

	//注意,如果一个组合字面一个Key也没有，此类型将为零值。
	//他的结果 &Test{} 和new(Test)是一样的。
	return &Test{}
}

func main() {
	//调用NewFile 看上面的注释应该清楚了得到的结果因该就是最开始定义的Test
	//结果会是fd,name,nepipe
	newfile:= NewFile(1,"Fuck you Test")
	fmt.Printf("示例五：fd>%d | name>%s | nepipe>%d \n",newfile.fd,newfile.name,newfile.nepipe)

	//调用NewFile2
	newfile2:= NewFile2(2,"Fuck you Test2")
	fmt.Printf("示例五：fd>%d | name>%s | nepipe>%d \n",newfile2.fd,newfile2.name,newfile2.nepipe)

	//调用NewFile3
	newfile3:= NewFile3(3,"Fuck you Test3")
	fmt.Printf("示例五：fd>%d | name>%s | nepipe>%d \n",newfile3.fd,newfile3.name,newfile3.nepipe)

	//调用NewFile4
	newfile4:= NewFile4(4,"Fuck you Test4")
	fmt.Printf("示例五：fd>%d | name>%s | nepipe>%d \n",newfile4.fd,newfile4.name,newfile4.nepipe)

	//调用NewFile5
	newfile5:= NewFile5(5,"Fuck you Test5")
	fmt.Printf("示例五：fd>%d | name>%s | nepipe>%d \n",newfile5.fd,newfile5.name,newfile5.nepipe)
}