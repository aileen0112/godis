package aof

import (
	"fmt"
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
