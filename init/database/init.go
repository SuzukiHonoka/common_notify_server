package init

import (
	confCommon "common_notify_server/config/common"
	database "common_notify_server/internal/database/common"
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
	case database.MethodSqlite:
		sqlite()
		break
	case database.MethodMysql:
		panic("not implemented")
	}
}
