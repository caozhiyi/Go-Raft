package raft

// heart beat request struct
type HeartBeatResquest struct {
	Term uint32                  // leader's term.
	Leader_id uint32             // leader's id, so followers can redirect clients.
	Prev_log_index uint64        // the followers replaces all entries no through and including this index.
	Prev_log_term  uint32        // term of last include index.
	Entries []string             // to add logs.
	Leader_commit uint64         // leader added log index.
}

// heart beat response struct
type HeartBeatResponse struct {
	Term uint32                  // now leader's term.
	Success bool                 // add success?
	Next_index uint64 		     // the node next index
}

// vote response struct
type VoteRequest struct {
    Term uint32                 // candidate's term.
	Candidate_id uint32         // candidate's id.
	Last_index uint64           // candidate last log index.
	Last_term uint32            // candidate last log term.
}

type VoteResponse struct {
	Term uint32                 // now leader's term.
	Vote_granted bool           // is vote?
}

type ClientRequest struct {
	Entries string              // commit entries.
}

type NodeInfoRequest struct {
	net_handle []string        // all node net address
}

type NodeInfoResponse struct {
	net_handle []string        // all node net address
}


const (
	success                    = 0
	not_leader                 = 1
	other_error                = 2
	send_again                 = 3    // now don't have a leader, client try again a minute after
)

type ClientResponse struct {
	Ret_code uint16           // return code
	Leader_net_handle string  // leader net handle
}