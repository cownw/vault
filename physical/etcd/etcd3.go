package etcd

import (
	"time"

	log "github.com/hashicorp/go-hclog"
)

type EtcdBackend struct {
	logger log.Logger
	path string
	haEnabled bool
	lockTimeout time.Duration
	requestTimeout time.Duration

	permitPool *
}