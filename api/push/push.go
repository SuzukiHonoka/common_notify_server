package api

import (
	api "common_notify_server/common"
	"common_notify_server/internal/errors"
	"common_notify_server/internal/notification"
	"common_notify_server/internal/utils"
	"net/http"
)

const actionPush = "push"

func Push(w http.ResponseWriter, r *http.Request) {
	if s := utils.ParseSession(w, r); s != nil {
		if ns := utils.ParseNotificationList(r); ns != nil {
			// add ns to cache
			notification.CachedNotifications[s.Bound.Credit.Email] = append(notification.CachedNotifications[s.Bound.Credit.Email],
				notification.NewNotification(s.Bound, ns.Title, ns.MessageChain))
			w.WriteHeader(http.StatusOK)
			utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionPush, true, ns)))
			return
		}
		http.Error(w, errors.NotificationsListParseFailed.Error(), http.StatusBadRequest)
	}
}
