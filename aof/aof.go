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

//FileToPro pas AOF file to db data
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
	/*
		for k, v := range pros {
			tmp := []byte(v)
			log.Println("--------", k)
			log.Println("--------", len(tmp), []byte(v))
		}
	*/
	//os.Exit(0)
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
