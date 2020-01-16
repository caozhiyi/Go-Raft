package raft 
// net interface,
// relize to rpc or other net
type INet interface {
    // start to listen
    Start(ip string, port uint16) error

    // connect to
    ConnectTo(ip string, port uint16) error
    // dis connect whit
    DisConnect(net_handle string) error

    // node info
    SendNodeInfoRequest(net_handle string, request *NodeInfoRequest) error
    SendNodeInfoResponse(net_handle string, response *NodeInfoResponse) error
    
    // heart beat
    SendHeartRequest(net_handle string, request *HeartBeatResquest) error
    SendHeartResponse(net_handle string, response *HeartBeatResponse) error
    // vote
    SendVoteRequest(net_handle string, request *VoteRequest) error
    SendVoteResponse(net_handle string, response *VoteResponse) error

    // client about 
    SendClientRequest(net_handle string, request *ClientRequest) error
    SendClientResponse(net_handle string, response *ClientResponse) error
    // client call back
    SetClientRecvCallBack(fn func(string, *ClientRequest)) error
    SetClientResponseCallBack(fn func(string, *ClientResponse)) error
    SetClientConnectCallBack(fn func(string)) error
    SetClientDisConnectCallBack(fn func(string)) error

    // set new connect call back
    SetNewConnectCallBack(fn func(string)) error
    // set disconnect call back
    SetDisConnectCallBack(fn func(string)) error
    // set heart request call back
    SetHeartRequestRecvCallBack(fn func(string, *HeartBeatResquest)) error
    // set heart response call back
    SetHeartResponseRecvCallBack(fn func(string, *HeartBeatResponse)) error
    // set vote request call back
    SetVoteRequestRecvCallBack(fn func(string, *VoteRequest)) error
    // set vote response call back
    SetVoteResponseRecvCallBack(fn func(string, *VoteResponse)) error
    // node info 
    SetNodeInfoRequestCallBack(fn func(string, *NodeInfoRequest)) error
    SetNodeInfoResponseCallBack(fn func(string, *NodeInfoResponse)) error
}