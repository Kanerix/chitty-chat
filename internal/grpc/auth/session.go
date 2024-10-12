package auth

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/exp/maps"
)

type SessionStore map[string]*Session

type InMemorySessionStore struct {
	mutex   sync.RWMutex
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

func (store *InMemorySessionStore) Save(username string, anonymous bool) (*Session, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if _, ok := store.storage[username]; ok {
		return nil, errors.New("username already has a active session")
	}

	session := &Session{
		Username:  username,
		Anonymous: anonymous,
		token:     uuid.New(),
	}
	store.storage[username] = session

	return session, nil
}

func (store *InMemorySessionStore) Get(username string) (*Session, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	session, ok := store.storage[username]
	if !ok {
		return nil, errors.New("no active session found for username")
	}

	return session, nil
}

func (store *InMemorySessionStore) Delete(username string) error {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	_, ok := store.storage[username]
	if !ok {
		return errors.New("no active session found for username")
	}

	store.storage[username] = nil

	return nil

}

func (store *InMemorySessionStore) List(page int) []string {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	keys := maps.Keys(store.storage)

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

func (session *Session) String() string {
	return fmt.Sprintf(
		"%s:%v:%s",
		session.Username,
		session.Anonymous,
		session.token,
	)
}

func StringToSession(s string) (*Session, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 3 {
		return nil, errors.New("parts of session string not equal to 3")
	}

	username := parts[0]
	anonymous := parts[1] == "true"
	fmt.Println(parts[2])
	token, err := uuid.Parse(parts[2])
	if err != nil {
		return nil, errors.New("error parsing token into UUID")
	}

	return &Session{username, anonymous, token}, nil
}
