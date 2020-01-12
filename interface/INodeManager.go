package raft
// node manager interface
// responsible for creating node agent
type INodeManager interface {
    // set current role
    SetRole(role *IRole)
    // send request to all
    SendHeartToAll(request *HeartBeatResquest)
    SendVoteToAll(request *VoteRequest)

    // get node numbers
    GetNodeCount() uint32
    // get all node info
    GetAllNode() map[string]INode

    // connect to all
    ConnectToAll(net_handle_list *string)
    // connect to a node
    ConnectTo(ip *string, port uint16)

    // new connect call back
    NewConnectCallBack(net_handle *string)
    // disconnect call back
    DisConnectCallBack(net_handle *string)

    // heart request call back
    HeartRequestRecvCallBack(net_handle *string, request *HeartBeatResquest)
    // heart response call back
    HeartResponseRecvCallBack(net_handle *string, response *HeartBeatResponse)

    // vote request call back
    VoteRequestRecvCallBack(net_handle *string, request *VoteRequest)
    // vote response call back
    VoteResponseRecvCallBack(net_handle *string, response *VoteResponse)
        
    // node find about
    NodeInfoRequestCallBack(net_handle *string, request *NodeInfoRequest)
    NodeInfoResponseCallBack(net_handle *string, response *NodeInfoResponse)
}