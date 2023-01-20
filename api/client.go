package api

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"golang.org/x/time/rate"
)

const (
	EnvVaultAddress   = "VAULT_ADDR"
	EnvVaultAgentAddr = "VAULT_AGENT_ADDR"
	EnvVaultCACert    = "VAULT_CACERT"
)

// Deprecated values
const (
	EnvVaultAgentAddress = "VAULT_AGENT_ADDR"
	EnvVaultInsecure     = "VAULT_SKIP_VERIFY"
)

type WrappingLookupFunc func(operation, path string) string

type Config struct {
	modifyLock sync.RWMutex

	Address string

	AgentAddress string

	HttpClient *http.Client

	MinRetryWait time.Duration

	MaxRetryWait time.Duration

	MaxRetries int

	Timeout time.Duration

	Error error

	Backoff retryablehttp.Backoff

	CheckRetry retryablehttp.CheckRetry

	Logger retryablehttp.LeveledLogger

	Limiter *rate.Limiter

	OutputCurlString bool

	OutputPolicy bool
}

type Client struct {
	modifyLock         sync.RWMutex
	addr               *url.URL
	config             *Config
	token              string
	headers            http.Header
	wrappingLookupFunc WrappingLookupFunc
	mfaCreds           []string
	policyOverride bool
	requestCallbacks []Req
}

type (
	RequestCallback func(*Request)
)