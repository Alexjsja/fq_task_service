package isstype

type IssueType string

const (
	TASK IssueType = "TASK"
	SUB_TASK IssueType = "SUB TASK"
	BUG IssueType = "BUG"
)