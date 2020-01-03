package raft

type IClientManager interface {
    SetRole(role *IRole)
    SendToAll(response *ClientResponse)
    // get client numbers
    GetNodeCount() uint32
    // client call back
    RecvClientRequest(net_handle *string, request* ClientRequest)
    ClientConnect(net_handle *string)
    ClientDisConnect(net_handle *string)
}