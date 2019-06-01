package tcpserver_test

import (
	"github.com/akashdeep-singh/errol-go/pkg/mocks"
	"github.com/akashdeep-singh/errol-go/pkg/server/tcpserver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTCPServer_GetConnectedClients(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRequestHandler := mocks.NewMockRequestHandler(mockCtrl)
	mockIo := mocks.NewMockIO(mockCtrl)
	mockRequestCodec := mocks.NewMockRequestCodec(mockCtrl)
	mockResponseCodec := mocks.NewMockResponseCodec(mockCtrl)
	testServer := &tcpserver.TCPServer{
		RequestHandler: mockRequestHandler,
		Io:             mockIo,
		RequestCodec:   mockRequestCodec,
		ResponseCodec:  mockResponseCodec,
		IdProviderFunc: tcpserver.NewID,
	}

	clients := testServer.GetConnectedClients()
	assert.Empty(t, clients)
}
