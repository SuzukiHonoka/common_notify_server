package api

import (
	api "common_notify_server/common"
	"common_notify_server/internal/errors"
	"common_notify_server/internal/notification"
	"common_notify_server/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

const actionCollect = "collect"

func Collect(w http.ResponseWriter, r *http.Request) {
	// collect notification push status
	if s := utils.ParseSession(w, r); s != nil {
		var uid uuid.UUID
		status := false
		// read flag from body
		if b, err := ioutil.ReadAll(r.Body); err == nil {
			// convert flag to bool
			if status, err = strconv.ParseBool(string(b)); err == nil {
				// get notification uuid
				if uid, err = uuid.Parse(mux.Vars(r)["uuid"]); err == nil {
					// find by uuid
					for _, n := range notification.CachedNotifications[s.Bound] {
						if n.Header.UUID == uid {
							// set flag
							n.Status.Pushed = status
							utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionCollect, true, nil)))
							return
						}
					}
				}
			}
		}
		// if collect failed
		w.WriteHeader(http.StatusBadRequest)
		utils.WriteReplyNoCheck(w, utils.VtoJson(*api.NewReply(actionCollect, false, errors.BadRequest.Error())))
	}
}
