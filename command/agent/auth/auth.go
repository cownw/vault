package auth

import (
	"context"
	"time"
)

const (
	defaultMinBackoff = 1 * time.Second
	defaultMaxBackoff = 5 * time.Minute
)

type AuthMethod interface {
	Authenticate(context.Context, *api)
}