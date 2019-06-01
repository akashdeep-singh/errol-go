package tcpserver_test

import (
	"github.com/akashdeep-singh/errol-go/pkg/models"
	"github.com/akashdeep-singh/errol-go/pkg/server/tcpserver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrolHandler_WhoAmI(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testHandler := tcpserver.ErrolHandler{}
	expectedResult := models.WhoAmIResponse{2, 2}
	result := testHandler.WhoAmI(2)
	assert.Equal(t, expectedResult, result)
}

func TestErrolHandler_GetConnectedClients(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testHandler := tcpserver.ErrolHandler{}
	expectedResult := models.GetConnectedClientsResponse{ConnectedClients: []uint64{1, 3, 4, 5}, Target: 2}
	result := testHandler.GetConnectedClients(2, []uint64{1, 3, 4, 5})
	assert.Equal(t, expectedResult, result)
}

func TestErrolHandler_Relay(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testHandler := tcpserver.ErrolHandler{}
	expectedResult := models.RelayResponse{Message: "Bonjour", Receivers: []uint64{2, 4, 6, 8}}
	result := testHandler.Relay(models.RelayRequest{Message: "Bonjour", Receivers: []uint64{2, 4, 6, 8}})
	assert.Equal(t, expectedResult, result)
}
