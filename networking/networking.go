package networking

import (
	"godis/protocol"
	"godis/server"
	"log"
	"net"
)

// ClientPointer for tmp alias
type ClientPointer *server.Client

// ProcessEvents process event
func ProcessEvents(conn net.Conn, s *server.Server) {
	c := createClient(conn, s)
	buff := make([]byte, 1024)
	defer conn.Close()
	for {
		n, err := conn.Read(buff)
		if err != nil {
			return
		}
		readQueryFromClient(c, conn, s)
		writeToClient(conn, c)
		log.Println("read from client bytes", n)
	}
}
func createClient(conn net.Conn, s *server.Server) (c *server.Client) {
	c = new(server.Client)
	selectDb(c, 1, s)
	//id := atomicGetIncr(s.NextClientID,client_id,1)
	id := 1
	c.ID = int32(id)
	c.QueryBuf = ""
	return c
}
func selectDb(c *server.Client, id int32, s *server.Server) {
	/*
		if _, ok := s.Db[id]; ok {
			c.Db = s.Db[id]
		}
	*/
	c.Db = s.Db[id]
}
func readQueryFromClient(c ClientPointer, conn net.Conn, s *server.Server) {
	buff := make([]byte, 1024*1024)
	n, err := conn.Read(buff)
	if err != nil {
		return
	}
	log.Println(n, conn.RemoteAddr().String(), conn.LocalAddr().String(), string(buff))
	c.QueryBuf = string(buff)
	//命令处理
	processInputBuffer(c, s)
}

func processInputBuffer(c ClientPointer, s *server.Server) {
	//c->reqtype == PROTO_REQ_MULTIBULK, 暂时不考虑其他协议
	processMultibulkBuffer(c)
	server.ProcessCommand(c, s)
}

//解析协议 保存到argc argv
func processMultibulkBuffer(c ClientPointer) {
	argv, argc := protocol.Protocol2Args(c.QueryBuf)
	c.Argc = int32(argc)
	for _, v := range argv {
		c.Argv = append(c.Argv, stringObject(v))
	}
}
func stringObject(s string) (o *server.GodisObject) {
	o = new(server.GodisObject)
	o.ObjectType = 0
	o.Ptr = s
	return o
}

func writeToClient(conn net.Conn, c *server.Client) {
	conn.Write([]byte(c.Buf))
}
