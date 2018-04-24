package aof

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"
)

var rdbFile = "godis.rdb"

//AppendToFile aof 持久化
func AppendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY|syscall.O_CREAT, 0644)
	fmt.Println(f, content, "write to aof")
	if err != nil {
		fmt.Println("aof file error" + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}

//FileToPro parse AOF file to db data
func FileToPro(fileName string) []string {
	// 以只写的模式，打开文件
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("aof file error" + err.Error())
	}
	defer f.Close()
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("aof file error" + err.Error())
	}
	//ret := bytes.Split(content, []byte("*"))
	ret := bytes.Split(content, []byte{'*'})
	//pros := make([]string, len(ret)-1)
	var pros = make([]string, len(ret)-1)
	for k, v := range ret[1:] {
		v := append(v[:0], append([]byte{'*'}, v[0:]...)...)
		//log.Println("convert result", tmp)
		pros[k] = string(v)
	}
	log.Println(len(pros), len(ret[1:]))
	return pros
}

/*
//FileToPro pas AOF file to db data
func FileToPro(fileName string) []string {
	// 以只写的模式，打开文件
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("aof file error" + err.Error())
	}
	defer f.Close()
	content, err := ioutil.ReadFile(fileName)
	log.Println(content)
	os.Exit(0)
	str := string(content)
	if err != nil {
		fmt.Println("aof file error" + err.Error())
	}
	//ret := bytes.Split(content, []byte("*"))
	ret := strings.Split(str, "*")
	//pros := new([]string)
	for k, v := range ret[1:] {
		log.Println("split result", k, v)
		//pros[k] = '*' + v
	}
	os.Exit(0)
	return ret[1:]
}
*/

/*
func readLines(fp os.File n int) string{
}
*/
// LoadData load aof data to server
/*
func (server *server.RedisServer)() LoadData() {
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

func exec(pro string, server *server.RedisServer) (retv interface{}, err error) {
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
