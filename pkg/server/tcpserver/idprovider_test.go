package tcpserver_test

import (
	"github.com/akashdeep-singh/errol-go/pkg/server/tcpserver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	result := tcpserver.NewID()
	assert.Equal(t, uint64(1), result)

	result = tcpserver.NewID()
	assert.Equal(t, uint64(2), result)

	result = tcpserver.NewID()
	assert.Equal(t, uint64(3), result)

	result = tcpserver.NewID()
	assert.Equal(t, uint64(4), result)

	result = tcpserver.NewID()
	assert.Equal(t, uint64(5), result)

	tcpserver.NewID()
	tcpserver.NewID()
	tcpserver.NewID()
	tcpserver.NewID()
	result = tcpserver.NewID()
	assert.Equal(t, uint64(10), result)
}
