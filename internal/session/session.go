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

type SessionsMap map[string]*Session // uuid as key

func NewSession(ip net.IP, user *user.User) *Session {
	// check the ip allocated session
	count := len(CachedSessionsMap.FindSessionByUser(user))
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
	//CachedSessions = append(CachedSessions, t)
	CachedSessionsMap[t.UUID.String()] = t
	log.Printf("new session: %s => %s\n", uid, user.Credit.Email)
	return t
}

func (x *SessionsMap) FindSessionByUser(user *user.User) SessionsList {
	var m SessionsList
	for _, v := range *x {
		if v.Bound == user {
			m = append(m, v)
		}
	}
	return m
}

func (x *SessionsMap) FindSessionByID(ip net.IP, uid string) *Session {
	var session *Session
	if v, ok := (*x)[uid]; ok {
		session = v
	}
	if session != nil {
		// double check
		if session.RemoteAddr.Equal(ip) && !x.CleanIfExpired(session) {
			log.Printf("user: %s => using session: %s\n", session.Bound.Credit.Email, session.UUID.String())
			return session
		}
	}
	return nil
}

// CleanIfExpired if expired, clean it and return true
func (x *SessionsMap) CleanIfExpired(session *Session) bool {
	// check if expired
	if session.ExpDate.Sub(time.Now()).Milliseconds() < 0 {
		// find and clean
		delete(*x, session.UUID.String())
		return true
	}
	return false
}
