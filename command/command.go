package command

import (
	"fmt"
)

//GodisCommand command struct
type GodisCommand struct {
	Name  string
	Proc  func(argv []string, argc int)
	Arity int
}
type cmdFunc func(argv []string, argc int)

func SetCommand(argv []string, argc int) {
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
