package server

//go:generate mockgen -source server.go -destination ../mocks/mock_server.go -package mocks

type Server interface {
	Initialize(portNumber string) error

	GetConnectedClients() []uint64

	Exit() error
}
