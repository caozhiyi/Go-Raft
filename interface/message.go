package raft;
// communication protocol struct
type HeartBeatResquest struct {
	term uint32
	leader_id uint32
	prev_log_index uint64
	prev_log_term uint64
	entries []string
	leader_commit uint64
}

type HeartBeatResponse struct {
	term uint32
	success bool
	next_index uint64
}

type VoteRequest struct {
	term uint32
	candidate_id uint32
	last_index uint64
	last_term uint32
}

type VoteResponse struct {
	term uint32
	vote_granted uint32
}

type ClientRequest struct {
	entries string
}

const (
	success                    = 0;
	not_leader                 = 1;
	other_error                = 2;
	send_again                 = 3;    // now don't have a leader, client try again a minute after
)

type ClientResponse struct {
	ret_code uint16
	leader_net_handle string
}

type NodeInfoRequest struct {
	net_handle []string
}

type NodeInfoResponse struct {
	net_handle []string
}
