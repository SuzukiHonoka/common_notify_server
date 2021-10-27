package api

import (
	api "common_notify_server/common"
	"common_notify_server/internal/notification"
	"common_notify_server/internal/utils"
	"net/http"
)

const actionFeeds = "feeds"

func GetFeeds(w http.ResponseWriter, r *http.Request) {
	// parse session
	if s := utils.ParseSession(r); s != nil {
		// get bounded notification
		nfs := notification.CachedNotifications.FindNotificationsByUser(s.Bound)
		utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionFeeds, true, nfs)))
		return
	}
	// session not validated
	w.WriteHeader(http.StatusUnauthorized)
	return
}
