package main

import (
	"fmt"
	"godis/gmap"
	"list"
	"reflect"
	"sds"
)

func main() {
	gmap.Pt()
	var node = new(list.Node)
	ty := reflect.TypeOf(sds.Sds{})
	fmt.Println(ty.Size(), node.Next)
	var str = sds.Sds{Len: 100, Alloc: 200}
	fmt.Println("hello godis", str)
	sds.Pt(str)
}

func init() {
	fmt.Println("start ->")
}
