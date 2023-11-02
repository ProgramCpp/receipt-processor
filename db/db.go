package db

type Db interface {
	Insert(interface{})string
}