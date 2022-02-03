package status

type IssueStatus string

const (
	TODO IssueStatus = "TODO"
	IN_WORK IssueStatus = "IN WORK"
	DONE IssueStatus = "DONE"
)