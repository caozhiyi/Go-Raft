package raft

const (
    vote_timer  = 0x01
    heart_timer = 0x02
)
// timer interface
type ITimer interface {
    // start timer thread
    Start()
    // reset all timer
    ResetTimer()
    // stop timer thread
    Stop()
    // add vote timer out call back
    SetVoteCallBack(fn func())
    // add heart timer out call back
    SetHeartCallBack(fn func())
    // start vote timer
    StartVoteTimer(time uint32) bool
    // heart timer
    StartHeartTimer(time uint32) bool
    StopHeartTimer()
}