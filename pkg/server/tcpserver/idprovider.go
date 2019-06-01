package tcpserver

var lastAssignedId uint64 = 0

type IDProviderFunc func() uint64

func NewID() uint64 {
	lastAssignedId = lastAssignedId + 1
	return lastAssignedId
}
