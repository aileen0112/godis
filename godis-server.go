package main

import (
	"fmt"
	"godis/info"
	"godis/networking"
	"godis/server"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// ConfigDefaultRdbFilename rdb filename
const ConfigDefaultRdbFilename = "dump.rdb"

// ConfigDefaultAofFilename aof filename
const ConfigDefaultAofFilename = "appendonly.aof"

// ConfigDefaultDBNum db nums
const ConfigDefaultDBNum = 16

// ConfigDefaultServerPort port
const ConfigDefaultServerPort = 9736

// ConfigDefaultDir working dir
const ConfigDefaultDir = "/tmp/"

func initDb() {
	fmt.Println("init db begin-->")
	godis.Db = make([]*server.GodisDb, godis.DbNum)
	for i := 0; i < godis.DbNum; i++ {
		godis.Db[i] = new(server.GodisDb)
		godis.Db[i].Dict = make(map[string]*server.GodisObject, 100)
	}
}

// Server : instance of godis-server
var godis = new(server.Server)

func main() {
	//var k = init()
	argv := os.Args
	argc := len(os.Args)
	if argc >= 2 {
		//j := 1 /* First option to parse in Args[] */
		/* Handle special options --help and --version */
		if argv[1] == "-v" || argv[1] == "--version" {
			version()
		}
		if argv[1] == "--help" || argv[1] == "-h" {
			usage()
		}
		if argv[1] == "--test-memory" {
			if argc == 3 {
				os.Exit(0)
			} else {
				println("Please specify the amount of memory to test in megabytes.\n")
				println("Example: ./godis-server --test-memory 4096\n\n")
				os.Exit(1)
			}
		}
	}
	godis = new(server.Server)
	initServerConfig()
	initServer()
	initDb()
	networking.LoadData(godis)
	log.Println("server init fin, ok", godis)

	c := make(chan os.Signal)
	//监听信号 ctrl+c kill
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go sigHandler(c)

	netListen, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(ConfigDefaultServerPort))
	if err != nil {
		log.Print("listen err ", "127.0.0.1:"+strconv.Itoa(ConfigDefaultServerPort))
	}
	//checkError(err)
	defer netListen.Close()
	log.Println("listen")

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		go networking.ProcessEvents(conn, godis)
	}
}

func handler(conn net.Conn, server *server.Server) {
	/*
		buff := make([]byte, 1024)
		for {
			n, err := conn.Read(buff)
			if err != nil {
				return
			}
			fmt.Println(n, conn.RemoteAddr().String(), conn.LocalAddr().String(), string(buff))
			//命令处理
			ret, err := do(string(buff), godis)
			//与客户端交互
			if v, ok := ret.(string); ok {
				conn.Write([]byte(v))
				fmt.Println([]byte(v))
			} else {
				//log.Fatal()
			}
			conn.Close()
		}
	*/
}

/*
func setCommand(server *Server, key string, value interface{}) error {
	server.db.Dict[key] = value
	fmt.Println(server.db.Dict, "server stat now in func setCommand", key, value)
	return nil
}
*/
func infoCommand(server *server.Server) (map[string]map[string]string, error) {
	serverInfo := info.GetServer()
	info := make(map[string]map[string]string)
	info["server"] = serverInfo
	return info, nil
}

func sigHandler(c chan os.Signal) {
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			exitHandler()
		case syscall.SIGUSR1:
			fmt.Println("signal usr1", s)
		case syscall.SIGUSR2:
			fmt.Println("signal usr2", s)
		default:
			fmt.Println("other signal", s)
		}
	}
}
func setupSignalHandlers() {
}

func exitHandler() {
	fmt.Println("开始退出...")
	fmt.Println("执行清理...")
	fmt.Println("结束退出...")
	os.Exit(0)
}
func start() {
	fmt.Println("hello godis")
}
func initServerConfig() {
	godis.Port = ConfigDefaultServerPort
	godis.DbNum = ConfigDefaultDBNum
	godis.RdbFilename = ConfigDefaultDir + ConfigDefaultRdbFilename
	godis.AofFilename = ConfigDefaultDir + ConfigDefaultAofFilename
	godis.NextClientID = 1 /* Client IDs, start from 1 .*/
}
func initServer() {
	setupSignalHandlers()

	godis.Pid = getpid()
	//server.clients = listCreate()
	//jserver.systemMemorysize = zmalloc_get_memory_size()

	//server.Db = db.InitDb()
	godis.Start = time.Now().UnixNano() / 1000000
	//var getf server.CmdFun

	type cmdFunc func(c *server.Client, s *server.Server)
	var gcc cmdFunc
	gcc = server.GetCommand
	fmt.Println(gcc)
	getCommand := &server.GodisCommand{Name: "get", Proc: server.GetCommand, Arity: -2}
	setCommand := &server.GodisCommand{Name: "set", Proc: server.SetCommand, Arity: -3}
	//getCommand := &server.GodisObject{ObjectType: 1, Ptr: gcc}
	//getV := map[string]*server.GodisCommand{}
	godis.Commands = map[string]*server.GodisCommand{
		"get": getCommand,
		"set": setCommand,
	}
	log.Println("server load data fin, ok")
	//server.AofBuf =
}

/*
func loadData() {
	c := networking.
	log.Println("file data loading ...")
	prefix := "/tmp/"
	pros := aof.FileToPro(prefix + ConfigDefaultAofFilename)
	//log.Println(pros, len(pros));os.Exit(0)
	for _, v := range pros {
		//do(v, godis)
		log.Println("loading... ", v)
	}
	log.Println("file data loading fin, ok")
}
*/
func liveSeconds(server server.Server) int64 {
	secM := time.Now().UnixNano()/1000000 - server.Start
	return secM / 1000
}
func liveDays(server server.Server) int64 {
	secM := time.Now().UnixNano()/1000000 - server.Start
	return secM / 1000 / 86400
}
func call(c *server.Client) {

}

/*
func (server *server.Server)() LoadData() {
	log.Println("file data loading ...")
	fileName := "/Users/zhen/go/dump.rdb"
	pros := FileToPro(fileName)
	log.Println(pros, len(pros))
	os.Exit(0)
	for _, v := range pros {
		exec(v, server)
	}
	log.Println("file data loading fin, ok")
}

func exec(pro string, server *server.Server) (retv interface{}, err error) {
	argv, argc := protocol.Protocol2Args(pro)
	log.Println("in func do *******", argv, argc)
	cmd := argv[0]
	switch {
	case len(cmd) == 3:
		if strings.Compare(cmd, "get") == 0 {
			command.GetCommand(argv, argc)
		}
	}
	if argc == 0 {
		log.Println("failed of do, pro: ", pro)
	} else {
		log.Println("file data loading result", pro, argv, argc)
		if argc == 3 && 0 == strings.Compare(argv[0], "set") {
			err := setCommand(server, argv[1], argv[2])
			AppendToFile("godis.rdb", pro)
			if err == nil {
				return 1, nil
			}
		} else if argc == 3 && 0 == strings.Compare(argv[0], "get") {
			retv, err := getCommand(server, argv[1])
			fmt.Println("get result ", retv)
			if err == nil {
				return retv, nil
			}
		} else if argc == 1 && 0 == strings.Compare(argv[0], "info") {
			retv, err := infoCommand(server)
			if err == nil {
				var outBuf string
				for k, v := range retv {
					outBuf += "\n# " + strings.ToUpper(k) + "\n"
					//fmt.Println("# ", strings.ToUpper(k))
					for innerk, innerv := range v {
						outBuf += innerk + ":" + innerv + "\n"
						//fmt.Println(innerk, ":", innerv)
					}
				}
				return outBuf, nil
			}
		}
	}
	return nil, err
}
*/
func getpid() int32 {
	return 1
}
func listCreate() {

}
func version() {
	println("Redis server v=0.0.1 sha=xxxxxxx:001 malloc=libc-go bits=64 ")
	os.Exit(0)
}
func usage() {
	println("Usage: ./godis-server [/path/to/redis.conf] [options]")
	println("       ./godis-server - (read config from stdin)")
	println("       ./godis-server -v or --version")
	println("       ./godis-server -h or --help")
	println("Examples:")
	println("       ./godis-server (run the server with default conf)")
	println("       ./godis-server /etc/redis/6379.conf")
	println("       ./godis-server --port 7777")
	println("       ./godis-server --port 7777 --slaveof 127.0.0.1 8888")
	println("       ./godis-server /etc/myredis.conf --loglevel verbose")
	println("Sentinel mode:")
	println("       ./godis-server /etc/sentinel.conf --sentinel")
	os.Exit(1)
}
