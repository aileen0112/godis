package db

import (
	"fmt"
)

//v0.1 use map as dict (golang version >= 1.9)
//type dict=map

//RedisDb db struct
type RedisDb struct {
	Dict map[string]interface{}
}

//InitDb init redis Db
func InitDb() *RedisDb {
	fmt.Println("init db begin-->")
	dict := make(map[string]interface{}, 100)
	db := new(RedisDb)
	db.Dict = dict
	return db
}
