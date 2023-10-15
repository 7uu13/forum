package middleware

import (
	"time"
)

var Sessions = make(map[string]Session)

type Session struct {
	Username string
	Expiry   time.Time
}

func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}
