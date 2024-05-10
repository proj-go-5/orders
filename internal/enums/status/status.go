package status

type Status string

const (
	NotActive Status = "not active"
	Actived   Status = "actived"
	Updated   Status = "updated"
	Finished  Status = "finished"
)
