package session

import (
	"common_notify_server/internal/user"
	"github.com/google/uuid"
	"log"
	"net"
	"time"
)

const maxSession = 1

type Session struct {
	RemoteAddr net.IP
	Bound      *user.USER
	UUID       uuid.UUID
	ExpDate    time.Time
}

func NewSession(ip net.IP, user *user.USER) *Session {
	var count int
	for _, session := range CachedSession {
		if ip.Equal(session.RemoteAddr) {
			count++
		}
	}
	if count >= maxSession {
		return nil
	}
	uid, _ := uuid.NewRandom()
	t := &Session{
		RemoteAddr: ip,
		Bound:      user,
		UUID:       uid,
		ExpDate:    time.Now().AddDate(0, 0, 7),
	}
	CachedSession = append(CachedSession, t)
	log.Println("total:", len(CachedSession), "=>", "new session:", uid, "for", user.Credit.Email)
	return t
}

func CheckSession(user *user.USER, uid string) bool {
	for _, v := range CachedSession {
		if v.Bound == user {
			return v.UUID.String() == uid
		}
	}
	return false
}
