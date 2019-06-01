package JSONcodec_test

import (
	"github.com/akashdeep-singh/errol/pkg/models"
	"github.com/akashdeep-singh/errol/pkg/serialization/JSONcodec"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var fakeEncodedWhoAmIRequest = []byte{123, 34, 116, 121, 112, 101, 34, 58, 48, 44, 34, 98, 111, 100, 121, 34, 58, 123, 34, 82, 101, 99, 101, 105, 118, 101, 114, 115, 34, 58, 110, 117, 108, 108, 44, 34, 77, 101, 115, 115, 97, 103, 101, 34, 58, 34, 34, 125, 125}

var fakeEncodedGetConnectedClientsRequest = []byte{123, 34, 116, 121, 112, 101, 34, 58, 49, 44, 34, 98, 111, 100, 121, 34, 58, 123, 34, 82, 101, 99, 101, 105, 118, 101, 114, 115, 34, 58, 110, 117, 108, 108, 44, 34, 77, 101, 115, 115, 97, 103, 101, 34, 58, 34, 34, 125, 125}

var fakeEncodedRelayRequest = []byte{123, 34, 116, 121, 112, 101, 34, 58, 50, 44, 34, 98, 111, 100, 121, 34, 58, 123, 34, 82, 101, 99, 101, 105, 118, 101, 114, 115, 34, 58, 91, 50, 44, 52, 44, 54, 44, 56, 93, 44, 34, 77, 101, 115, 115, 97, 103, 101, 34,
	58, 34, 84, 101, 115, 116, 32, 109, 101, 115, 115, 97, 103, 101, 34, 125, 125}

func TestJSONRequestCodec_EncodeRequest_nonNilBody(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	fakeRequest := models.NewRelayRequest([]uint64{2, 4, 6, 8}, "Test message")
	requestCodec := JSONcodec.JSONRequestCodec{}
	result := requestCodec.EncodeRequest(fakeRequest)
	assert.Equal(t, fakeEncodedRelayRequest, result)
}

func TestJSONRequestCodec_EncodeRequest_nilBody(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	fakeRequest := models.NewWhoAmIRequest()
	requestCodec := JSONcodec.JSONRequestCodec{}
	result := requestCodec.EncodeRequest(fakeRequest)
	assert.Equal(t, fakeEncodedWhoAmIRequest, result)
}

func TestJSONRequestCodec_DecodeRequest_WhoAmIRequest_success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	encodedRequest := fakeEncodedWhoAmIRequest
	requestCodec := JSONcodec.JSONRequestCodec{}
	result, err := requestCodec.DecodeRequest(encodedRequest)
	assert.Nil(t, err)
	assert.Equal(t, models.WhoAmIRequestType, result.RequestType())
}

func TestJSONRequestCodec_DecodeRequest_GetConnectedClientsRequest_success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	encodedRequest := fakeEncodedGetConnectedClientsRequest
	requestCodec := JSONcodec.JSONRequestCodec{}
	result, err := requestCodec.DecodeRequest(encodedRequest)
	assert.Nil(t, err)
	assert.Equal(t, models.GetConnectedClientsRequestType, result.RequestType())
}

func TestJSONRequestCodec_DecodeRequest_RelayRequest_success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	encodedRequest := fakeEncodedRelayRequest
	requestCodec := JSONcodec.JSONRequestCodec{}
	result, err := requestCodec.DecodeRequest(encodedRequest)
	assert.Nil(t, err)
	assert.Equal(t, models.RelayRequestType, result.RequestType())
}

func TestJSONRequestCodec_DecodeRequest_incorrectRequestType_error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	encodedRequest := []byte{123, 34, 116, 121, 112, 101, 34, 58, 51, 44, 34, 98, 111, 100, 121, 34, 58, 123, 34, 82, 101, 99, 101, 105, 118, 101, 114, 115, 34, 58, 110, 117, 108, 108, 44, 34, 77, 101, 115, 115, 97, 103, 101, 34, 58, 34, 34, 125, 125}
	requestCodec := JSONcodec.JSONRequestCodec{}
	result, err := requestCodec.DecodeRequest(encodedRequest)
	assert.Error(t, err)
	assert.Equal(t, models.RequestDecodingError, err)
	assert.Equal(t, nil, result)
}
