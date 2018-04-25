package main

import (
	"bufio"
	"fmt"
	"godis/protocol"
	"godis/server"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

//GodisClient connect client
type GodisClient struct {
	Cmd  *server.GodisCommand
	Argv []string
	Argc int
}

/*
func InitClient(argv []string, argc int) GodisClient {
	//cmd = command.Search(argv[0])
	cmd := command.InitCommand(argv, argc)
	fmt.Println(cmd, "initcommand of initclient")
	client := GodisClient{cmd, argv, argc}
	return client
}
*/

func main() {
	server := "127.0.0.1:9736"

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hi Godis")

	for {
		fmt.Print(server + "> ")
		text, _ := reader.ReadString('\n')
		//clear 回车换行
		text = strings.Replace(text, "\n", "", -1)
		pro := protocol.Cmd2Protocol(text)
		//fmt.Println(pro)
		tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
		if err != nil {
			log.Print("conn err, clent")
		}
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			fmt.Println(tcpAddr, "client failed", conn, err)
		}
		//validation of pro
		sendPro2Server(pro, server, conn)
		// take response of server
		result, err := ioutil.ReadAll(conn)
		checkError(err)
		//fmt.Println("result ", string(result))
		if len(result) == 0 {
			fmt.Println(server+"> ", "nil")
		} else {
			fmt.Println(server+">", string(result))
		}

		//if strings.Compare("exit" text) == 0 {}
	}

}
func sendPro2Server(pro string, server string, conn net.Conn) (err error) {
	data := []byte(pro)
	conn.Write(data)
	return nil
}
func checkError(err error) {
	if err != nil {
		log.Println("socket error: ", err.Error())
		os.Exit(1)
	}
}
