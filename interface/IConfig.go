package raft
// config interface.
// realize it to get config from disk or any where.
type IConfig interface {
    // get commit disk file name
    GetCommitDiskFile() string
    // get candidate time out
    GetVoteTimerRandomRange() (uint32, uint32)
    // get current node id, must be unique
    GetNodeId() uint32
    // get listen port
    GetPort() uint16
    // get listen ip address
    GetIp() string
    // get log file name if print log
    GetLogFile() string
    // need print log?
    PrintLog() bool
    // get raft process thread
    GetThreadNum() bool
    // get log level
    GetLogLevel() uint16
    // get node address info list
    GetNodeInfo() string
    // get heart time out
    GetHeartTime() uint32
}