package raft

type INode interface {
    // net handle
    SetNetHandle(handle *string) 
    GetNetHandle() string;

    // get send to the node next entries index
    SetNextIndex(index uint64)
    GetNextIndex() uint64
    
    // get the node newest entries index
    SetMatchIndex(index uint64)
    GetMatchIndex() uint64

    // node info
    SendNodeInfoRequest(request *NodeInfoRequest)
    SendNodeInfoResponse(response *NodeInfoResponse)

    // heart beat
    SendHeartRequest(request *HeartBeatResquest)
    SendHeartResponse(response *HeartBeatResponse)

    // vote
    SendVoteRequest(request *VoteRequest)
    SendVoteResponse(response *VoteResponse)
}