package main

import (
	"fmt"
	"time"
	//"regexp"
	//"os/exec"
	//"regexp"
	"strings"
)
const (
	// See http://golang.org/pkg/time/#Parse
	timeFormat = "2006-01-02 15:04:05.000000000"
)
func main() {

	//caution : format string is `2006-01-02 15:04:05.000000000`
	current := time.Now()

	fmt.Println("origin : ", current.String())
	// origin :  2016-09-02 15:53:07.159994437 +0800 CST

	fmt.Println("mm-dd-yyyy : ", current.Format("01-02-2006"))
	// mm-dd-yyyy :  09-02-2016

	fmt.Println("yyyy-mm-dd : ", current.Format("2006-01-02"))
	// yyyy-mm-dd :  2016-09-02

	// separated by .
	fmt.Println("yyyy.mm.dd : ", current.Format("2006.01.02"))
	// yyyy.mm.dd :  2016.09.02


	fmt.Println("yyyy-mm-dd HH:mm:ss : ", current.Format("2006-01-02 15:04:05"))
	// yyyy-mm-dd HH:mm:ss :  2016-09-02 15:53:07

	// StampMicro
	fmt.Println("yyyy-mm-dd HH:mm:ss: ", current.Format("2006-01-02 15:04:05.000000"))
	// yyyy-mm-dd HH:mm:ss:  2016-09-02 15:53:07.159994

	//StampNano
	fmt.Println("yyyy-mm-dd HH:mm:ss: ", current.Format("2006-01-02 15:04:05.000000000"))
	// yyyy-mm-dd HH:mm:ss:  2016-09-02 15:53:07.159994437
	//fmt.Println("yyyy-mm-dd HH:mm:ss: ", timestr.Format("2006-01-02 15:04:05.000000000"))
	//timestr := current.String()
	//2018-03-25 20:01:01.540574484 +0800 CST m=+0.000450648
	//[0-9]{4}-[0-9]]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}.[0-9]{9}).*
	//correcttime := timestr[0:29]
	//fmt.Println("Then:"+correcttime)
	//
	//then, err := time.Parse(timeFormat, correcttime)
	//if err != nil {
	//	fmt.Println(err)
	//
	//}
	//fmt.Println(then)
	//fmt.Println("Latency：")
	//time.Sleep(time.Microsecond)
	//fmt.Println("Now:", time.Now().String())
	//now := time.Now().Format("2006-01-02 15:04:05.000000000")
	//fmt.Println("now after format"+now)
	//fmt.Println(time.Now())



	before := time.Now()
	fmt.Println("before:"+before.String())
	//time.Sleep(2*time.Second)
	after := time.Now()
	fmt.Println("after :"+after.String())
	gap := after.Sub(before)
	fmt.Print("Gap: ")
	fmt.Println(gap)

	before1str := before.String()[0:29]
	before1, err := time.Parse(timeFormat, before1str)
	if err != nil {
		fmt.Println(err)
	}
	after1str := after.String()[0:29]
	after1, err := time.Parse(timeFormat,  after1str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("before1")
	fmt.Println(before1)
	fmt.Print("after1")
	fmt.Println(after1)
	gap1 := after1.Sub(before1)
	fmt.Print("gap1")
	fmt.Println(gap1)

	timestr := "2018-03-25 22:04:27.18205301 +0800 CST m=+0.022590237"

	//r := regexp.MustCompile(`^([\d]{4}[\w-][\d]{2}[\w-][\d]{2}\s[\d]{2}:[\d]{2}:[\d]{2}:[\d]{0,9})\s.*$`)
	// 需要匹配的字符串
	//r := regexp.MustCompile(`^([\w.\-+]+\s[\w.\-+]+)\s\+0800\sCST\s[\w.+]+.*$`)
	arrary := strings.Split(timestr," ")
	time := arrary[0]+" "+arrary[1]

	// 返回所有匹配的字符串
	//strings := r.FindAllString(timestr, -1)
	fmt.Println(time)
}
