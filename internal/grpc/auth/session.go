package grpc

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type SessionStorage map[string]*Session

type InMemorySessionStore struct {
	mutex   sync.RWMutex
	storage SessionStorage
}

type Session struct {
	username  string
	anonymous bool
	token     uuid.UUID
	ttl       time.Duration
}

func NewInMemorySessionStore() *InMemorySessionStore {
	return &InMemorySessionStore{}
}

func (store *InMemorySessionStore) Save(username string, anonymous bool) (*Session, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if _, ok := store.storage[username]; ok {
		return nil, errors.New("username already has a active session")
	}

	session := &Session{
		username:  username,
		anonymous: anonymous,
		token:     uuid.New(),
		ttl:       time.Duration(5 * float64(time.Minute)),
	}
	store.storage[username] = session
	return session, nil
}

func (store *InMemorySessionStore) Get(username string) (*Session, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	session_token, ok := store.storage[username]
	if !ok {
		return nil, errors.New("no active session found for username")
	}

	return session_token, nil
}

func (session *Session) String() string {
	return fmt.Sprintf("%s:%v:%s", session.username, session.anonymous, session.token)
}
