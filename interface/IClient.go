package raft
// client agent interface
// every connected client will create one
type IClient interface {
	// net handle
	SetNetHandle(handle string) error
	GetNetHandle() (string, error)

	// send response to client
	SendToClient(response *ClientResponse)
}
