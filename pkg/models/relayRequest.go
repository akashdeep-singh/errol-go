package models

type RelayRequest struct {
	Receivers []uint64
	Message   string
}

func NewRelayRequest(receivers []uint64, message string) RelayRequest {
	return RelayRequest{receivers, message}
}

func (_ RelayRequest) RequestType() RequestType {
	return RelayRequestType
}

func (relayRequest RelayRequest) Body() interface{} {
	return relayRequest
}
