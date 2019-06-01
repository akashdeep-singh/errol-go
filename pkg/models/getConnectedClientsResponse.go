package models

type GetConnectedClientsResponse struct {
	ConnectedClients []uint64
	Target           uint64
}

func NewGetConnectedClientsResponse(data []uint64) GetConnectedClientsResponse {
	return GetConnectedClientsResponse{ConnectedClients: data}
}

func (_ GetConnectedClientsResponse) ResponseType() ResponseType {
	return GetConnectedClientsResponseType
}

func (getConnectedClientsResponse GetConnectedClientsResponse) Body() interface{} {
	return getConnectedClientsResponse.ConnectedClients
}

func (getConnectedClientsResponse GetConnectedClientsResponse) Targets() []uint64 {
	var targets []uint64
	return append(targets, getConnectedClientsResponse.Target)
}
