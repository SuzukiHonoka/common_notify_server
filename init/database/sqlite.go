package init

import (
	conf "common_notify_server/config/database"
	helper "common_notify_server/internal/database/sqlite"
	"common_notify_server/internal/user"
	"common_notify_server/internal/utils"
	"log"
)

func sqlite() {
	log.Println("connect method: sqlite")
	user.Helper = new(helper.Helper)
	if !utils.PathExist(conf.DBPath) {
		// panic("db file does not exist")
		log.Println("creating new DB")
		// create sqlite file and init table
		user.Helper.Create()
		return
	}
	user.Helper.Connect()
	user.Helper.Refresh()
}
