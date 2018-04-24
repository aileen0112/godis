package server

import (
	"fmt"
	"godis_back/list"
	"log"
	"os"
)

//Client connect client
type Client struct {
	Cmd      *GodisCommand
	Argv     []*GodisObject
	Argc     int32
	Db       *GodisDb
	ID       int32
	QueryBuf string
	Buf      string
}

//GodisCommand command struct
type GodisCommand struct {
	Name  string
	Proc  cmdFunc
	Arity int
}

//version 0.0.1
//type cmdFunc func(argv []string, argc int)
//version 0.0.1
type cmdFunc func(c *Client, s *Server)

// GodisObject object wrapper of godis data structure
type GodisObject struct {
	ObjectType int
	//Encoding   int
	Ptr interface{}
}

// Server struct
type Server struct {
	Db               []*GodisDb
	DbNum            int
	Start            int64
	Port             int32
	RdbFilename      string
	AofFilename      string
	NextClientID     int32
	SystemMemorySize int32
	Clients          int32
	Pid              int32
	commands         map[string]GodisCommand
	//AofBuf           string
	//info   map[string]interface{}
	AofBuf []string
}

//use map[string]* as type dict
type dict map[string]*GodisObject

//GodisDb db struct
type GodisDb struct {
	Dict    dict
	Expires dict
	ID      int32
}

/*
func SetCommand(argv []string, argc int) {
	if v := argv[0]; strings.Compare(v, "set") != 0 {
		fmt.Println("error set command")
	}
	//key := argv[1]
	fmt.Println(argv, argc, "setcommand")
}
*/

// SetSet cmd get
func SetSet() {
	fmt.Println("hello setset")
}

//func processCommand(c *Client) {}

// cmd map
//var cmdMap = map[string]cmdFunc{"set": SetCommand, "get": GetCommand}

/*
// InitCommand func
func InitCommand(argv []string, argc int) *GodisCommand {
	cmd := new(GodisCommand)
	if f := Search(argv[0]); f != nil {
		cmd.Name = argv[0]
		cmd.Proc = f
		cmd.Arity = 3
		fmt.Println(cmd)
		return cmd
	}
	return nil
}
*/

// SetCommand cmd of set
func SetCommand(key string, value interface{}) error {
	//server.Db.Dict[key] = value
	fmt.Println("server in func setCommand", key, value)
	return nil
}

// Search cmd of search
func Search(name string) {
	/*
			if v, exist := cmdMap[name]; exist {
				return v
			}
		return nil
	*/
}

// LpushCommand lpush command
func LpushCommand(server *Server, key string, value interface{}) error {
	list := list.Create()
	list.AddHead(value)
	obj := new(GodisObject)
	obj.ObjectType = 2
	obj.Ptr = list
	server.Db[0].Dict[key] = obj
	fmt.Println(server.Db[0].Dict, "server stat now in func lpushCommand", key, value)
	return nil
}

// Llen len of list
func Llen(server *Server, key string) int {
	return -1
}

// ProcessCommand process
func ProcessCommand(c *Client, s *Server) {
	v := c.Argv[0].Ptr
	name, ok := v.(string)
	if !ok {
		log.Println("error cmd")
		os.Exit(1)
	}
	c.Cmd = lookupCommand(name, s)
	call(c, s)
}
func lookupCommand(name string, s *Server) *GodisCommand {
	if cmd, ok := s.commands[name]; ok {
		return &cmd
	}
	return nil
}
func call(c *Client, s *Server) {
	c.Cmd.Proc(c, s)
}

//GetCommand get
func GetCommand(c *Client, s *Server) {
	o := lookupKey(c.Db, c.Argv[1])
	addReply(c, o)
}
func lookupKey(db *GodisDb, key *GodisObject) (ret *GodisObject) {
	o, ok := db.Dict[key.Ptr.(string)]
	if !ok {
		return o
	}
	return nil
}

func addReply(c *Client, o *GodisObject) {
	Obj2Protocol(c, o)
}

// Obj2Protocol ret obj to protocol
func Obj2Protocol(c *Client, o *GodisObject) {
	if c.Cmd.Name == "set" {
		c.Buf = "OK"
	} else if c.Cmd.Name == "get" {
		c.Buf = o.Ptr.(string)
	} else {
		c.Buf = "nil"
	}
}
