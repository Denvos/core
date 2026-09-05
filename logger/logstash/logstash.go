package logstash

import (
	"encoding/json"
	"net"
	"time"
)

type Logstash struct {
	conn   net.Conn
	host   string
	port   string
	timeout time.Duration
}

func New(host, port string) (*Logstash, error) {
	conn, err := net.DialTimeout("tcp", host+":"+port, 5*time.Second)
	if err != nil {
		return nil, err
	}
	return &Logstash{
		conn:   conn,
		host:   host,
		port:   port,
		timeout: 5 * time.Second,
	}, nil
}

func (l *Logstash) Write(p []byte) (int, error) {
	var entry map[string]interface{}
	if err := json.Unmarshal(p, &entry); err != nil {
		entry = map[string]interface{}{
			"message": string(p),
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		}
	}
	data, err := json.Marshal(entry)
	if err != nil {
		return 0, err
	}
	l.conn.SetWriteDeadline(time.Now().Add(l.timeout))
	_, err = l.conn.Write(append(data, '\n'))
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (l *Logstash) Close() error {
	return l.conn.Close()
}
