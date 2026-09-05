package disk

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type Cache struct {
	mu   sync.RWMutex
	dir  string
	data map[string][]byte
}

func New(dir string) (*Cache, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	c := &Cache{
		dir:  dir,
		data: make(map[string][]byte),
	}
	if err := c.load(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Cache) Set(key string, value interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	c.data[key] = data
	return ioutil.WriteFile(filepath.Join(c.dir, key), data, 0644)
}

func (c *Cache) Get(key string, dest interface{}) (bool, error) {
	c.mu.RLock()
	data, ok := c.data[key]
	c.mu.RUnlock()
	if !ok {
		return false, nil
	}
	if err := json.Unmarshal(data, dest); err != nil {
		return false, err
	}
	return true, nil
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
	return os.Remove(filepath.Join(c.dir, key))
}

func (c *Cache) Clear() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string][]byte)
	return os.RemoveAll(c.dir)
}

func (c *Cache) load() error {
	files, err := ioutil.ReadDir(c.dir)
	if err != nil {
		return err
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		data, err := ioutil.ReadFile(filepath.Join(c.dir, f.Name()))
		if err != nil {
			continue
		}
		c.data[f.Name()] = data
	}
	return nil
}
