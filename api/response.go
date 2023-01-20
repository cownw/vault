package api

import "net/http"

// Response is a raw response that wraps an HTTP response.
type Response struct {
	*http.Response
}

func (r *Response) DecodeJSON(out interface{}) error {
	return
}