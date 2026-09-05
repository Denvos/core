package silent

import "io"

type Silent struct{}

func New() *Silent {
	return &Silent{}
}

func (s *Silent) Write(p []byte) (int, error) {
	return len(p), nil
}
