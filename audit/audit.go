package audit

import "context"

type Backend interface {
	LogRequest(context.Context, *)
}