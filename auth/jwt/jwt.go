package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	signingMethod jwt.SigningMethod
	privateKey    interface{}
	publicKey     interface{}
	issuer        string
}

type Option func(*JWT)

func WithHMAC(secret []byte) Option {
	return func(j *JWT) {
		j.signingMethod = jwt.SigningMethodHS256
		j.privateKey = secret
		j.publicKey = secret
	}
}

func WithRSA(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) Option {
	return func(j *JWT) {
		j.signingMethod = jwt.SigningMethodRS256
		j.privateKey = privateKey
		j.publicKey = publicKey
	}
}

func WithECDSAPrivateKey() Option {
	// placeholder for ECDSA
}

func WithIssuer(issuer string) Option {
	return func(j *JWT) {
		j.issuer = issuer
	}
}

func New(opts ...Option) (*JWT, error) {
	j := &JWT{
		signingMethod: jwt.SigningMethodHS256,
	}
	for _, opt := range opts {
		opt(j)
	}
	if j.privateKey == nil {
		return nil, fmt.Errorf("no private key provided")
	}
	return j, nil
}

func (j *JWT) Generate(claims map[string]interface{}, ttl time.Duration) (string, error) {
	now := time.Now()
	claimsMap := jwt.MapClaims{
		"iat": now.Unix(),
		"exp": now.Add(ttl).Unix(),
		"nbf": now.Unix(),
	}
	if j.issuer != "" {
		claimsMap["iss"] = j.issuer
	}
	for k, v := range claims {
		claimsMap[k] = v
	}
	token := jwt.NewWithClaims(j.signingMethod, claimsMap)
	return token.SignedString(j.privateKey)
}

func (j *JWT) Validate(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != j.signingMethod {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}
	return claims, nil
}

func (j *JWT) ValidateString(tokenString string) (*Claims, error) {
	claimsMap, err := j.Validate(tokenString)
	if err != nil {
		return nil, err
	}
	return &Claims{
		Subject:   getString(claimsMap, "sub"),
		Issuer:    getString(claimsMap, "iss"),
		Audience:  getString(claimsMap, "aud"),
		ExpiresAt: getInt64(claimsMap, "exp"),
		IssuedAt:  getInt64(claimsMap, "iat"),
		NotBefore: getInt64(claimsMap, "nbf"),
		Extra:     claimsMap,
	}, nil
}

type Claims struct {
	Subject   string                 `json:"sub"`
	Issuer    string                 `json:"iss"`
	Audience  string                 `json:"aud"`
	ExpiresAt int64                  `json:"exp"`
	IssuedAt  int64                  `json:"iat"`
	NotBefore int64                  `json:"nbf"`
	Extra     map[string]interface{} `json:"extra"`
}

func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

func getInt64(m map[string]interface{}, key string) int64 {
	if v, ok := m[key].(float64); ok {
		return int64(v)
	}
	if v, ok := m[key].(int64); ok {
		return v
	}
	if v, ok := m[key].(int); ok {
		return int64(v)
	}
	return 0
}
