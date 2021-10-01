package main

type DBInstance struct{}

var dbInstance *DBInstance

func GetInstance() *DBInstance {
	if dbInstance == nil {
		dbInstance = &DBInstance{}
	}
	return dbInstance
}

func main() {
	GetInstance()
}
