package raft

type IClient interface {
    // net handle
    SetNetHandle(handle *string)
    GetNetHandle() string

    // send response to client
    SendToClient(response *ClientResponse)
}