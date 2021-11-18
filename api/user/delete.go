package api

import (
	api "common_notify_server/common"
	"common_notify_server/internal/notification"
	"common_notify_server/internal/session"
	"common_notify_server/internal/user"
	"common_notify_server/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
)

const actionDelete = "delete"

func UserDelete(w http.ResponseWriter, r *http.Request) {
	if s := utils.ParseSession(w, r); s != nil {
		// check user group
		if s.Bound.Group.ID != user.AdminGP.ID {
			utils.WriteReplyNoCheck(w, http.StatusUnauthorized, utils.VtoJson(*api.NewReply(actionDelete, false, nil)))
			return
		}
		// get user to be deleted
		u := user.CachedUsersMap.FindUserByEmail(mux.Vars(r)["user"])
		if u == nil {
			utils.WriteReplyNoCheck(w, http.StatusNotFound, utils.VtoJson(*api.NewReply(actionDelete, false, nil)))
			return
		}
		// also delete related session and notifications
		session.CachedSessionsMap.DeleteSessionByUser(s.Bound)
		notification.CachedNotifications.DeleteNotificationsByUser(u)
		if user.CachedUsersMap.DeleteUser(u) {
			utils.WriteReplyNoCheck(w, http.StatusOK, utils.VtoJson(*api.NewReply(actionDelete, true, u)))
			return
		}
		utils.WriteReplyNoCheck(w, http.StatusInternalServerError, utils.VtoJson(*api.NewReply(actionDelete, false, nil)))
	}
}
