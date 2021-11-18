package api

import (
	api "common_notify_server/common"
	"common_notify_server/internal/session"
	"common_notify_server/internal/utils"
	"log"
	"net/http"
	"time"
)

const actionLogout = "logout"

func UserLogout(w http.ResponseWriter, r *http.Request) {
	if s := utils.ParseSession(w, r); s != nil {
		log.Printf("user: %s => logout\n", s.Bound.Credit.Email)
		now := time.Now()
		for _, v := range session.CachedSessionsMap.FindSessionByUser(s.Bound) {
			v.ExpDate = now // invalidate the session
		}
		utils.WriteReplyNoCheck(w, http.StatusOK, utils.VtoJson(*api.NewReply(actionLogout, true, nil)))
	}
}
