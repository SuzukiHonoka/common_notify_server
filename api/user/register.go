package api

import (
	api "common_notify_server/common"
	"common_notify_server/internal/user"
	"common_notify_server/internal/utils"
	"net/http"
)

const actionRegister = "register"

func UserRegister(w http.ResponseWriter, r *http.Request) {
	if r.ParseForm() != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	email := r.FormValue("email")
	pass := r.FormValue("password")
	if utils.IsEmpty(email, pass) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := user.Register(email, pass, nil)
	if err != nil {
		utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionRegister, false, err.Error())))
		return
	}
	utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionRegister, u != nil, u)))
}
