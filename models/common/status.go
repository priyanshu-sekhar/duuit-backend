package common

type Status int
const (
	COMPLETED Status = iota
	SKIPPED
	ABORTED
)
