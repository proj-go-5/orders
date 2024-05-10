package status

type Status string

const (
	NotActive   Status = "not active"
	Added       Status = "added"
	Updated     Status = "updated"
	Deactivated Status = "deactivated"
)
