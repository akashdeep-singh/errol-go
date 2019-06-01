package main

import (
	"fmt"
	"github.com/akashdeep-singh/errol-go/pkg/client/tcpclient"
	"github.com/akashdeep-singh/errol-go/pkg/io/networkIO"
	"github.com/akashdeep-singh/errol-go/pkg/serialization/JSONcodec"
	"github.com/akashdeep-singh/errol-go/pkg/server/tcpserver"
	"time"
)

const ServerPort = "9090"

func main() {
	server := tcpserver.TCPServer{
		RequestHandler: &tcpserver.ErrolHandler{},
		Io:             &networkIO.NetworkIO{},
		RequestCodec:   &JSONcodec.JSONRequestCodec{},
		ResponseCodec:  &JSONcodec.JSONResponseCodec{},
		IdProviderFunc: tcpserver.NewID,
	}
	go func() {
		err := server.Initialize(ServerPort)
		if err != nil {
			fmt.Println("Error initializing server")
			panic("server initialization error")
		}
	}()

	time.Sleep(2 * time.Second)

	var clients [10]tcpclient.TCPClient
	for i := 0; i < 10; i++ {
		clients[i] = tcpclient.TCPClient{
			Io:            &networkIO.NetworkIO{},
			RequestCodec:  &JSONcodec.JSONRequestCodec{},
			ResponseCodec: &JSONcodec.JSONResponseCodec{},
		}
		go func(i int) {
			err := clients[i].CreateConnection(ServerPort)
			if err != nil {
				fmt.Println("Error creating connection with server")
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("list of clients: ", server.GetConnectedClients())

	for i := range server.GetConnectedClients() {
		go func(i int) {
			err := clients[i].WhoAmI()
			if err != nil {
				fmt.Println("Error sending WhoAmI request to server")
			}
		}(i)
	}

	time.Sleep(5 * time.Second)

	go func() {
		receivers := []uint64{2, 4, 6, 8}
		err := clients[5].RelayMessage(receivers, "Test message")
		if err != nil {
			fmt.Println("Error sending RelayMessage request to server")
		}
	}()

	time.Sleep(2 * time.Second)

	go func() {
		err := server.Exit()
		if err != nil {
			fmt.Println("Error stopping server")
			panic("server stoppage error")
		}
	}()
	return
}
