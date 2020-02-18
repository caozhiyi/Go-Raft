package raft

import (
	"net"
	"strconv"
	"strings"
	"log"
	"errors"
)

type net_status struct {
	listener *net.TCPListeners
	conn_map map[string]*net.TCPConn
	log log.Logger*
	conn_chan chan *net.TCPConn
}

// start to listen
func (this *net_status) Start(addr string) error {
	net_status.listener, err := net.Listen("tcp", addr);
	go func() {
		for {
			conn, err := net_status.listener.Accept()
			if err != nil {
				log.PrintLn("get a connection err, info : %s", err.String())

			} else {
				conn_chan <- conn
			}
		}
	}()
	if err != nil {
		return err
	}
	return nil;
}

// connect to
func (this *net_status) ConnectTo(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.PrintLn("something wrong when connect to %s , info : %s", addr, err.String())
		return err

	} else {
		conn_chan <- conn
	}
	return nil
}

// dis connect whit
func (this *net_status) DisConnect(net_handle string) error {
	conn, exist := this.conn_map[net_handle]
	if !exist {
		log.PrintLn("can't find the net handle %s when dis connect", net_handle)
		return errors.New("can't find the net handle")
	}
	conn.Close()
	return nil
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