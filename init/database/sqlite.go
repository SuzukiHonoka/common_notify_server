package init

import (
	"log"
	conf "nfly/config/database"
	helper "nfly/internal/database/sqlite"
	"nfly/internal/user"
	"nfly/internal/utils"
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
