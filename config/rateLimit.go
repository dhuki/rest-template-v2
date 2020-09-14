package config

import (
	"time"

	"github.com/juju/ratelimit"
)

// implement simple rate limiter
func NewRateLimit() *ratelimit.Bucket {
	return ratelimit.NewBucket(1*time.Minute, 5)
}
