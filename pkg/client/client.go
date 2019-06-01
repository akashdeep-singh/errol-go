package client

//go:generate mockgen -source client.go -destination ../mocks/mock_client.go -package mocks

type Client interface {
	CreateConnection(portNumber string) error

	Disconnect() error

	WhoAmI() error

	GetConnectedClients() error

	RelayMessage(receivers []uint64, payload string) error
}
