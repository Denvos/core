package tokens

import (
	"sync"
	"time"
)

type Token struct {
	ID        string
	Value     string
	Type      string // refresh, access, etc.
	Identity  string
	Scopes    []string
	ExpiresAt time.Time
	Revoked   bool
}

type Store struct {
	mu     sync.RWMutex
	tokens map[string]Token
}

func NewStore() *Store {
	return &Store{tokens: make(map[string]Token)}
}

func (s *Store) Add(token Token) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tokens[token.ID] = token
}

func (s *Store) Get(id string) (Token, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.tokens[id]
	return t, ok
}

func (s *Store) Revoke(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if t, ok := s.tokens[id]; ok {
		t.Revoked = true
		s.tokens[id] = t
	}
}

func (s *Store) Cleanup() {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	for id, t := range s.tokens {
		if t.ExpiresAt.Before(now) || t.Revoked {
			delete(s.tokens, id)
		}
	}
}
