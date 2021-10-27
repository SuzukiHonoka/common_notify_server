package session

import (
	"common_notify_server/internal/user"
	"common_notify_server/internal/utils"
	"github.com/google/uuid"
	"log"
	"net"
	"net/http"
	"time"
)

const maxSession = 8

type Session struct {
	RemoteAddr net.IP
	Bound      *user.User // only bond single user per session
	UUID       uuid.UUID
	ExpDate    time.Time
}

func NewSession(ip net.IP, user *user.User) *Session {
	// check the ip allocated session
	var count int
	for _, session := range CachedSessions {
		if ip.Equal(session.RemoteAddr) {
			count++
		}
	}
	// reject to alloc if count greater than max limitation
	if count >= maxSession {
		return nil
	}
	// alloc
	uid, _ := uuid.NewRandom()
	t := &Session{
		RemoteAddr: ip,
		Bound:      user,
		UUID:       uid,
		ExpDate:    time.Now().AddDate(0, 0, 7), // exp in next 7 days
	}
	// add to cache
	CachedSessions = append(CachedSessions, t)
	log.Println("total:", len(CachedSessions), "=>", "new session:", uid, "for", user.Credit.Email)
	return t
}

func FindUserBySessionID(r *http.Request, uid string) *user.User {
	for _, session := range CachedSessions {
		id, err := uuid.Parse(uid)
		// if parse uid failed
		if err != nil {
			return nil
		}
		// double check
		if session.UUID == id && session.RemoteAddr.Equal(utils.ParseIP(r)) {
			return session.Bound
		}
	}
	return nil
}
