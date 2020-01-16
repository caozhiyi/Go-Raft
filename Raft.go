package raft
// raft export interface

// push entries to raft cluster
func PushEntries(string entries)

// set entries commit recv chan
func SetCommitEntriesChan(entries_recv chan<-string);