package sessions

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        string
	Identity  string
	Data      map[string]interface{}
	CreatedAt time.Time
	ExpiresAt time.Time
}

type Store struct {
	mu       sync.RWMutex
	sessions map[string]Session
}

func NewStore() *Store {
	return &Store{sessions: make(map[string]Session)}
}

func (s *Store) Create(identity string, ttl time.Duration) *Session {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := uuid.New().String()
	sess := Session{
		ID:        id,
		Identity:  identity,
		Data:      make(map[string]interface{}),
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(ttl),
	}
	s.sessions[id] = sess
	return &sess
}

func (s *Store) Get(id string) (*Session, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sess, ok := s.sessions[id]
	if !ok {
		return nil, false
	}
	if time.Now().After(sess.ExpiresAt) {
		return nil, false
	}
	return &sess, true
}

func (s *Store) Delete(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, id)
}

func (s *Store) Cleanup() {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	for id, sess := range s.sessions {
		if now.After(sess.ExpiresAt) {
			delete(s.sessions, id)
		}
	}
}
