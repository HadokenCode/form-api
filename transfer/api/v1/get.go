package v1

import "github.com/kamilsk/form-api/domain"

// GetRequest represents `GET /api/v1/{Schema.ID}` request.
type GetRequest struct {
	UUID domain.UUID
}

// GetResponse represents `GET /api/v1/{Schema.ID}` response.
type GetResponse struct {
	Schema domain.Schema
	Error  error
}
