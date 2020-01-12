package raft 
// net interface,
// relize to rpc or other net
type INet interface {
    // start to listen
    Init(thread_num uint16)
    // start to listen
    Start(ip string, port uint16) bool

    // connect to
    ConnectTo(ip string, port uint16)
    // dis connect whit
    DisConnect(net_handle string)

    // node info
    SendNodeInfoRequest(net_handle string, request *NodeInfoRequest)
    SendNodeInfoResponse(net_handle string, response *NodeInfoResponse)
    
    // heart beat
    SendHeartRequest(net_handle string, request *HeartBeatResquest)
    SendHeartResponse(net_handle string, response *HeartBeatResponse)
    // vote
    SendVoteRequest(net_handle string, request *VoteRequest)
    SendVoteResponse(net_handle string, response *VoteResponse)

    // client about
    SendClientRequest(net_handle string, request *ClientRequest)
    SendClientResponse(net_handle string, response *ClientResponse)
    // client call back
    SetClientRecvCallBack(fn func(string, *ClientRequest))
    SetClientResponseCallBack(fn func(string, *ClientResponse))
    SetClientConnectCallBack(fn func(string))
    SetClientDisConnectCallBack(fn func(string))

    // set new connect call back
    SetNewConnectCallBack(fn func(string))
    // set disconnect call back
    SetDisConnectCallBack(fn func(string))
    // set heart request call back
    SetHeartRequestRecvCallBack(fn func(string, *HeartBeatResquest))
    // set heart response call back
    SetHeartResponseRecvCallBack(fn func(string, *HeartBeatResponse))
    // set vote request call back
    SetVoteRequestRecvCallBack(fn func(string, *VoteRequest))
    // set vote response call back
    SetVoteResponseRecvCallBack(fn func(string, *VoteResponse))
    // node info 
    SetNodeInfoRequestCallBack(fn func(string, *NodeInfoRequest))
    SetNodeInfoResponseCallBack(fn func(string, *NodeInfoResponse))
}