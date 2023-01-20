package api

import "errors"

var ErrSecretNotFound = errors.New("secret not found")

type KVSecret struct {
	Data map[string]interface{}
	VersionMetadata *
}