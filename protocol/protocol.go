package protocol

import (
	"fmt"
	"strings"
)

//Cmd2Protocol cmd line to redis protocol
func Cmd2Protocol(cmd string) string {
	//todo trim useless space
	//cmd := "set alpha 12412"
	ret := strings.Split(cmd, " ")
	//todo validate cmd and params
	var pro string
	for k, value := range ret {
		if k == 0 {
			pro = fmt.Sprintf("*%d\r\n", len(value))
		}
		pro += fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)
	}
	return pro
}
