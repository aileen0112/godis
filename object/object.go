package object

// RedisObject object wrapper of godis data structure
type RedisObject struct {
	ObjectType int
	//Encoding   int
	Ptr interface{}
}
