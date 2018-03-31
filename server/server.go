package main

import (
	"fmt"
	"godis/aof"
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
	db     *db.RedisDb
	dbnum  int
	AofBuf []string
}

func main() {
	//var k = init()
	server := initServer()
	log.Println("server init fin, ok")
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
		ret, err := do(string(buff), server)
		if v, ok := ret.(string); ok {
			conn.Write([]byte(v))
			fmt.Println([]byte(v))
		}
		conn.Close()
	}
}

func setCommand(server *RedisServer, key string, value interface{}) error {
	server.db.Dict[key] = value
	fmt.Println(server.db.Dict, "server stat now in func setCommand", key, value)
	return nil
}
func getCommand(server *RedisServer, key string) (interface{}, error) {
	v, ok := server.db.Dict[key]
	if !ok {
		return nil, nil
	}
	return v, nil
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
func start() {
	fmt.Println("hello godis")
}

func initServer() *RedisServer {
	server := new(RedisServer)
	server.db = db.InitDb()
	loadData(server)
	log.Println("server load data fin, ok")
	return server
	//server.AofBuf =
}

func loadData(server *RedisServer) {
	log.Println("file data loading ...")
	fileName := "/Users/zhen/go/dump.rdb"
	pros := aof.FileToPro(fileName)
	//log.Println(pros, len(pros));os.Exit(0)
	for _, v := range pros {
		do(v, server)
	}
	log.Println("file data loading fin, ok")
}
func do(pro string, server *RedisServer) (retv interface{}, err error) {
	argv, argc := protocol.Protocol2Args(pro)
	log.Println("in func do *******", argv, argc)
	if argc == 0 {
		log.Println("failed of do, pro: ", pro)
	} else {
		log.Println("file data loading result", pro, argv, argc)
		if argc == 3 && 0 == strings.Compare(argv[0], "set") {
			err := setCommand(server, argv[1], argv[2])
			aof.AppendToFile("godis.rdb", pro)
			if err == nil {
				return retv, nil
			}
		} else if argc == 3 && 0 == strings.Compare(argv[0], "get") {
			retv, err := getCommand(server, argv[1])
			fmt.Println("get result ", retv)
			if err == nil {
				return retv, nil
			}
		}
	}
	return nil, err
}
