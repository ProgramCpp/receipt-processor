package db

type Db interface {
	Put(interface{}) string
	Get(string) (interface{}, bool)
}
