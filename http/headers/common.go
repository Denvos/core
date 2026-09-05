package headers

import (
    "net/http"
    "strings"
)

func GetContentType(r *http.Request) string {
    return r.Header.Get(HeaderContentType)
}

func SetContentType(w http.ResponseWriter, contentType string) {
    w.Header().Set(HeaderContentType, contentType)
}

func GetHeader(r *http.Request, key string) string {
    return r.Header.Get(key)
}

func SetHeader(w http.ResponseWriter, key, value string) {
    w.Header().Set(key, value)
}

func AddHeader(w http.ResponseWriter, key, value string) {
    w.Header().Add(key, value)
}

func GetBearerToken(r *http.Request) string {
    auth := r.Header.Get(HeaderAuthorization)
    if len(auth) > 7 && strings.EqualFold(auth[:7], "Bearer ") {
        return auth[7:]
    }
    return ""
}

func GetBasicAuth(r *http.Request) (username, password string, ok bool) {
    return r.BasicAuth()
}

func GetRequestID(r *http.Request) string {
    id := r.Header.Get(HeaderXRequestID)
    if id == "" {
        id = r.Header.Get(HeaderXCorrelationID)
    }
    return id
}

func SetRequestID(w http.ResponseWriter, id string) {
    if id != "" {
        w.Header().Set(HeaderXRequestID, id)
    }
}
