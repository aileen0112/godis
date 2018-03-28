package protocol

import (
	"fmt"
	"log"
	"strconv"
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

//Protocol2Args reverse of Cmd2Protocol
func Protocol2Args(protocol string) (argv []string, argc int) {
	parts := strings.Split(strings.Trim(protocol, " "), "\r\n")
	if len(parts) == 0 {
		//errors.New("invalid proto 1")
		log.Println("invalid")
	}
	argc, err := strconv.Atoi(parts[0][1:])
	if err != nil {
		//errors.New("invalid proto 2")
		log.Println("invalid")
	}
	j := 0
	var vlen []int
	fmt.Println(protocol, parts)
	for _, v := range parts[1:] {
		if len(v) == 0 {
			continue
		}
		//todo valid len of params
		if v[0] == '$' {
			tmpl, err := strconv.Atoi(v[1:])
			if err == nil {
				vlen = append(vlen, tmpl)
			}
		} else {
			fmt.Println("before :", vlen, v, argv)
			if j < len(vlen) && vlen[j] == len(v) {
				j++
				argv = append(argv, v)
			}
			fmt.Println("after :", vlen, v, argv)
		}
	}
	fmt.Println(argc, argv, parts)
	return argv, argc
}
