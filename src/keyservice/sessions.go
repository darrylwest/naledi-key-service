package keyservice

import (
    "code.google.com/p/go-uuid/uuid"
    "sync"
    "time"
)

type Session struct {
    ssid  string
    expires int64

}

func (ss *Session) GetSSID() string {
    return ss.ssid
}

func (ss *Session) GetExpires() int64 {
    return ss.expires
}

type SessionMap struct {
    hash map[string]Session
    sync.RWMutex
}

var (
    sessions SessionMap
    dfltTimeout int64
)

func init() {
    sessions = SessionMap{ hash: make(map[string]Session )}
    dfltTimeout = 60 * 5 // five minutes
}

func (ss *SessionMap) Len() int {
    return len(ss.hash)
}

func GetSessions() *SessionMap {
    return &sessions
}

func CreateSession(expires int64) *Session {
    session := new(Session)

    session.ssid = uuid.New()

    if expires <= 0 {
        session.expires = time.Now().Unix() + dfltTimeout
    } else {
        session.expires = expires
    }

    // session meta data

    sessions.Lock()
    sessions.hash[ session.ssid ] = *session
    sessions.Unlock()

    return session
}

func ValidateSession(ssid string) bool {
	var valid = false
	sessions.Lock()

	if session, ok := sessions.hash[ssid]; ok {
		if session.expires >= time.Now().Unix() {
			valid = true
		} else {
			delete(sessions.hash, ssid)
		}

		log.Info("session exists: %s and valid: %v, t: %d", session.ssid, valid, session.expires)
	} else {
		log.Warn("attempted session validation with unknown key: %s", session)
	}

	sessions.Unlock()

	return valid
}

func FindSession(ssid string) (*Session, bool) {
	sessions.RLock()
	ss, ok := sessions.hash[ssid]
	sessions.RUnlock()

	if ok {
		return &ss, ok
	} else {
		return nil, false
	}
}

func RemoveSession(ssid string) {
	sessions.Lock()
	delete(sessions.hash, ssid)
	sessions.Unlock()
}

func PurgeSessions() {
	log.Info("purge sessions: %d", len(sessions.hash))
	now := time.Now().Unix()

	for k, session := range sessions.hash {
		log.Info("key: %s = time: %d", k, session.expires)

		if session.expires < now {
			log.Info("purge key: %s = time: %d", k, session.expires)
			sessions.Lock()
			delete(sessions.hash, k)
			sessions.Unlock()
		}
	}
}
