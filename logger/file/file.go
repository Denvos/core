package file

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type File struct {
	path string
	file *os.File
}

func New(path string) (*File, error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return &File{path: path, file: f}, nil
}

func (f *File) Write(p []byte) (int, error) {
	return f.file.Write(p)
}

func (f *File) Close() error {
	return f.file.Close()
}

func Write(f *File, entry map[string]interface{}) {
	data, _ := json.Marshal(entry)
	f.file.Write(append(data, '\n'))
}
