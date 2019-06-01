package models

import (
	"errors"
)

type ResponseType int

const (
	WhoAmIResponseType ResponseType = iota
	GetConnectedClientsResponseType
	RelayResponseType
)

var ResponseDecodingError = errors.New("decoding response failed")

type Response interface {
	ResponseType() ResponseType
	Body() interface{}
	Targets() []uint64
}

type ResponseStruct struct {
	Type ResponseType `json:"type"`
	Body interface{}  `json:"body"`
}
