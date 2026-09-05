package retry

import (
	"time"
)

type Config struct {
	MaxAttempts int
	Delay       time.Duration
	Backoff     float64
	Jitter      time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		MaxAttempts: 3,
		Delay:       100 * time.Millisecond,
		Backoff:     2.0,
		Jitter:      10 * time.Millisecond,
	}
}

func (c *Config) NextDelay(attempt int) time.Duration {
	d := time.Duration(float64(c.Delay) * (c.Backoff * float64(attempt)))
	if c.Jitter > 0 {
		j := time.Duration(float64(c.Jitter) * (1 + float64(attempt)*0.5))
		return d + j
	}
	return d
}
