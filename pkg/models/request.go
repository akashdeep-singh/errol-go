package models

import (
	"errors"
)

type RequestType int

const (
	WhoAmIRequestType RequestType = iota
	GetConnectedClientsRequestType
	RelayRequestType
)

var RequestDecodingError = errors.New("decoding request failed")

type Request interface {
	RequestType() RequestType
	Body() interface{}
}

type RequestStruct struct {
	Type RequestType  `json:"type"`
	Body RelayRequest `json:"body"`
}
