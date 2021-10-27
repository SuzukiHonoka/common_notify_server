package session

import (
	"common_notify_server/internal/user"
	"github.com/google/uuid"
	"log"
	"net"
	"time"
)

const maxSession = 8

type Session struct {
	RemoteAddr net.IP
	Bound      *user.User // only bond single user per session
	UUID       uuid.UUID
	ExpDate    time.Time
}

type SessionsList []*Session

func NewSession(ip net.IP, user *user.User) *Session {
	// check the ip allocated session
	count := len(CachedSessions.FindSessionByIP(ip))
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

func (x *SessionsList) FindSessionByIP(ip net.IP) SessionsList {
	var tmp SessionsList
	for _, session := range *x {
		if session.RemoteAddr.Equal(ip) {
			tmp = append(tmp, session)
		}
	}
	return tmp
}

func (x *SessionsList) FindSessionByID(ip net.IP, uid string) *Session {
	for _, session := range *x {
		id, err := uuid.Parse(uid)
		// if parse uid failed
		if err != nil {
			return nil
		}
		// double check
		if session.UUID == id && session.RemoteAddr.Equal(ip) {
			return session
		}
	}
	return nil
}
