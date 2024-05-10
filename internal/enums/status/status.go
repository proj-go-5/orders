package status

type Status string

const (
	NotActive Status = "not_active"
	Active    Status = "active"
	Updated   Status = "updated"
	Finished  Status = "finished"
)
