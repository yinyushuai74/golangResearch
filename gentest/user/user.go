package user

//go:generate gentest -type Status
type Status int
const (
	Offline Status = iota
	Online
	Disable
	Deleted
)
type Color int
const (
	Write Color = iota
	Red
	Blue
)