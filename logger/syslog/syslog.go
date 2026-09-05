package syslog

import (
	"io"
	"log/syslog"
)

type Syslog struct {
	w *syslog.Writer
}

func New(priority syslog.Priority, tag string) (*Syslog, error) {
	w, err := syslog.New(priority, tag)
	if err != nil {
		return nil, err
	}
	return &Syslog{w: w}, nil
}

func (s *Syslog) Write(p []byte) (int, error) {
	err := s.w.Info(string(p))
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (s *Syslog) Close() error {
	return s.w.Close()
}
