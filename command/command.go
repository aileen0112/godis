package command

import (
	"fmt"
	"godis/list"
	"godis/object"
	"godis/server"
	"strings"
)

//GodisCommand command struct
type GodisCommand struct {
	Name  string
	Proc  func(argv []string, argc int)
	Arity int
}
type cmdFunc func(argv []string, argc int)

func SetCommand(argv []string, argc int) {
	if v := argv[0]; strings.Compare(v, "set") != 0 {
		fmt.Println("error set command")
	}
	//key := argv[1]
	fmt.Println(argv, argc, "setcommand")
}
func GetCommand(argv []string, argc int) {
	fmt.Println(argv)
}

var cmdMap = map[string]cmdFunc{"set": SetCommand, "get": GetCommand}

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

func Search(name string) cmdFunc {
	if v, exist := cmdMap[name]; exist {
		return v
	}
	return nil
}

// LpushCommand lpush command
func LpushCommand(server *server.RedisServer, key string, value interface{}) error {
	list := list.Create()
	list.AddHead(value)
	obj := new(object.RedisObject)
	obj.ObjectType = 2
	obj.Ptr = list
	server.Db.Dict[key] = obj
	fmt.Println(server.Db.Dict, "server stat now in func lpushCommand", key, value, server.Db.Dict)
	return nil
}
