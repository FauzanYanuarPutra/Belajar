package database

var connect string

func init() {
	connect = "MySQL"
}

func ConnectDatabase() string {
	return connect
}