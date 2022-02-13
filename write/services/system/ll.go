package system

import (
	"sync"
	"time"
)

var LL = &LoginLimiter{
	Date:  time.Now().Format("2006-01-02"),
	Count: 0,
	mutex: &sync.RWMutex{},
}

type LoginLimiter struct {
	Date  string `json:"date"`
	Count uint   `json:"count"`
	mutex *sync.RWMutex
}

// ?
func (ll *LoginLimiter) State() bool {
	ll.mutex.RLock()
	date := time.Now().Format("2006-01-02")
	if ll.Date != date {
		ll.mutex.RUnlock()
		ll.mutex.Lock()
		ll.Date = date
		ll.Count = 0
		ll.mutex.Unlock()
		return true
	} else {
		state := (ll.Count >= 3)
		ll.mutex.RUnlock()
		return !state
	}
}

func (ll *LoginLimiter) Inc() {
	ll.mutex.Lock()
	defer ll.mutex.Unlock()
	ll.Count++
}
