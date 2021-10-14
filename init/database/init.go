package init

import (
	confCommon "common_notify_server/config/common"
	value "common_notify_server/res"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func init() {
	// todo: check db connectivity
	// todo: create table if not exist
	// todo: read data from table
	// todo: pass data to super
	log.Println("database initializing")
}

func DataBase() {
	switch confCommon.DBMethod {
	case value.MethodSqlite:
		sqlite()
		break
	case value.MethodMysql:
		panic("not implemented")
	}
}
