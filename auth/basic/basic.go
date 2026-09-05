package basic

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type Basic struct {
	mu    sync.RWMutex
	users map[string]string // username -> hashed password
}

func New() *Basic {
	return &Basic{
		users: make(map[string]string),
	}
}

func (b *Basic) Set(username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	b.users[username] = string(hash)
	return nil
}

func (b *Basic) Validate(username, password string) (bool, error) {
	b.mu.RLock()
	hash, ok := b.users[username]
	b.mu.RUnlock()
	if !ok {
		return false, nil
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (b *Basic) Delete(username string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	delete(b.users, username)
}
