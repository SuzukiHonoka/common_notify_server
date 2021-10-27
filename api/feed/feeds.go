package api

import (
	"common_notify_server/api/session"
	api "common_notify_server/common"
	"common_notify_server/internal/notification"
	"common_notify_server/internal/utils"
	"net/http"
)

const actionFeeds = "feeds"

func GetFeeds(w http.ResponseWriter, r *http.Request) {
	// parse session
	if s := utils.ParseSession(r); len(s) > 0 {
		// get bounded user from session
		if u := session.FindUserBySessionID(r, s); u != nil {
			// get bounded notification
			nfs := notification.CachedNotifications.FindNotificationsByUser(u)
			utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionFeeds, true, nfs)))
			return
		}
		// session not validated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionFeeds, false, "check parameters")))
}
