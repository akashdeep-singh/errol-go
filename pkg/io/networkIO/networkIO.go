package networkIO

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/akashdeep-singh/errol-go/pkg/io"
	"net"
)

type NetworkIO struct {
	io.IO
}

const StreamSizeLimit = 1024000

var DataOverFlowError = errors.New("data overflow error")

func (networkIO *NetworkIO) Write(writerI interface{}, data []byte) error {
	writer := writerI.(net.Conn)
	if len(data) > StreamSizeLimit {
		return DataOverFlowError
	}
	err := binary.Write(writer, binary.LittleEndian, int32(len(data)))
	if err != nil {
		return err
	}
	_, err = writer.(net.Conn).Write(data)
	return err
}

func (networkIO *NetworkIO) Read(readerI interface{}) ([]byte, error) {
	reader := readerI.(net.Conn)
	var expectedBytes int32
	err := binary.Read(reader, binary.LittleEndian, &expectedBytes)
	if err != nil {
		return nil, err
	}

	if expectedBytes > StreamSizeLimit {
		return nil, DataOverFlowError
	}
	totalBytesRead := 0
	buffer := make([]byte, StreamSizeLimit/10)

	var data bytes.Buffer
	for {
		bytesRead, err := reader.Read(buffer)
		if err != nil {
			return nil, err
		}
		data.Write(buffer[:bytesRead])
		totalBytesRead = totalBytesRead + bytesRead

		if totalBytesRead >= int(expectedBytes) {
			break
		}
	}
	return data.Bytes(), nil
}
