package raft

import (
	"net"
)

type net_status struct {

}

// start to listen
func (this *net_status) Start(ip string, port uint16) error {
	
	net.Listen("tcp", port);
}

// connect to
func (this *net_status) ConnectTo(ip string, port uint16) error {

}

// dis connect whit
func (this *net_status) DisConnect(net_handle string) error {

}

// node info
func (this *net_status) SendNodeInfoRequest(net_handle string, request *NodeInfoRequest) error {
	
}

func (this *net_status) SendNodeInfoResponse(net_handle string, response *NodeInfoResponse) error {
	
}

// heart beat
func (this *net_status) SendHeartRequest(net_handle string, request *HeartBeatResquest) error {
	
}

func (this *net_status) SendHeartResponse(net_handle string, response *HeartBeatResponse) error {
	
}

// vote
func (this *net_status) SendVoteRequest(net_handle string, request *VoteRequest) error {
	
}

func (this *net_status) SendVoteResponse(net_handle string, response *VoteResponse) error {
	
}

// client about
func (this *net_status) SendClientRequest(net_handle string, request *ClientRequest) error {
	
}

func (this *net_status) SendClientResponse(net_handle string, response *ClientResponse) error {
	
}

// client call back
func (this *net_status) SetClientRecvCallBack(fn func(string, *ClientRequest)) error {
	
}

func (this *net_status) SetClientResponseCallBack(fn func(string, *ClientResponse)) error {
	
}

func (this *net_status) SetClientConnectCallBack(fn func(string)) error {
	
}

func (this *net_status) SetClientDisConnectCallBack(fn func(string)) error {
	
}

// set new connect call back
func (this *net_status) SetNewConnectCallBack(fn func(string)) error {
	
}

// set disconnect call back
func (this *net_status) SetDisConnectCallBack(fn func(string)) error {
	
}

// set heart request call back
func (this *net_status) SetHeartRequestRecvCallBack(fn func(string, *HeartBeatResquest)) error {
	
}

// set heart response call back
func (this *net_status) SetHeartResponseRecvCallBack(fn func(string, *HeartBeatResponse)) error {
	
}

// set vote request call back
func (this *net_status) SetVoteRequestRecvCallBack(fn func(string, *VoteRequest)) error {
	
}

// set vote response call back
func (this *net_status) SetVoteResponseRecvCallBack(fn func(string, *VoteResponse)) error {
	
}

// node info 
func (this *net_status) SetNodeInfoRequestCallBack(fn func(string, *NodeInfoRequest)) error {
	
}

func (this *net_status) SetNodeInfoResponseCallBack(fn func(string, *NodeInfoResponse)) error {
	
}