package api

import (
	"net/http"
	api "nfly/common"
	"nfly/internal/errors"
	"nfly/internal/session"
	"nfly/internal/user"
	"nfly/internal/utils"
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
				http.SetCookie(w, &http.Cookie{
					Name:    "session",
					Value:   s.UUID.String(),
					Expires: s.ExpDate,
				})
				utils.WriteReplyNoCheck(w, http.StatusOK, utils.VtoJson(*api.NewReply(actionLogin, true, u)))
				return
			}
			// if alloc failed
			utils.WriteReplyNoCheck(w, http.StatusLocked, utils.VtoJson(*api.NewReply(actionLogin, false, errors.SessionPoolMaxReached.Error())))
			return
		}
		// if login failed
		utils.WriteReplyNoCheck(w, http.StatusUnauthorized, utils.VtoJson(*api.NewReply(actionLogin, false, err.Error())))
	}
}
