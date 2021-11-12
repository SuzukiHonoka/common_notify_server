package api

import (
	api "common_notify_server/common"
	"common_notify_server/internal/notification"
	"common_notify_server/internal/session"
	"common_notify_server/internal/user"
	"common_notify_server/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

const actionDelete = "delete"

func UserDelete(w http.ResponseWriter, r *http.Request) {
	if s := utils.ParseSession(w, r); s != nil {
		// check user group
		if s.Bound.Group.ID != user.AdminGP.ID {
			http.Error(w, string(utils.VtoJson(*api.NewReply(actionDelete, false, nil))), http.StatusUnauthorized)
			return
		}
		// get user to be deleted
		u := user.CachedUsersMap.FindUserByEmail(mux.Vars(r)["user"])
		if u == nil {
			http.Error(w, string(utils.VtoJson(*api.NewReply(actionDelete, false, nil))), http.StatusNotFound)
			return
		}
		// also delete related session and notifications
		now := time.Now()
		for _, v := range session.CachedSessionsMap.FindSessionByUser(u) {
			v.ExpDate = now // invalidate the session
		}
		notification.CachedNotifications.DeleteNotificationsByUser(u)
		if user.CachedUsersMap.DeleteUser(u) {
			utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionDelete, true, u)))
			return
		}
		http.Error(w, string(utils.VtoJson(*api.NewReply(actionDelete, false, nil))), http.StatusInternalServerError)
	}
}
