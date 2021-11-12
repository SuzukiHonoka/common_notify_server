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
			utils.WriteReplyNoCheck(w, http.StatusOK, utils.VtoJson(*api.NewReply(actionPush, true, ns)))
			return
		}
		utils.WriteStringReplyNoCheck(w, http.StatusBadRequest, errors.NotificationsListParseFailed.Error())
	}
}
