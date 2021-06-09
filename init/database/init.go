package nfly

import (
	config "common_notify_server/config/main"
	util "common_notify_server/pkg/util"
	value "common_notify_server/res"
	"fmt"
)

func init() {
	// todo: check db connectivity
	switch config.Method {
		case value.MethodSqlite:
			if !util.Exist(config.Path){
				panic("db file does not exist")
			}
			fmt.Println("connect method: sqlite => db file exist")
		break
		case value.MethodMysql:
			break
	}
	// todo: create table if not exist
	// todo: read data from table
	// todo: pass data to super
}





