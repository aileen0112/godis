package main

import (
	"bufio"
	"fmt"
	"godis/protocol"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	server := "127.0.0.1:2046"

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hi Godis")
	//fmt.Println("********************")

	for {
		fmt.Print(server + "> ")
		text, _ := reader.ReadString('\n')
		//clear 回车换行
		text = strings.Replace(text, "\n", "", -1)
		pro := protocol.Cmd2Protocol(text)
		//fmt.Println(pro)
		//validation of pro
		sendPro2Server(pro, server)

		//if strings.Compare("exit" text) == 0 {}
	}

}
func sendPro2Server(pro string, server string) (err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		log.Print("conn err, clent")
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(tcpAddr, "client failed", conn, err)
		return err
	}
	data := []byte(pro)
	conn.Write(data)
	return nil
}
