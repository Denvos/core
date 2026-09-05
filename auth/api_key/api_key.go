package apikey

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sync"
	"time"
)

type APIKey struct {
	mu          sync.RWMutex
	keys        map[string]keyInfo
	hashed      bool
}

type keyInfo struct {
	ID        string
	Scopes    []string
	CreatedAt time.Time
	LastUsed  time.Time
	Metadata  map[string]interface{}
}

type Option func(*APIKey)

func WithHashing() Option {
	return func(a *APIKey) {
		a.hashed = true
	}
}

func New(opts ...Option) *APIKey {
	a := &APIKey{
		keys:   make(map[string]keyInfo),
		hashed: false,
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func (a *APIKey) Generate(scopes []string, metadata map[string]interface{}) (string, string, error) {
	id := randomString(8)
	key := randomString(32)
	raw := id + ":" + key
	if a.hashed {
		hash := sha256.Sum256([]byte(raw))
		key = base64.StdEncoding.EncodeToString(hash[:])
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.keys[key] = keyInfo{
		ID:        id,
		Scopes:    scopes,
		CreatedAt: time.Now(),
		Metadata:  metadata,
	}
	return id, raw, nil
}

func (a *APIKey) Validate(raw string) (string, []string, error) {
	if a.hashed {
		hash := sha256.Sum256([]byte(raw))
		hashed := base64.StdEncoding.EncodeToString(hash[:])
		return a.validateHashed(hashed)
	}
	return a.validatePlain(raw)
}

func (a *APIKey) validatePlain(raw string) (string, []string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	info, ok := a.keys[raw]
	if !ok {
		return "", nil, fmt.Errorf("invalid API key")
	}
	return info.ID, info.Scopes, nil
}

func (a *APIKey) validateHashed(hashed string) (string, []string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	info, ok := a.keys[hashed]
	if !ok {
		return "", nil, fmt.Errorf("invalid API key")
	}
	return info.ID, info.Scopes, nil
}

func (a *APIKey) Revoke(raw string) error {
	if a.hashed {
		hash := sha256.Sum256([]byte(raw))
		hashed := base64.StdEncoding.EncodeToString(hash[:])
		return a.revokeHashed(hashed)
	}
	return a.revokePlain(raw)
}

func (a *APIKey) revokePlain(raw string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.keys[raw]; !ok {
		return fmt.Errorf("key not found")
	}
	delete(a.keys, raw)
	return nil
}

func (a *APIKey) revokeHashed(hashed string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.keys[hashed]; !ok {
		return fmt.Errorf("key not found")
	}
	delete(a.keys, hashed)
	return nil
}

func randomString(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)[:n]
}
