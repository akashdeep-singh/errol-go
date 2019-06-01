package tcpserver

import (
	"errors"
	"fmt"
	"github.com/akashdeep-singh/errol/pkg/common"
	"github.com/akashdeep-singh/errol/pkg/io"
	"github.com/akashdeep-singh/errol/pkg/models"
	"github.com/akashdeep-singh/errol/pkg/serialization"
	io2 "io"
	"net"
)

type TCPServer struct {
	connections    map[uint64]net.Conn
	exit           chan bool
	listener       net.Listener
	RequestHandler RequestHandler
	Io             io.IO
	RequestCodec   serialization.RequestCodec
	ResponseCodec  serialization.ResponseCodec
	IdProviderFunc IDProviderFunc
}

func (server *TCPServer) Initialize(portNumber string) error {
	if server.listener != nil {
		fmt.Println("Server already up")
		return nil
	} else {
		fmt.Println("Server initializing")

		listener, err := net.Listen(common.Network, ":"+portNumber)
		if err != nil {
			fmt.Println("Error while starting listener:", err.Error())
			return err
		}
		defer func() {
			err := server.Exit()
			if err != nil {
				fmt.Println("Error while exiting")
			}
		}()

		server.listener = listener
		server.connections = make(map[uint64]net.Conn)
		server.exit = make(chan bool)

		fmt.Println("Server started. Listening on ", portNumber)

		for {
			connection, err := listener.Accept()

			if err != nil {
				fmt.Println("Error accepting:", err.Error())

				select {
				case <-server.exit:
					return nil
				default:
				}

				continue
			}
			id := server.IdProviderFunc()
			server.connections[id] = connection

			go func() {
				err := server.handleClientRequest(id)
				if err != nil {
					fmt.Println("Error while handling client request")
				}
			}()
		}
	}
}

func (server *TCPServer) GetConnectedClients() []uint64 {
	var list []uint64
	for id := range server.connections {
		list = append(list, id)
	}
	return list
}

func (server *TCPServer) Exit() error {
	if server.listener != nil {
		fmt.Println("Server exiting.")
		close(server.exit)

		for id, conn := range server.connections {
			if conn != nil {
				fmt.Println("Closing connection", id)
				err := conn.Close()
				if err != nil {
					return err
				}
				delete(server.connections, id)
			}
		}
		fmt.Println("Closing listener")
		err := server.listener.Close()
		if err != nil {
			return err
		} else {
			return nil
		}
	} else {
		fmt.Println("Server not running")
		return nil
	}
}

func (server *TCPServer) handleClientRequest(id uint64) error {
	for {
		request, err := server.Io.Read(server.connections[id])
		if err != nil {
			if err == io2.EOF {
				break
			}
			fmt.Println("Error reading request:", err.Error())
			return err
		}
		fmt.Println("Request from client:", string(request))

		decoded, err := server.RequestCodec.DecodeRequest(request)
		if err != nil {
			fmt.Println("Error decoding request:", err.Error())
			return err
		}

		var response models.Response
		switch decoded.RequestType() {
		case models.WhoAmIRequestType:
			response = server.RequestHandler.WhoAmI(id)
		case models.GetConnectedClientsRequestType:
			response = server.RequestHandler.GetConnectedClients(id, server.GetConnectedClients())
		case models.RelayRequestType:
			response = server.RequestHandler.Relay(decoded.(models.RelayRequest))
		default:
			err := errors.New("unsupported request type")
			return err
		}

		go func() {
			err := server.sendResponse(response)
			if err != nil {
				fmt.Println("Error while sending response")
			}
		}()
	}
	return nil
}

func (server *TCPServer) sendResponse(response models.Response) error {
	encoded := server.ResponseCodec.EncodeResponse(response)
	for _, id := range response.Targets() {
		err := server.Io.Write(server.connections[id], encoded)
		if err != nil {
			fmt.Println("Error sending response:", err.Error())
			return err
		}
	}
	return nil
}

// TODO: detect client connection close
