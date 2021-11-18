package api

import (
	api "common_notify_server/common"
	"common_notify_server/internal/session"
	"common_notify_server/internal/utils"
	"log"
	"net/http"
)

const actionLogout = "logout"

func UserLogout(w http.ResponseWriter, r *http.Request) {
	if s := utils.ParseSession(w, r); s != nil {
		log.Printf("user: %s => logout\n", s.Bound.Credit.Email)
		session.CachedSessionsMap.DeleteSessionByUser(s.Bound)
		utils.WriteReplyNoCheck(w, http.StatusOK, utils.VtoJson(*api.NewReply(actionLogout, true, nil)))
	}
}
