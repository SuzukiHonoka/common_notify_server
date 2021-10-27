package api

import (
	api "common_notify_server/common"
	"common_notify_server/internal/notification"
	"common_notify_server/internal/utils"
	"net/http"
)

const actionPush = "push"

func Push(w http.ResponseWriter, r *http.Request) {
	if s := utils.ParseSession(w, r); s != nil {
		// get bounded notification
		nfs := notification.CachedNotifications.FindNotificationsByUser(s.Bound)
		utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionPush, true, nfs)))
		return
	}
}
