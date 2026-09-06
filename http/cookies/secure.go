package cookies

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "fmt"
    "strings"
)

type SecureCookie struct {
    secret []byte
}

func NewSecureCookie(secret []byte) *SecureCookie {
    return &SecureCookie{secret: secret}
}

func (s *SecureCookie) Encode(name, value string) (string, error) {
    // Simple encoding: value + signature
    h := hmac.New(sha256.New, s.secret)
    h.Write([]byte(name + ":" + value))
    sig := base64.URLEncoding.EncodeToString(h.Sum(nil))
    return base64.URLEncoding.EncodeToString([]byte(value + "|" + sig)), nil
}

func (s *SecureCookie) Decode(name, encoded string) (string, error) {
    data, err := base64.URLEncoding.DecodeString(encoded)
    if err != nil {
        return "", err
    }
    parts := strings.SplitN(string(data), "|", 2)
    if len(parts) != 2 {
        return "", fmt.Errorf("invalid cookie format")
    }
    value := parts[0]
    sig := parts[1]

    h := hmac.New(sha256.New, s.secret)
    h.Write([]byte(name + ":" + value))
    expected := base64.URLEncoding.EncodeToString(h.Sum(nil))

    if !hmac.Equal([]byte(sig), []byte(expected)) {
        return "", fmt.Errorf("invalid signature")
    }
    return value, nil
}
