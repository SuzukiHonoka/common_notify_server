package sqlite

import (
	config "common_notify_server/config/database"
	database "common_notify_server/internal/database/common"
	"common_notify_server/internal/notification"
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
	var users []*user.User
	for row.Next() {
		var email string
		var password string
		var group user.Type
		err = row.Scan(&email, &password, &group)
		utils.CheckErrors(err)
		users = append(users, &user.User{
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
	p := usr.(*user.User)
	s, err := h.DB.Prepare(res.InsertUser)
	utils.CheckErrors(err)
	defer s.Close()
	_, err = s.Exec(p.Credit.Email, p.Credit.Password, p.Group.ID)
	utils.CheckErrors(err)
	log.Println("user added:", p.Credit.Email)
}

func (h *Helper) DelUser(usr interface{}) {
	p := usr.(*user.User)
	s, err := h.DB.Prepare(res.DeleteUser)
	utils.CheckErrors(err)
	defer s.Close()
	_, err = s.Exec(p.Credit.Email)
	utils.CheckErrors(err)
	log.Println("user deleted:", p.Credit.Email)
}

func (h *Helper) Refresh() {
	user.CachedUsers = h.GetUsers().([]*user.User)
	title := "test"
	notification.CachedNotifications[user.CachedUsers[0]] = notification.Notifications{
		notification.NewNotification(user.CachedUsers[0], &title, notification.MessageChain{
			notification.NewTextMessage(title),
			notification.NewBinaryMessage([]byte{1, 2, 1, 2, 3, 4, 5, 6}),
		}),
	}
}
