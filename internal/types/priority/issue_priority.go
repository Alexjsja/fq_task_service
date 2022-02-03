package priority

type IssuePriority string

const (
	MINOR IssuePriority = "MINOR"
	MATTER IssuePriority = "MATTER"
	MAJOR IssuePriority = "MAJOR"
)