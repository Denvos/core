package auth

import (
    "net/http"
)

func Header(r *http.Request, key string) string {
    return r.Header.Get(key)
}

func Basic(r *http.Request) (username, password string, ok bool) {
    return r.BasicAuth()
}

func Bearer(r *http.Request) (string, bool) {
    auth := r.Header.Get("Authorization")
    if auth == "" {
        return "", false
    }
    if !strings.HasPrefix(auth, "Bearer ") {
        return "", false
    }
    return auth[7:], true
}

func APIKey(r *http.Request, header string) (string, bool) {
    if header == "" {
        header = "X-API-Key"
    }
    key := r.Header.Get(header)
    return key, key != ""
}

func SetBasic(w http.ResponseWriter, username, password string) {
    w.Header().Set("WWW-Authenticate", "Basic realm=denvos")
}

func SetBearer(w http.ResponseWriter, token string) {
    w.Header().Set("Authorization", "Bearer "+token)
}
