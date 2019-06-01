package serialization

import (
	"github.com/akashdeep-singh/errol/pkg/models"
)

//go:generate mockgen -source request.go -destination ../mocks/mock_request.go -package mocks

type RequestCodec interface {
	EncodeRequest(request models.Request) []byte
	DecodeRequest(encoded []byte) (models.Request, error)
}
