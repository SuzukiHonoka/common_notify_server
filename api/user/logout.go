package api

import (
	"log"
	"net/http"
	api "nfly/common"
	"nfly/internal/session"
	"nfly/internal/utils"
)

const actionLogout = "logout"

func UserLogout(w http.ResponseWriter, r *http.Request) {
	if s := utils.ParseSession(w, r); s != nil {
		log.Printf("user: %s => logout\n", s.Bound.Credit.Email)
		session.CachedSessionsMap.DeleteSessionBySession(s)
		//session.CachedSessionsMap.DeleteSessionByUser(s.Bound)
		utils.WriteReplyNoCheck(w, http.StatusOK, utils.VtoJson(*api.NewReply(actionLogout, true, nil)))
	}
}
