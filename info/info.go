package info

import (
	"godis/server"
	"os"
	"strconv"
)

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
	return server.TCPPORT
}
func getIP() string {
	return server.TCPIP
}
