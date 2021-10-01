package db

type DBInstance struct{}

var dbInstance *DBInstance

func Init() {
	dbInstance = &DBInstance{}
}

func GetInstance() *DBInstance {
	return dbInstance
}
