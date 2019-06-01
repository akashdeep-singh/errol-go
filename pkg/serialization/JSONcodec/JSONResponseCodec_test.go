package JSONcodec_test

import (
	"github.com/akashdeep-singh/errol/pkg/models"
	"github.com/akashdeep-singh/errol/pkg/serialization/JSONcodec"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var fakeEncodedRelayResponse = []uint8([]byte{0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x32, 0x2c, 0x22, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x3a, 0x22, 0x54, 0x65, 0x73, 0x74, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x7d})

var fakeEncodedWhoAmIResponse = []uint8([]byte{0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x30, 0x2c, 0x22, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x3a, 0x32, 0x7d})

var fakeEncodedGetConnectedClientsResponse = []uint8([]byte{0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x31, 0x2c, 0x22, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x3a, 0x5b, 0x31, 0x2c, 0x32, 0x2c, 0x33, 0x5d, 0x7d})

func TestJSONResponseCodec_EncodeResponse_RelayResponse_success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	fakeResponse := models.NewRelayResponse("Test message")
	responseCodec := JSONcodec.JSONResponseCodec{}
	result := responseCodec.EncodeResponse(fakeResponse)
	assert.Equal(t, result, fakeEncodedRelayResponse)
}

func TestJSONResponseCodec_EncodeResponse_WhoAmIResponse_success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	fakeResponse := models.NewWhoAmIResponse(2)
	responseCodec := JSONcodec.JSONResponseCodec{}
	result := responseCodec.EncodeResponse(fakeResponse)
	assert.Equal(t, result, fakeEncodedWhoAmIResponse)
}

func TestJSONResponseCodec_EncodeResponse_GetConnectedClientsResponse_success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	fakeResponse := models.NewGetConnectedClientsResponse([]uint64{1, 2, 3})
	responseCodec := JSONcodec.JSONResponseCodec{}
	result := responseCodec.EncodeResponse(fakeResponse)
	assert.Equal(t, result, fakeEncodedGetConnectedClientsResponse)
}

func TestJSONResponseCodec_DecodeResponse_WhoAmIResponse_success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	encodedResponse := fakeEncodedWhoAmIResponse
	responseCodec := JSONcodec.JSONResponseCodec{}
	result, err := responseCodec.DecodeResponse(encodedResponse)
	assert.Nil(t, err)
	assert.Equal(t, models.WhoAmIResponseType, result.ResponseType())
}

func TestJSONResponseCodec_DecodeResponse_GetConnectedClientsResponse_success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	encodedResponse := fakeEncodedGetConnectedClientsResponse
	responseCodec := JSONcodec.JSONResponseCodec{}
	result, err := responseCodec.DecodeResponse(encodedResponse)
	assert.Nil(t, err)
	assert.Equal(t, models.GetConnectedClientsResponseType, result.ResponseType())
	assert.Equal(t, []uint64{1, 2, 3}, result.Body())
}

func TestJSONResponseCodec_DecodeResponse_RelayResponse_success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	encodedResponse := fakeEncodedRelayResponse
	responseCodec := JSONcodec.JSONResponseCodec{}
	result, err := responseCodec.DecodeResponse(encodedResponse)
	assert.Nil(t, err)
	assert.Equal(t, models.RelayResponseType, result.ResponseType())
	assert.Equal(t, "Test message", result.Body())
}

func TestJSONResponseCodec_DecodeResponse_incorrectResponseType_error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	encodedResponse := []uint8([]byte{0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x33, 0x2c, 0x22, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x3a, 0x22, 0x54, 0x65, 0x73, 0x74, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x7d})
	responseCodec := JSONcodec.JSONResponseCodec{}
	result, err := responseCodec.DecodeResponse(encodedResponse)
	assert.Error(t, err)
	assert.Equal(t, models.ResponseDecodingError, err)
	assert.Equal(t, nil, result)

}
