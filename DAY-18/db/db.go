package db

type DBInstance struct{}

var dbInstance *DBInstance

func init() {
	dbInstance = &DBInstance{}
}

func GetInstance() *DBInstance {
	return dbInstance
}
