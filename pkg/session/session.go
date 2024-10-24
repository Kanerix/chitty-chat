package session

import (
	"errors"
	"fmt"
	"path"
	"sync"

	"github.com/google/uuid"
	"github.com/kanerix/chitty-chat/pkg/config"
	"golang.org/x/exp/maps"
)

type SessionKey struct{}

type SessionStore map[string]*Session

var SessionFile = path.Join(config.ChippyPath, "session.txt")

type InMemorySessionStore struct {
	sync.RWMutex
	storage SessionStore
}

type Session struct {
	Username  string
	Anonymous bool
	token     uuid.UUID
}

func NewInMemorySessionStore() *InMemorySessionStore {
	return &InMemorySessionStore{
		storage: make(SessionStore),
	}
}

func (s *InMemorySessionStore) Save(username string, anonymous bool) (*Session, error) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.storage[username]; ok {
		return nil, errors.New("username already has a active session")
	}

	session := &Session{
		Username:  username,
		Anonymous: anonymous,
		token:     uuid.New(),
	}
	s.storage[username] = session

	return session, nil
}

func (s *InMemorySessionStore) Get(username string) (*Session, error) {
	s.RLock()
	defer s.RUnlock()

	session, ok := s.storage[username]
	if !ok {
		return nil, errors.New("no active session found for username")
	}

	return session, nil
}

func (s *InMemorySessionStore) Delete(username string) error {
	s.RLock()
	defer s.RUnlock()

	_, ok := s.storage[username]
	if !ok {
		return errors.New("no active session found for username")
	}

	s.storage[username] = nil

	return nil

}

func (s *InMemorySessionStore) List(page int) []string {
	s.RLock()
	defer s.RUnlock()

	keys := maps.Keys(s.storage)

	start := (page - 1) * 10
	end := start + 10

	if start > len(keys) {
		return []string{}
	}

	if end > len(keys) {
		end = len(keys)
	}

	return keys[start:end]
}

func (s *Session) String() string {
	return fmt.Sprintf("%s:%v:%s", s.Username, s.Anonymous, s.token)
}

var (
	ErrSessionKeyNotFound = errors.New("session key not found")
)
