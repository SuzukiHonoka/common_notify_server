package api

import (
	api "common_notify_server/common"
	"common_notify_server/internal/user"
	"common_notify_server/internal/utils"
	"net/http"
)

const actionRegister = "register"

func UserRegister(w http.ResponseWriter, r *http.Request) {
	// parse account
	if email, pass, err := utils.ParseAccount(r); err == nil {
		// check form values
		if utils.IsNotEmpty(email, pass) {
			// user register
			var u *user.User
			u, err = user.Register(email, pass, nil)
			// if register success
			if err == nil {
				utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionRegister, true, u)))
				return
			}
			// if register failed
			utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionRegister, false, err.Error())))
			return
		}
	}
	// if parse failed
	w.WriteHeader(http.StatusBadRequest)
	utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionLogin, false, "check parameters")))
}
