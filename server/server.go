package main

import (
	"fmt"
	"godis/db"
	"godis/protocol"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

//RedisServer struct
type RedisServer struct {
	db    *db.RedisDb
	dbnum int
}

func main() {
	//var k = init()
	server := new(RedisServer)
	server.db = start()
	fmt.Println(server, "server")
	netListen, err := net.Listen("tcp", "127.0.0.1:2046")
	if err != nil {
		log.Print("listen err")
	}
	//checkError(err)
	defer netListen.Close()
	log.Println("listen")

	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go sig(c)

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		go handler(conn, server)
	}
}

func handler(conn net.Conn, server *RedisServer) {
	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			return
		}
		fmt.Println(n, conn.RemoteAddr().String(), conn.LocalAddr().String(), string(buff))
		do(string(buff), server)
	}
}

func do(pro string, server *RedisServer) {
	argv, argc := protocol.Protocol2Args(pro)
	if argc == 3 && 0 == strings.Compare(argv[0], "set") {
		setCommand(server, argv[1], argv[2])
	} else if argc == 3 && 0 == strings.Compare(argv[0], "get") {
		get := getCommand(server, argv[1])
		fmt.Println("get result ", get)
	}
}

func setCommand(server *RedisServer, key string, value interface{}) {
	server.db.Dict[key] = value
	fmt.Println(server.db.Dict, "server stat now in func do")
}
func getCommand(server *RedisServer, key string) interface{} {
	v, ok := server.db.Dict[key]
	if !ok {
		return nil
	}
	return v
}

func start() *db.RedisDb {
	db := db.InitDb()
	return db
}

func sig(c chan os.Signal) {
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("退出", s)
			ExitFunc()
		case syscall.SIGUSR1:
			fmt.Println("usr1", s)
		case syscall.SIGUSR2:
			fmt.Println("usr2", s)
		default:
			fmt.Println("other", s)
		}
	}
}

//ExitFunc exit smoothly
func ExitFunc() {
	fmt.Println("开始退出...")
	fmt.Println("执行清理...")
	fmt.Println("结束退出...")
	os.Exit(0)
}
