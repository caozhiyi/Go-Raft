package raft

const (
    leader_role    = 1
    candidate_role = 2
    follower_role  = 3
)

// role interface
// leader, follower and candidate common function
type IRole interface {
    // get current role type
    GetRole() int
    // role init when role change
    ItsMyTurn()
    // get a heart message
    RecvVoteRequest(node *INode, vote_request *VoteRequest)
    // get a vote response?
    RecvVoteResponse(node *INode, vote_response *VoteResponse)
    // get a heart message
    RecvHeartBeatRequest(node *INode, heart_request *HeartBeatResquest)
    // get a heart message
    RecvHeartBeatResponse(node *INode, heart_response *HeartBeatResponse)
    // get a client request
    RecvClientRequest(client *IClient, request *ClientRequest)
    // when candidate timer out. follower and candidate
    CandidateTimeOut()
    // when heart beat timer out
    HeartBeatTimerOut()
}
