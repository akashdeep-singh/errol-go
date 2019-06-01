package serialization

import (
	"github.com/akashdeep-singh/errol/pkg/models"
)

//go:generate mockgen -source response.go -destination ../mocks/mock_response.go -package mocks

type ResponseCodec interface {
	EncodeResponse(response models.Response) []byte
	DecodeResponse(encoded []byte) (models.Response, error)
}
