package status

type Status string

const (
	NotActive Status = "not active"
	Active    Status = "active"
	Updated   Status = "updated"
	Finished  Status = "finished"
)
