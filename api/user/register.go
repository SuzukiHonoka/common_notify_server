package api

import (
	api "common_notify_server/common"
	"common_notify_server/internal/user"
	"common_notify_server/internal/utils"
	"net/http"
)

const actionRegister = "register"

func UserRegister(w http.ResponseWriter, r *http.Request) {
	if email, pass, err := utils.ParseAccount(r); err == nil {
		if utils.IsNotEmpty(email, pass) {
			var u *user.USER
			u, err = user.Register(email, pass, nil)
			if err == nil {
				utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionRegister, true, u)))
				return
			}
			utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionRegister, false, err.Error())))
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionLogin, false, "check parameters")))
}
