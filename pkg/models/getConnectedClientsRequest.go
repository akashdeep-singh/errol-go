package models

type GetConnectedClientsRequest struct {
}

func NewGetConnectedClientsRequest() GetConnectedClientsRequest {
	return GetConnectedClientsRequest{}
}

func (_ GetConnectedClientsRequest) RequestType() RequestType {
	return GetConnectedClientsRequestType
}

func (_ GetConnectedClientsRequest) Body() interface{} {
	return nil
}
