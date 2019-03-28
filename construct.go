package drift

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ConstructDatabaseState(values ...interface{}) (state DatabaseState, err error) {
	state = DatabaseState{
		Tables: make(map[string]TableState),
	}

	db, err := gorm.Open("sqlite3", ":memory:")
	for _, value := range values {
		scope := db.NewScope(value)
		modelStruct := scope.GetModelStruct()
		tableName := modelStruct.TableName(db)

		tableState := TableState{
			Name:    tableName,
			Columns: make(map[string]ColumnState),
		}

		fmt.Println(scope.TableName())
		for _, field := range modelStruct.PrimaryFields {
			fmt.Println("  + primary:", field.Name)
			tableState.Columns[field.Name] = ColumnState{
				FQName:       tableState.Name + "." + field.Name,
				Name:         field.Name,
				InPrimaryKey: true,
			}
		}
		for _, field := range modelStruct.StructFields {
			fmt.Println("  - field:", field.Name)
			tableState.Columns[field.Name] = ColumnState{
				FQName:       tableState.Name + "." + field.Name,
				Name:         field.Name,
				InPrimaryKey: false,
			}
		}

		state.Tables[tableName] = tableState
	}
	return
}
