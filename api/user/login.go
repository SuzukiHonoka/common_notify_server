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
	// parse account
	if email, pass, err := utils.ParseAccount(w, r); err == nil {
		// user login
		var u *user.User
		u, err = user.Login(email, pass)
		// if login success
		if err == nil {
			// alloc session
			if s := session.NewSession(utils.ParseIP(r), u); s != nil {
				// set header session value
				w.Header().Set("Session", s.UUID.String())
				utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionLogin, true, u)))
				return
			}
			// if alloc failed
			utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionLogin, false, errors.SessionPoolMaxReached.Error())))
			return
		}
		// if login failed
		w.WriteHeader(http.StatusUnauthorized)
		utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionLogin, false, err.Error())))
		return
	}
}
