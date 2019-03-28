package drift

import "github.com/nsogame/drift/ddl"

// Diff describes the change set between two database states
type DatabaseDiff struct {
	Changes []interface{}
}

type TableDiff struct {
	Changes []interface{}
}

func CalculateDatabaseDiff(left, right DatabaseState) (diff DatabaseDiff) {
	diff = DatabaseDiff{}
	intersection := make(map[string]TableState)

	// iterate over old tables
	for key, value := range left.Tables {
		_, ok := right.Tables[key]
		if ok {
			// this table is still in the new version
			intersection[key] = value
		} else {
			// this table got dropped!
			diff.Changes = append(diff.Changes, ddl.DropTable{key})
		}
	}

	// iterate over new tables
	for key, _ := range right.Tables {
		_, ok := intersection[key]
		if ok {
			// calculate diff of changes
		} else {
			// this table is new!
			// TODO: column info
			diff.Changes = append(diff.Changes, ddl.CreateTable{key})
		}
	}
	return
}

func CalculateTableDiff(left, right TableState) (diff TableDiff) {
	diff = TableDiff{}
	return
}
