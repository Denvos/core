package cookies

import (
    "crypto/rand"
    "encoding/base64"
    "net/http"
    "sync"
    "time"
)

type SessionStore struct {
    mu       sync.RWMutex
    sessions map[string]map[string]interface{}
    ttl      time.Duration
}

func NewSessionStore(ttl time.Duration) *SessionStore {
    return &SessionStore{
        sessions: make(map[string]map[string]interface{}),
        ttl:      ttl,
    }
}

func (s *SessionStore) Create(w http.ResponseWriter, data map[string]interface{}) string {
    id := generateSessionID()
    s.mu.Lock()
    s.sessions[id] = data
    s.mu.Unlock()

    cookie := &http.Cookie{
        Name:     "session",
        Value:    id,
        Path:     "/",
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteLaxMode,
        MaxAge:   int(s.ttl.Seconds()),
    }
    http.SetCookie(w, cookie)
    return id
}

func (s *SessionStore) Get(r *http.Request) (map[string]interface{}, bool) {
    cookie, err := r.Cookie("session")
    if err != nil {
        return nil, false
    }
    s.mu.RLock()
    data, ok := s.sessions[cookie.Value]
    s.mu.RUnlock()
    return data, ok
}

func (s *SessionStore) GetValue(r *http.Request, key string) (interface{}, bool) {
    data, ok := s.Get(r)
    if !ok {
        return nil, false
    }
    val, ok := data[key]
    return val, ok
}

func (s *SessionStore) SetValue(r *http.Request, key string, value interface{}) bool {
    cookie, err := r.Cookie("session")
    if err != nil {
        return false
    }
    s.mu.Lock()
    data, ok := s.sessions[cookie.Value]
    if !ok {
        s.mu.Unlock()
        return false
    }
    data[key] = value
    s.mu.Unlock()
    return true
}

func (s *SessionStore) Delete(r *http.Request, w http.ResponseWriter) {
    cookie, err := r.Cookie("session")
    if err != nil {
        return
    }
    s.mu.Lock()
    delete(s.sessions, cookie.Value)
    s.mu.Unlock()

    http.SetCookie(w, &http.Cookie{
        Name:   "session",
        Value:  "",
        Path:   "/",
        MaxAge: -1,
    })
}

func generateSessionID() string {
    b := make([]byte, 32)
    if _, err := rand.Read(b); err != nil {
        panic(err)
    }
    return base64.URLEncoding.EncodeToString(b)
}
