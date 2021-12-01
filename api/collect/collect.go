package api

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	api "nfly/common"
	"nfly/internal/errors"
	"nfly/internal/notification"
	"nfly/internal/utils"
	"strconv"
)

const actionCollect = "collect"

func Collect(w http.ResponseWriter, r *http.Request) {
	// collect notification push status
	if s := utils.ParseSession(w, r); s != nil {
		var uid uuid.UUID
		// read flag from body
		if b, err := io.ReadAll(r.Body); err == nil {
			parsed := true
			var status bool
			// 2-way
			if b[0] <= 0x01 {
				status = b[0] == 0x01
			} else {
				// convert flag to bool
				if status, err = strconv.ParseBool(string(b)); err != nil {
					parsed = false
				}
			}
			if parsed {
				// get notification uuid
				if uid, err = uuid.Parse(mux.Vars(r)["uuid"]); err == nil {
					// find by uuid
					for _, n := range notification.CachedNotifications[s.Bound.Credit.Email] {
						if n.Header.UUID == uid {
							// set flag
							n.Status.Pushed = status
							utils.WriteReplyNoCheck(w, http.StatusOK, utils.VtoJson(*api.NewReply(actionCollect, true, nil)))
							return
						}
					}
					// not found
					utils.WriteReplyNoCheck(w, http.StatusNotFound, utils.VtoJson(*api.NewReply(actionCollect, false, errors.NotificationNotFound.Error())))
					return
				}
			}
		}
		// if collect failed
		utils.WriteReplyNoCheck(w, http.StatusBadRequest, utils.VtoJson(*api.NewReply(actionCollect, false, errors.BadRequest.Error())))
	}
}
