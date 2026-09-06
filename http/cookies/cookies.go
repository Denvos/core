package cookies

import (
    "net/http"
    "time"
)

type Cookie struct {
    Name     string
    Value    string
    Path     string
    Domain   string
    Expires  time.Time
    MaxAge   int
    Secure   bool
    HttpOnly bool
    SameSite http.SameSite
}

func New(name, value string) *Cookie {
    return &Cookie{
        Name:     name,
        Value:    value,
        Path:     "/",
        SameSite: http.SameSiteLaxMode,
    }
}

func (c *Cookie) SetPath(path string) *Cookie {
    c.Path = path
    return c
}

func (c *Cookie) SetDomain(domain string) *Cookie {
    c.Domain = domain
    return c
}

func (c *Cookie) SetExpires(t time.Time) *Cookie {
    c.Expires = t
    c.MaxAge = 0
    return c
}

func (c *Cookie) SetMaxAge(seconds int) *Cookie {
    c.MaxAge = seconds
    c.Expires = time.Time{}
    return c
}

func (c *Cookie) SetSecure(secure bool) *Cookie {
    c.Secure = secure
    return c
}

func (c *Cookie) SetHttpOnly(httpOnly bool) *Cookie {
    c.HttpOnly = httpOnly
    return c
}

func (c *Cookie) SetSameSite(sameSite http.SameSite) *Cookie {
    c.SameSite = sameSite
    return c
}

func (c *Cookie) ToHTTPCookie() *http.Cookie {
    return &http.Cookie{
        Name:     c.Name,
        Value:    c.Value,
        Path:     c.Path,
        Domain:   c.Domain,
        Expires:  c.Expires,
        MaxAge:   c.MaxAge,
        Secure:   c.Secure,
        HttpOnly: c.HttpOnly,
        SameSite: c.SameSite,
    }
}
