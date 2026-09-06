package auth

import (
    "crypto/sha256"
    "encoding/hex"
)

type BasicAuthStore struct {
    users map[string]string
}

func NewBasicAuthStore() *BasicAuthStore {
    return &BasicAuthStore{
        users: make(map[string]string),
    }
}

func (s *BasicAuthStore) Set(username, password string) {
    hash := sha256.Sum256([]byte(password))
    s.users[username] = hex.EncodeToString(hash[:])
}

func (s *BasicAuthStore) Verify(username, password string) bool {
    hash, ok := s.users[username]
    if !ok {
        return false
    }
    check := sha256.Sum256([]byte(password))
    return hash == hex.EncodeToString(check[:])
}

func (s *BasicAuthStore) Delete(username string) {
    delete(s.users, username)
}

func (s *BasicAuthStore) Has(username string) bool {
    _, ok := s.users[username]
    return ok
}

func (s *BasicAuthStore) Count() int {
    return len(s.users)
}
