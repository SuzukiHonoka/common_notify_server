package sqlite

import (
	config "common_notify_server/config/database"
	database "common_notify_server/internal/database/common"
	"common_notify_server/internal/user"
	"common_notify_server/internal/utils"
	"common_notify_server/res"
	"database/sql"
	"log"
	"os"
)

type Helper struct {
	DB *sql.DB
}

func (h *Helper) Connect() {
	// not to defer close due to other reference
	var err error
	h.DB, err = sql.Open("sqlite3", config.DBPath)
	utils.CheckErrors(err)
}

func (h *Helper) Close() {
	_ = h.DB.Close()
}

func (h *Helper) Create() {
	_, err := os.Create(config.DBPath)
	utils.CheckErrors(err)
	h.DB, err = sql.Open("sqlite3", config.DBPath)
	utils.CheckErrors(err)
	_, err = h.DB.Exec(res.CreateTableIfNotExist)
	utils.CheckErrors(err)
}

func (h *Helper) GetUsers() interface{} {
	row, err := h.DB.Query(res.SelectAllUserFromDB)
	utils.CheckErrors(err)
	defer database.RowClose(row)
	var users []*user.USER
	for row.Next() {
		var email string
		var password string
		var group user.Type
		err = row.Scan(&email, &password, &group)
		utils.CheckErrors(err)
		users = append(users, &user.USER{
			Credit: user.Credit{
				Email:    email,
				Password: password,
			},
			Group: *user.AdminGP,
		})
		log.Println("load user:", email)
	}
	return users
}

func (h *Helper) AddUser(usr interface{}) {
	p := usr.(*user.USER)
	s, err := h.DB.Prepare(res.InsertUser)
	utils.CheckErrors(err)
	_, err = s.Exec(p.Credit.Email, p.Credit.Password, p.Group.ID)
	utils.CheckErrors(err)
	log.Println("added the user to db:", p.Credit.Email)
}

func (h *Helper) Refresh() {
	user.CachedUsers = h.GetUsers().([]*user.USER)
}
