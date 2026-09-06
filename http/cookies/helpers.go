package cookies

import (
    "net/http"
    "strings"
)

func Get(r *http.Request, name string) string {
    cookie, err := r.Cookie(name)
    if err != nil {
        return ""
    }
    return cookie.Value
}

func GetCookie(r *http.Request, name string) (*http.Cookie, error) {
    return r.Cookie(name)
}

func Set(w http.ResponseWriter, cookie *http.Cookie) {
    http.SetCookie(w, cookie)
}

func SetString(w http.ResponseWriter, name, value string) {
    http.SetCookie(w, &http.Cookie{
        Name:     name,
        Value:    value,
        Path:     "/",
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteLaxMode,
    })
}

func Delete(w http.ResponseWriter, name string) {
    http.SetCookie(w, &http.Cookie{
        Name:     name,
        Value:    "",
        Path:     "/",
        MaxAge:   -1,
        HttpOnly: true,
        Secure:   true,
    })
}

func GetAll(r *http.Request) []*http.Cookie {
    return r.Cookies()
}

func ParseHeader(header string) map[string]string {
    cookies := make(map[string]string)
    if header == "" {
        return cookies
    }
    for _, part := range strings.Split(header, ";") {
        part = strings.TrimSpace(part)
        if part == "" {
            continue
        }
        parts := strings.SplitN(part, "=", 2)
        if len(parts) == 2 {
            cookies[parts[0]] = parts[1]
        } else {
            cookies[part] = ""
        }
    }
    return cookies
}
