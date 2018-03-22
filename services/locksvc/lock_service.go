package locksvc

import (
	"fmt"
	"sync"
	"time"
)

// Locker is safe to use concurrently.
type Locker struct {
	v   map[string]bool
	mux sync.Mutex
}

var locker *Locker

const (
	// MatchType match
	MatchType = iota
	// MatchSetType match set
	MatchSetType = iota // c0 == 0
	// MatchSetOffer offer
	MatchSetOffer = iota // c1 == 1
)

// LockerInstance get lock singleton
func LockerInstance() *Locker {
	if locker == nil {
		locker = &Locker{v: make(map[string]bool)}
	}
	return locker
}

// GetLockKey get match lock key for update
func GetLockKey(locktype int, ids ...uint) string {
	key := ""
	id := ""
	for _, v := range ids {
		id = fmt.Sprint(id, "_", v)
	}
	switch locktype {
	case MatchType:
		key = fmt.Sprint("MATCH", id)
	case MatchSetType:
		key = fmt.Sprint("MATCH_SET", id)
	case MatchSetOffer:
		key = fmt.Sprint("MATCH_SET_OFFER", id)
	}
	return key
}

// LockWithTicker lock before update database
func (lock *Locker) LockWithTicker(key string, seconds float32) bool {
	tickChan := time.NewTicker(time.Millisecond * 10).C
	timeoutChan := make(chan bool)
	go func() {
		time.Sleep(time.Second * time.Duration(seconds))
		timeoutChan <- true
	}()
	for {
		select {
		case <-tickChan:
			if lock.Lock(key) {
				return true
			}
		case <-timeoutChan:
			return false
		}
	}
}

// Lock within concurrency
func (lock *Locker) Lock(key string) bool {
	lock.mux.Lock()
	defer lock.mux.Unlock()
	if lock.v[key] {
		return false
	}
	lock.v[key] = true
	return true
}

// Unlock within concurrency
func (lock *Locker) Unlock(key string) {
	lock.mux.Lock()
	defer lock.mux.Unlock()
	delete(lock.v, key)
}

func UsingLockJob(key, worker string) {
	for {
		fmt.Println(worker, "apply authorization")
		l := LockerInstance().LockWithTicker(key, 2)
		if !l {
			fmt.Println(worker, "faild")
			continue
		}
		fmt.Println(worker, "doing job...")
		time.Sleep(time.Second * 3)
		LockerInstance().Unlock(key)
	}
}
