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
	for _, session := range CachedSessions {
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
	CachedSessions = append(CachedSessions, t)
	log.Println("total:", len(CachedSessions), "=>", "new session:", uid, "for", user.Credit.Email)
	return t
}

func FindUserBySession(uid string) *user.USER {
	for _, session := range CachedSessions {
		if session.UUID.String() == uid {
			return session.Bound
		}
	}
	return nil
}

func CheckSession(user *user.USER, uid string) bool {
	for _, v := range CachedSessions {
		if v.Bound == user {
			return v.UUID.String() == uid
		}
	}
	return false
}
