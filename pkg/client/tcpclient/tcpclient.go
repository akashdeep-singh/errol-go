package tcpclient

import (
	"fmt"
	"github.com/akashdeep-singh/errol/pkg/common"
	"github.com/akashdeep-singh/errol/pkg/io"
	"github.com/akashdeep-singh/errol/pkg/models"
	"github.com/akashdeep-singh/errol/pkg/serialization"
	"net"
)

type TCPClient struct {
	connection    net.Conn
	Io            io.IO
	RequestCodec  serialization.RequestCodec
	ResponseCodec serialization.ResponseCodec
}

func (client *TCPClient) Disconnect() error {
	return client.connection.Close()
}

func (client *TCPClient) CreateConnection(portNumber string) error {
	fmt.Println("Connecting to server at port: ", portNumber)
	connection, err := net.Dial(common.Network, ":"+portNumber)
	if err != nil {
		fmt.Println("Error while creating connection:", err.Error())
		return err
	}
	client.connection = connection

	go func() {
		for {
			message, _ := client.readServerMessage()
			if message != nil {
				switch message.ResponseType() {
				case models.WhoAmIResponseType:
					fmt.Println("Received whoami response: ", message.(models.WhoAmIResponse).Id)
				case models.GetConnectedClientsResponseType:
					fmt.Println("Received whoami response: ", message.(models.GetConnectedClientsResponse).ConnectedClients)
				case models.RelayResponseType:
					fmt.Println("Received relay message: ",
						string(message.(models.RelayResponse).Message[:]))
				}
			}
		}
	}()
	return nil
}

func (client *TCPClient) WhoAmI() error {
	request := models.NewWhoAmIRequest()
	err := client.sendRequest(request)
	if err != nil {
		fmt.Println("Error sending request to server")
	}
	return err
}

func (client *TCPClient) GetConnectedClients() error {
	request := models.NewGetConnectedClientsRequest()
	err := client.sendRequest(request)
	if err != nil {
		fmt.Println("Error sending request to server")
	}
	return err
}

func (client *TCPClient) RelayMessage(receivers []uint64, payload string) error {
	request := models.NewRelayRequest(receivers, payload)
	err := client.sendRequest(request)
	if err != nil {
		fmt.Println("Error sending request to server")
	}
	return err
}

func (client *TCPClient) sendRequest(request models.Request) error {
	encoded := client.RequestCodec.EncodeRequest(request)
	err := client.Io.Write(client.connection, encoded)
	return err
}

func (client *TCPClient) readServerMessage() (models.Response, error) {
	message, err := client.Io.Read(client.connection)
	if err != nil {
		fmt.Println("Error receiving server message:", err.Error())
		return nil, err
	}
	fmt.Println("Received server message:", string(message))

	decoded, err := client.ResponseCodec.DecodeResponse(message)
	if err != nil {
		fmt.Println("Error decoding server message:", err.Error())
		return nil, err
	}
	return decoded.(models.Response), nil
}

// TODO: implement request IDs and request queueing so that the requesting methods can return the result
