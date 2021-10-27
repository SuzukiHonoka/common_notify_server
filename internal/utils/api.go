package utils

import (
	"common_notify_server/api/session"
	"common_notify_server/internal/errors"
	"encoding/json"
	"net"
	"net/http"
	"strings"
)

var RealIPHeader = []string{"X-REAL-IP", "X-FORWARDED-FOR", ""}

// ParseAccount parse from post form data
func ParseAccount(r *http.Request) (string, string, error) {
	// parse form
	if err := r.ParseForm(); err != nil {
		return "", "", err
	}
	// get form data
	email := r.FormValue("email")
	pass := r.FormValue("password")
	// check if empty
	if IsNotEmpty(email, pass) {
		return email, pass, nil
	}
	return "", "", errors.ParseAccountFailed
}

// VtoJson convert struct to json
func VtoJson(v interface{}) []byte {
	r, _ := json.Marshal(v)
	return r
}

// WriteReplyNoCheck write []byte to resp without any check
func WriteReplyNoCheck(w http.ResponseWriter, msg []byte) {
	_, _ = w.Write(msg)
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

// ParseSession from http header
func ParseSession(r *http.Request) *session.Session {
	return session.CachedSessions.FindSessionByID(ParseIP(r), r.Header.Get("Session"))
}
