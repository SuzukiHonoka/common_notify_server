package utils

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
)

var RealIPHeader = []string{"X-REAL-IP", "X-FORWARDED-FOR", ""}

func ParseAccount(r *http.Request) (string, string, error) {
	if err := r.ParseForm(); err != nil {
		return "", "", err
	}
	return r.FormValue("email"), r.FormValue("password"), nil
}

func VtoJson(v interface{}) []byte {
	r, _ := json.Marshal(v)
	return r
}

func WriteReplyNoCheck(w http.ResponseWriter, msg []byte) {
	_, _ = w.Write(msg)
}

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

func ParseSession(r *http.Request) string {
	return r.Header.Get("Session")
}
