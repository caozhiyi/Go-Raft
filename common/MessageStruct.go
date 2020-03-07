package raft

const (
	HeartBeatRequest  = iota
	HeartBeatResponse
	VoteRequest
	VoteResponse
	ClientRequst
	ClientResponse
	NodeInfoRequest
	NodeInfoResponse
)

const (
	UnknowType = 01
	UsrClient  = 02
	RaftNode   = 04
)

type Header struct {
	Length uint32
	Type uint32
}
