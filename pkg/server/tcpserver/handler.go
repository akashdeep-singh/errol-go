package tcpserver

import "github.com/akashdeep-singh/errol/pkg/models"

//go:generate mockgen -source handler.go -destination ../../mocks/mock_handler.go -package mocks RequestHandler

type RequestHandler interface {
	WhoAmI(id uint64) models.WhoAmIResponse
	GetConnectedClients(id uint64, clients []uint64) models.GetConnectedClientsResponse
	Relay(request models.RelayRequest) models.RelayResponse
}

type ErrolHandler struct {
	RequestHandler
}

func (h *ErrolHandler) WhoAmI(id uint64) models.WhoAmIResponse {
	return models.WhoAmIResponse{Id: id, Target: id}
}

func (h *ErrolHandler) GetConnectedClients(id uint64, clients []uint64) models.GetConnectedClientsResponse {
	var connectedClients []uint64
	for _, clientId := range clients {
		if clientId != id {
			connectedClients = append(connectedClients, clientId)
		}
	}
	return models.GetConnectedClientsResponse{ConnectedClients: connectedClients, Target: id}
}

func (h *ErrolHandler) Relay(request models.RelayRequest) models.RelayResponse {
	return models.RelayResponse{Message: request.Message, Receivers: request.Receivers}
}
