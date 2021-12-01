package api

import (
	"net/http"
	api "nfly/common"
	"nfly/internal/errors"
	"nfly/internal/notification"
	"nfly/internal/utils"
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
