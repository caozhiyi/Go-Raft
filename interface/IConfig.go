package raft

type IConfig interface {

    GetCommitDiskFile() string

    GetVoteTimerRandomRange() (uint32, uint32)

    GetNodeId() uint32

    GetPort() uint16

    GetIp() string

    GetLogFile() string

    PrintLog() bool

    GetThreadNum() bool

    GetLogLevel() uint16

    GetNodeInfo() string

    GetHeartTime() uint32
}