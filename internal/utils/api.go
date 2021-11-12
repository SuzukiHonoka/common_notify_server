package utils

import (
	"common_notify_server/internal/errors"
	"common_notify_server/internal/notification"
	"common_notify_server/internal/session"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

var RealIPHeader = []string{"X-REAL-IP", "X-FORWARDED-FOR", ""}

// ParseAccount parse from post form data
func ParseAccount(w http.ResponseWriter, r *http.Request) (string, string, error) {
	if err := r.ParseForm(); err == nil {
		// get form data
		email := r.FormValue("email")
		pass := r.FormValue("password")
		// check if empty
		if IsNotEmpty(email, pass) {
			// content type
			w.Header().Set("Content-Type", "application/json")
			return email, pass, nil
		}
	}
	WriteStringReplyNoCheck(w, http.StatusBadRequest, errors.ParseAccountFailed.Error())
	return "", "", errors.ParseAccountFailed
}

// ParseSession from http header
func ParseSession(w http.ResponseWriter, r *http.Request) *session.Session {
	if cs, err := r.Cookie("session"); err == nil {
		if s := session.CachedSessionsMap.FindSessionByID(ParseIP(r), cs.Value); s != nil {
			// content type
			w.Header().Set("Content-Type", "application/json")
			return s
		}
	}
	// session not validated
	WriteStringReplyNoCheck(w, http.StatusUnauthorized, errors.SessionInvalid.Error())
	return nil
}

// VtoJson convert struct to json
func VtoJson(v interface{}) []byte {
	r, _ := json.Marshal(v)
	return r
}

// WriteReplyNoCheck write []byte to resp without any check
func WriteReplyNoCheck(w http.ResponseWriter, code int, msg []byte) {
	w.WriteHeader(code)
	_, _ = w.Write(msg)
}

func WriteStringReplyNoCheck(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	_, _ = fmt.Fprint(w, msg)
}

// ParseIP from http header
func ParseIP(r *http.Request) net.IP {
	for i, header := range RealIPHeader {
		var ips string
		switch i {
		case 0:
			fallthrough
		case 1:
			ips = r.Header.Get(header)
			if len(ips) == 0 {
				continue
			}
			if i == 1 {
				ips = strings.Split(ips, ",")[0]
			}
		case 2:
			ips = r.RemoteAddr
		}
		if ip := net.ParseIP(ips); ip != nil {
			return ip
		}
	}
	return nil
}

// CloseBodyNoCheck closes the post body without any check
func CloseBodyNoCheck(body io.ReadCloser) {
	_ = body.Close()
}

// ParseNotificationList Parse post json payload to notification.Message
func ParseNotificationList(r *http.Request) *notification.Message {
	var tmp notification.Message
	defer CloseBodyNoCheck(r.Body)
	bodyType := r.Header.Get("Content-Type")
	if len(bodyType) > 0 && bodyType == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err == nil {
			return &tmp
		}
	}
	return nil
}
