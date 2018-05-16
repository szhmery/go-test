package main
import (
	"bufio"
	"fmt"
	"net"
	"time"
	//"strconv"
	//"strings"
	//"strings"
)
const (
	// See http://golang.org/pkg/time/#Parse
	timeFormat = "2006-01-02 15:04:05.00000000"
)

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}
func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	for  {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println("From client's msg："+string(message))
		clientTimeStampStr := stringToTime(message)

		serverTimeStampStr := stringToTime(time.Now().String())
		fmt.Println("Client TimeStamp:"+clientTimeStampStr)
		fmt.Println("Server TimeStamp:"+serverTimeStampStr)

		clientTimeStamp, err := time.Parse(timeFormat, clientTimeStampStr)
		if err != nil {
			fmt.Println(err)
		}

		serverTimeStamp, err := time.Parse(timeFormat, serverTimeStampStr)
		if err != nil {
			fmt.Println(err)
		}


		gap := serverTimeStamp.Sub(clientTimeStamp)
		fmt.Print("Latency:")
		fmt.Println(gap)
	}
}

func stringToTime  (timestr string) string {
	//this is the way to get the correct time, no matter the length of time
	//arrary := strings.Split(timestr," ")
	//correcttimestr := arrary[0]+" "+arrary[1]

	//this is the way to get the last 8 bit number
	return timestr[0:28]
}
