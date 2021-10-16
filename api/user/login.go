package api

import (
	"common_notify_server/api/session"
	api "common_notify_server/common"
	"common_notify_server/internal/errors"
	"common_notify_server/internal/user"
	"common_notify_server/internal/utils"
	"net/http"
)

const actionLogin = "login"

func UserLogin(w http.ResponseWriter, r *http.Request) {
	if email, pass, err := utils.ParseAccount(r); err == nil {
		if utils.IsNotEmpty(email, pass) {
			var u *user.USER
			u, err = user.Login(email, pass)
			if err == nil {
				if s := session.NewSession(utils.ParseIP(r), u); s != nil {
					w.Header().Set("Session", s.UUID.String())
					utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionLogin, true, u)))
				}
				utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionLogin, false, errors.SessionPoolMaxReached.Error())))
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionLogin, false, err.Error())))
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionLogin, false, "check parameters")))
}
