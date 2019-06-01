package tcpclient_test

import (
	"errors"
	"github.com/akashdeep-singh/errol/pkg/client/tcpclient"
	"github.com/akashdeep-singh/errol/pkg/mocks"
	"github.com/akashdeep-singh/errol/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var fakeEncodedWhoAmIRequest = []byte{123, 34, 116, 121, 112, 101, 34, 58, 48, 44, 34, 98, 111, 100, 121, 34, 58, 123, 34, 82, 101, 99, 101, 105, 118, 101, 114, 115, 34, 58, 110, 117, 108, 108, 44, 34, 77, 101, 115, 115, 97, 103, 101, 34, 58, 34, 34, 125, 125}

var fakeEncodedGetConnectedClientsRequest = []byte{123, 34, 116, 121, 112, 101, 34, 58, 49, 44, 34, 98, 111, 100, 121, 34, 58, 123, 34, 82, 101, 99, 101, 105, 118, 101, 114, 115, 34, 58, 110, 117, 108, 108, 44, 34, 77, 101, 115, 115, 97, 103, 101, 34, 58, 34, 34, 125, 125}

var fakeEncodedRelayRequest = []byte{123, 34, 116, 121, 112, 101, 34, 58, 50, 44, 34, 98, 111, 100, 121, 34, 58, 123, 34, 82, 101, 99, 101, 105, 118, 101, 114, 115, 34, 58, 91, 50, 44, 52, 44, 54, 44, 56, 93, 44, 34, 77, 101, 115, 115, 97, 103, 101, 34,
	58, 34, 84, 101, 115, 116, 32, 109, 101, 115, 115, 97, 103, 101, 34, 125, 125}

func TestTCPClient_CreateConnection_error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testClient := tcpclient.TCPClient{
		Io:            &mocks.MockIO{},
		RequestCodec:  &mocks.MockRequestCodec{},
		ResponseCodec: &mocks.MockResponseCodec{},
	}

	err := testClient.CreateConnection("22")
	assert.Error(t, err)

}

func TestTCPClient_WhoAmI_success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockIo := mocks.NewMockIO(mockCtrl)
	mockRequestCodec := mocks.NewMockRequestCodec(mockCtrl)
	mockResponseCodec := mocks.NewMockResponseCodec(mockCtrl)
	testClient := &tcpclient.TCPClient{
		Io:            mockIo,
		RequestCodec:  mockRequestCodec,
		ResponseCodec: mockResponseCodec,
	}
	fakeRequest := models.NewWhoAmIRequest()
	mockRequestCodec.EXPECT().EncodeRequest(fakeRequest).Return(fakeEncodedWhoAmIRequest).Times(1)
	mockIo.EXPECT().Write(gomock.Any(), fakeEncodedWhoAmIRequest).Return(nil).Times(1)
	err := testClient.WhoAmI()
	assert.Nil(t, err)
}

func TestTCPClient_WhoAmI_errorInWrite(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockIo := mocks.NewMockIO(mockCtrl)
	mockRequestCodec := mocks.NewMockRequestCodec(mockCtrl)
	mockResponseCodec := mocks.NewMockResponseCodec(mockCtrl)
	testClient := &tcpclient.TCPClient{
		Io:            mockIo,
		RequestCodec:  mockRequestCodec,
		ResponseCodec: mockResponseCodec,
	}
	fakeRequest := models.NewWhoAmIRequest()
	mockRequestCodec.EXPECT().EncodeRequest(fakeRequest).Return([]byte{}).Times(1)
	mockIo.EXPECT().Write(gomock.Any(), []byte{}).Return(errors.New("some error")).Times(1)
	err := testClient.WhoAmI()
	assert.Error(t, err)
}

func TestTCPClient_GetConnectedClients_success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockIo := mocks.NewMockIO(mockCtrl)
	mockRequestCodec := mocks.NewMockRequestCodec(mockCtrl)
	mockResponseCodec := mocks.NewMockResponseCodec(mockCtrl)
	testClient := &tcpclient.TCPClient{
		Io:            mockIo,
		RequestCodec:  mockRequestCodec,
		ResponseCodec: mockResponseCodec,
	}
	fakeRequest := models.NewGetConnectedClientsRequest()
	mockRequestCodec.EXPECT().EncodeRequest(fakeRequest).Return(fakeEncodedGetConnectedClientsRequest).Times(1)
	mockIo.EXPECT().Write(gomock.Any(), fakeEncodedGetConnectedClientsRequest).Return(nil).Times(1)
	err := testClient.GetConnectedClients()
	assert.Nil(t, err)
}

func TestTCPClient_GetConnectedClients_errorInWrite(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockIo := mocks.NewMockIO(mockCtrl)
	mockRequestCodec := mocks.NewMockRequestCodec(mockCtrl)
	mockResponseCodec := mocks.NewMockResponseCodec(mockCtrl)
	testClient := &tcpclient.TCPClient{
		Io:            mockIo,
		RequestCodec:  mockRequestCodec,
		ResponseCodec: mockResponseCodec,
	}
	fakeRequest := models.NewGetConnectedClientsRequest()
	mockRequestCodec.EXPECT().EncodeRequest(fakeRequest).Return(fakeEncodedGetConnectedClientsRequest).Times(1)
	mockIo.EXPECT().Write(gomock.Any(), fakeEncodedGetConnectedClientsRequest).Return(errors.New("some error")).Times(1)
	err := testClient.GetConnectedClients()
	assert.Error(t, err)
}

func TestTCPClient_RelayMessage_success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockIo := mocks.NewMockIO(mockCtrl)
	mockRequestCodec := mocks.NewMockRequestCodec(mockCtrl)
	mockResponseCodec := mocks.NewMockResponseCodec(mockCtrl)
	testClient := &tcpclient.TCPClient{
		Io:            mockIo,
		RequestCodec:  mockRequestCodec,
		ResponseCodec: mockResponseCodec,
	}
	receivers := []uint64{2, 4, 6, 8}
	message := "Test message"
	fakeRequest := models.NewRelayRequest(receivers, message)
	mockRequestCodec.EXPECT().EncodeRequest(fakeRequest).Return(fakeEncodedRelayRequest).Times(1)
	mockIo.EXPECT().Write(gomock.Any(), fakeEncodedRelayRequest).Return(nil).Times(1)
	err := testClient.RelayMessage(receivers, message)
	assert.Nil(t, err)
}

func TestTCPClient_RelayMessage_errorInWrite(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockIo := mocks.NewMockIO(mockCtrl)
	mockRequestCodec := mocks.NewMockRequestCodec(mockCtrl)
	mockResponseCodec := mocks.NewMockResponseCodec(mockCtrl)
	testClient := &tcpclient.TCPClient{
		Io:            mockIo,
		RequestCodec:  mockRequestCodec,
		ResponseCodec: mockResponseCodec,
	}
	receivers := []uint64{2, 4, 6, 8}
	message := "Test message"
	fakeRequest := models.NewRelayRequest(receivers, message)
	mockRequestCodec.EXPECT().EncodeRequest(fakeRequest).Return(fakeEncodedRelayRequest).Times(1)
	mockIo.EXPECT().Write(gomock.Any(), fakeEncodedRelayRequest).Return(errors.New("some error")).Times(1)
	err := testClient.RelayMessage(receivers, message)
	assert.Error(t, err)
}
