package ldap

import (
	"crypto/tls"
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

type LDAP struct {
	conn   *ldap.Conn
	config Config
}

type Config struct {
	Host         string
	Port         int
	BaseDN       string
	BindDN       string
	BindPassword string
	UserFilter   string
	GroupFilter  string
	TLS          bool
	InsecureTL   bool
}

func New(cfg Config) (*LDAP, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	var conn *ldap.Conn
	var err error
	if cfg.TLS {
		tlsCfg := &tls.Config{InsecureSkipVerify: cfg.InsecureTL}
		conn, err = ldap.DialTLS("tcp", addr, tlsCfg)
	} else {
		conn, err = ldap.Dial("tcp", addr)
	}
	if err != nil {
		return nil, err
	}
	return &LDAP{conn: conn, config: cfg}, nil
}

func (l *LDAP) Bind() error {
	if l.config.BindDN == "" || l.config.BindPassword == "" {
		return nil // anonymous bind
	}
	return l.conn.Bind(l.config.BindDN, l.config.BindPassword)
}

func (l *LDAP) Authenticate(username, password string) (bool, error) {
	if err := l.Bind(); err != nil {
		return false, err
	}
	searchReq := ldap.NewSearchRequest(
		l.config.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf(l.config.UserFilter, username),
		[]string{"dn"},
		nil,
	)
	result, err := l.conn.Search(searchReq)
	if err != nil {
		return false, err
	}
	if len(result.Entries) == 0 {
		return false, nil
	}
	userDN := result.Entries[0].DN
	if err := l.conn.Bind(userDN, password); err != nil {
		return false, nil
	}
	return true, nil
}
