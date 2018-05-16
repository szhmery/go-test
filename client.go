package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
	"strconv"
)
var quitSemaphore chan bool

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer func () {
		fmt.Println("The client connection is close!")
		conn.Close()
	}()
	fmt.Println("client connected!")
	//go onMessageRecived(conn)
	onSendMsg(conn)
	//b := []byte("client time\n")
	//conn.Write(b)
	//<-quitSemaphore
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for i := 0; i < 10; i++ {
		fmt.Println("This is " + strconv.Itoa(i) + " time to read")
		msg, err := reader.ReadString('\n')
		fmt.Println("From sever's msg:"+msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
		time.Sleep(1*time.Microsecond)
		time.Sleep(1*time.Nanosecond)
		b := []byte(msg)
		conn.Write(b)

	}
	//quitSemaphore <- true
	close(quitSemaphore)
}

func onSendMsg(conn *net.TCPConn) {
	for i := 1; i <= 10; i++ {
		fmt.Println("Client sending " + strconv.Itoa(i) + " packets to server")
		msg := time.Now().String() + "\n"
		fmt.Println("Client send msg: "+msg)
		b := []byte(msg)
		conn.Write(b)
	}
	//quitSemaphore <- true
	//close(quitSemaphore)
}
