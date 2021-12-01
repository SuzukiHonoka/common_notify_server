package api

import (
	"net/http"
	api "nfly/common"
	"nfly/internal/notification"
	"nfly/internal/utils"
)

const actionFeeds = "feeds"

func GetFeeds(w http.ResponseWriter, r *http.Request) {
	// parse session
	if s := utils.ParseSession(w, r); s != nil {
		// get bounded notification
		nfs := notification.CachedNotifications.FindNotificationsByUser(s.Bound)
		utils.WriteReplyNoCheck(w, http.StatusOK, utils.VtoJson(*api.NewReply(actionFeeds, true, nfs)))
	}
}
