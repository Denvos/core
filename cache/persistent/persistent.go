package persistent

import (
	"encoding/gob"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type Store struct {
	mu   sync.RWMutex
	file string
	data map[string][]byte
}

func New(file string) (*Store, error) {
	s := &Store{
		file: file,
		data: make(map[string][]byte),
	}
	if err := s.load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	return s, nil
}

func (s *Store) Set(key string, value []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	return s.save()
}

func (s *Store) Get(key string) ([]byte, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.data[key]
	return v, ok
}

func (s *Store) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	return s.save()
}

func (s *Store) Clear() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = make(map[string][]byte)
	return s.save()
}

func (s *Store) load() error {
	data, err := ioutil.ReadFile(s.file)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(bytes.NewReader(data))
	return dec.Decode(&s.data)
}

func (s *Store) save() error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(s.data); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(s.file), 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(s.file, buf.Bytes(), 0644)
}
