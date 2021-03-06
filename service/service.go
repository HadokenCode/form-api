package service

import (
	"github.com/kamilsk/form-api/domain"
	"github.com/kamilsk/form-api/errors"
	"github.com/kamilsk/form-api/transfer/api/v1"
)

// New returns a new instance of Form API service.
func New(dao Storage) *FormAPI {
	return &FormAPI{dao: dao}
}

// FormAPI is the primary application service.
type FormAPI struct {
	dao Storage
}

// HandleGetV1 handles an input request.
func (s *FormAPI) HandleGetV1(request v1.GetRequest) v1.GetResponse {
	var response v1.GetResponse
	response.Schema, response.Error = s.dao.Schema(request.UUID)
	return response
}

// HandlePostV1 handles an input request.
func (s *FormAPI) HandlePostV1(request v1.PostRequest) v1.PostResponse {
	var (
		response v1.PostResponse
		verified map[string][]string
	)

	{ // TODO encrypt/decrypt marker
		marker := domain.UUID(request.EncryptedMarker)
		if !marker.IsValid() {
			marker, response.Error = s.dao.UUID()
			if response.Error != nil {
				return response
			}
		}
		response.EncryptedMarker = string(marker)
	}

	response.Schema, response.Error = s.dao.Schema(request.UUID)
	if response.Error != nil {
		return response
	}
	verified, response.Error = response.Schema.Apply(request.Data)
	if response.Error != nil {
		response.Error = errors.Validation(errors.InvalidFormDataMessage, response.Error,
			"trying to add data for schema %q", request.UUID)
		return response
	}

	// issue #110: add cookie
	// TODO use context column
	verified["_token"] = []string{response.EncryptedMarker}

	response.ID, response.Error = s.dao.AddData(request.UUID, verified)
	return response
}
