package raft
// client manager interface
// responsible for creating client agent
type IClientManager interface {
    // set current role
    SetRole(role *IRole)
    // send response to all client agent
    SendToAll(response *ClientResponse)
    // get client numbers
    GetNodeCount() uint32
    // client call back
    RecvClientRequest(net_handle string, request* ClientRequest)
    // connection about
    ClientConnect(net_handle string)
    ClientDisConnect(net_handle string)
}