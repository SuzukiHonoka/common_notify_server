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
	if s := utils.ParseSession(r); len(s) > 0 {
		if u := session.FindUserBySession(s); u != nil {
			nfs := notification.FindNotificationsByUser(u)
			utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionFeeds, true, nfs)))
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
