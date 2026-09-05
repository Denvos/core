package splunk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Splunk struct {
	url      string
	token    string
	index    string
	source   string
	sourcetype string
	host     string
	client   *http.Client
}

type Option func(*Splunk)

func WithIndex(index string) Option {
	return func(s *Splunk) {
		s.index = index
	}
}

func WithSource(source string) Option {
	return func(s *Splunk) {
		s.source = source
	}
}

func WithSourcetype(sourcetype string) Option {
	return func(s *Splunk) {
		s.sourcetype = sourcetype
	}
}

func WithHost(host string) Option {
	return func(s *Splunk) {
		s.host = host
	}
}

func New(url, token string, opts ...Option) *Splunk {
	s := &Splunk{
		url:      url,
		token:    token,
		client:   &http.Client{},
		source:   "denvos",
		sourcetype: "denvos-log",
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Splunk) Write(p []byte) (int, error) {
	var payload map[string]interface{}
	if err := json.Unmarshal(p, &payload); err != nil {
		payload = map[string]interface{}{
			"message": string(p),
		}
	}
	if s.index != "" {
		payload["index"] = s.index
	}
	if s.source != "" {
		payload["source"] = s.source
	}
	if s.sourcetype != "" {
		payload["sourcetype"] = s.sourcetype
	}
	if s.host != "" {
		payload["host"] = s.host
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest("POST", s.url+"/services/collector", bytes.NewReader(data))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Authorization", "Splunk "+s.token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("splunk error: %s", body)
	}
	return len(p), nil
}
