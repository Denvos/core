package invalidation

import "time"

type Strategy interface {
	ShouldInvalidate(created time.Time, hits int64) bool
}

type TimeBased struct {
	TTL time.Duration
}

func (t *TimeBased) ShouldInvalidate(created time.Time, hits int64) bool {
	return time.Since(created) > t.TTL
}

type CountBased struct {
	MaxHits int64
}

func (c *CountBased) ShouldInvalidate(created time.Time, hits int64) bool {
	return hits >= c.MaxHits
}

type Composite struct {
	Strategies []Strategy
}

func (c *Composite) ShouldInvalidate(created time.Time, hits int64) bool {
	for _, s := range c.Strategies {
		if s.ShouldInvalidate(created, hits) {
			return true
		}
	}
	return false
}
