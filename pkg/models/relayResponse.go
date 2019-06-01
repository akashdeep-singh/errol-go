package models

type RelayResponse struct {
	Message   string
	Receivers []uint64
}

func NewRelayResponse(data string) RelayResponse {
	return RelayResponse{Message: data}
}

func (_ RelayResponse) ResponseType() ResponseType {
	return RelayResponseType
}

func (relayRequest RelayResponse) Body() interface{} {
	return relayRequest.Message
}

func (relayRequest RelayResponse) Targets() []uint64 {
	return relayRequest.Receivers
}
