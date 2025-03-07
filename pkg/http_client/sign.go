package http_client

import "net/http"

type Sign interface {
	Sign(request *http.Request)
}
