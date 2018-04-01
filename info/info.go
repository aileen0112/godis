package info

import (
	"os"
	"strconv"
)

//TCPIP ip
const TCPIP = "127.0.0.1"

//TCPPORT port
const TCPPORT = 2046

//GetServer server section of info cmd
func GetServer() map[string]string {
	pid := getPid()
	port := getPort()
	info := make(map[string]string)
	info["process_id"] = strconv.Itoa(pid)
	info["tcp_port"] = strconv.Itoa(port)
	return info
}

func getPid() int {
	return os.Getpid()
}
func getPort() int {
	return TCPPORT
}
