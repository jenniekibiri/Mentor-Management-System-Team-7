package cache

import (
	"context"
	"time"
)

type Cache interface {
	// BlacklistSession adds a session token to the blacklist cache with an expiration duration.
	BlacklistSession(ctx context.Context, sessionToken string, expirationTime time.Duration) error

	// IsSessionBlacklisted checks if a session token is blacklisted by querying the Redis cache.
	IsSessionBlacklisted(ctx context.Context, sessionToken string) (bool, error)
}
