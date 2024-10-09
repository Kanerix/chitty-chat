package main

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

type InMemorySessionStore struct {
	mutex    sync.RWMutex
	sessions map[string]*Session
}

type Session struct {
	token uuid.UUID
	ttl   time.Duration
}

func NewInMemorySessionStore() *InMemorySessionStore {
	return &InMemorySessionStore{}
}

func (store *InMemorySessionStore) Save(username string) (*Session, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if _, ok := store.sessions[username]; ok {
		return nil, errors.New("username already has a active session")
	}

	session := &Session{
		token: uuid.New(),
		ttl:   time.Duration(5 * float64(time.Minute)),
	}
	store.sessions[username] = session
	return session, nil
}

func (store *InMemorySessionStore) Get(username string) (*Session, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	session_token, ok := store.sessions[username]
	if !ok {
		return nil, errors.New("no active session found for username")
	}

	return session_token, nil
}
