package drift

// DatabaseState describes the general overall state of a database
type DatabaseState struct {
	Tables map[string]TableState
}

// TableState describes the state of a table
type TableState struct {
	Name    string
	Columns map[string]ColumnState
}

type ColumnState struct {
	FQName       string
	Name         string
	InPrimaryKey bool
}

func EmptyState() DatabaseState {
	return DatabaseState{}
}
