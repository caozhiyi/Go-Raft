package raft
// commit entries interface
// realize it to commit entries to disk or any where.
type ICommitEntries interface {
    // submit log that after merge.
    PushCompleteEntries(entries string) bool
    // submit log with index and term
    PushEntries(index uint64, term uint32, entries string) bool
    // get log vector after index
    GetEntries(index uint64, entries []string) bool
    // sync get log after index.
    GetEntriesSync(index uint64) bool
    //set get log call back
    SetEntriesCallBack(fn func([]string))
    //get the newest index
    GetNewestIndex() uint64
}