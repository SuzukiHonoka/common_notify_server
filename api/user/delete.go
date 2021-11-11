package api

import (
	api "common_notify_server/common"
	"common_notify_server/internal/session"
	"common_notify_server/internal/user"
	"common_notify_server/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

const ActionDelete = "delete"

func UserDelete(w http.ResponseWriter, r *http.Request) {
	if s := utils.ParseSession(w, r); s != nil {
		// check user group
		if s.Bound.Group.ID != user.AdminGP.ID {
			http.Error(w, string(utils.VtoJson(*api.NewReply(ActionDelete, false, nil))), http.StatusUnauthorized)
			return
		}
		// get user to be deleted
		u := user.CachedUsers.FindUserByEmail(mux.Vars(r)["user"])
		if u == nil {
			http.Error(w, string(utils.VtoJson(*api.NewReply(ActionDelete, false, nil))), http.StatusNotFound)
			return
		}
		// todo: also delete related session and notifications
		if user.CachedUsers.DeleteUser(u) {
			session.CachedSessions.FindSessionByUser(u).ExpDate = time.Now() // invalidate the session
			utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(ActionDelete, true, u)))
			return
		}
		http.Error(w, string(utils.VtoJson(*api.NewReply(ActionDelete, false, nil))), http.StatusInternalServerError)
	}
}
