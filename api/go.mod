module github.com/cownw/vault/api

go 1.19

replace github.com/cownw/vault/sdk => ../sdk

require (
	github.com/hashicorp/go-retryablehttp v0.7.2
	golang.org/x/time v0.3.0

	github.com/cownw/vault/sdk v0.7.0
)

require github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
