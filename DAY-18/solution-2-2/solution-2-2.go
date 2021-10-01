package main

import "sync"

type DBInstance struct{}

var (
	dbInstance *DBInstance
	once       = &sync.Once{}
)

func GetInstance() *DBInstance {
	once.Do(func() {
		dbInstance = &DBInstance{}
	})
	return dbInstance
}

func main() {
	GetInstance()
}
