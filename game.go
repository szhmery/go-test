package main


import (
"bufio"
"fmt"
"math/rand"
"os"
"strconv"
"time"
)
var (
	endNum int //设置成生数的范围
)
func main() {
	i := createRandomNumber(endNum)
	//fmt.Println("成生定规范围内的整数:", i)    //本句调试用


	fmt.Println("请输入整数,范围为:0-", endNum)


	flag := true
	reader := bufio.NewReader(os.Stdin)


	for flag {
		data, _, _ := reader.ReadLine()
		//fmt.Print(a,b)

		command, err := strconv.Atoi(string(data)) //string to int,并作输入式格判断
		if err != nil {
			fmt.Println("式格不对，请输入数字")
		} else {


			fmt.Println("你输入的数字:", command)


			if command == i {
				flag = false
				fmt.Println("祝贺你，答对了~")
			} else if command < i {
				fmt.Println("你输入的数字小于成生的数字，别悲观！再来一次~")
			} else if command > i {
				fmt.Println("你输入的数字大于成生的数字，别悲观！再来一次~")
			}
		}
	}
}
func init() {
	endNum = 10
}


//成生定规范围内的整数
//设置肇端数字范围，0开始,endNum截止
func createRandomNumber(endNum int) int {

	r := rand.New(rand.NewSource(time.Now().Unix()))

	return r.Intn(endNum)
}
