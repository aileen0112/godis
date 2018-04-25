package server

import (
	"fmt"
	"godis/aof"
	"godis/list"
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
	Commands         map[string]*GodisCommand
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

// SetSet cmd get
func SetSet() {
	fmt.Println("hello setset")
}

//func processCommand(c *Client) {}

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
func SetCommand(c *Client, s *Server) {
	objKey := c.Argv[1]
	objValue := c.Argv[2]
	if stringKey, ok1 := objKey.Ptr.(string); ok1 {
		if stringValue, ok2 := objValue.Ptr.(string); ok2 {
			c.Db.Dict[stringKey] = stringObject(stringValue)
		}
	}
	addReply(c, objKey)
	//server.Db.Dict[key] = stringObject(string(objValue.Ptr)
	fmt.Println("func setCommand", c.Db, s.Db[0])
}
func stringObject(s string) (o *GodisObject) {
	o = new(GodisObject)
	o.ObjectType = 0
	o.Ptr = s
	return o
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
	cmd := lookupCommand(name, s)
	fmt.Println("lookup result ", cmd)
	if cmd != nil {
		c.Cmd = cmd
		fmt.Println("ProcessCommand", c.Cmd, s.Db)
		call(c, s)
	}
}
func lookupCommand(name string, s *Server) *GodisCommand {
	fmt.Println("lookupCommand", name, s.Commands)
	if cmd, ok := s.Commands[name]; ok {
		return cmd
	}
	return nil
}
func call(c *Client, s *Server) {
	fmt.Println("server call ", c.Cmd.Name, c.Cmd.Proc, s.Db)
	c.Cmd.Proc(c, s)
	aof.AppendToFile(s.AofFilename, c.QueryBuf)
}

//GetCommand get
func GetCommand(c *Client, s *Server) {
	o := lookupKey(c.Db, c.Argv[1])
	fmt.Println("GetCommand ", o)
	if o != nil {
		addReply(c, o)
	} else {
		addReplyNil(c)
	}
}
func lookupKey(db *GodisDb, key *GodisObject) (ret *GodisObject) {
	o, ok := db.Dict[key.Ptr.(string)]
	fmt.Println("lookupKey ", o, ok, db.Dict, key.Ptr)
	if ok {
		fmt.Println("ok, key exist ", o)
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
func addReplyNil(c *Client) {
	c.Buf = "nil"
}
