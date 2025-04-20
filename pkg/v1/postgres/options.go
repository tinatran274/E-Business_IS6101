package postgres

import (
	"time"

	custom_log "10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/log"
)

type Option func(*Postgres)

func MaxPoolSize(size int) Option {
	return func(c *Postgres) {
		c.maxPoolSize = size
	}
}

func ConnAttempts(attempts int) Option {
	return func(c *Postgres) {
		c.connAttempts = attempts
	}
}

func ConnTimeout(timeout time.Duration) Option {
	return func(c *Postgres) {
		c.connTimeout = timeout
	}
}

func WithTracer(tracer *custom_log.LogTracer) Option {
	return func(pg *Postgres) {
		pg.tracer = tracer
	}
}
