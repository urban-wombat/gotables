package gotables

import (
	"bytes"
	//	"encoding/json"
	"fmt"
	"strings"
)

/*
Marshall gotables TableSet to JSON
*/
func (table *Table) getTableAsJSON() (jsonString string, err error) {

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(`{ "%s":`, table.tableName))
	buf.WriteByte('[')
	for rowIndex := 0; rowIndex < len(table.rows); rowIndex++ {
		buf.WriteByte('{')
		for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
			buf.WriteByte('"')
			buf.WriteString(table.colNames[colIndex])
			buf.WriteByte('"')
			buf.WriteByte(':')
			var val interface{}
			val, err = table.GetValByColIndex(colIndex, rowIndex)
			if err != nil {
				where()
				return "", err
			}
			switch val.(type) {
			case string:
				buf.WriteString(`"` + val.(string) + `"`)
			case int, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64:
				valStr, err := table.GetValAsStringByColIndex(colIndex, rowIndex)
				if err != nil {
					where()
					return "", err
				}
				buf.WriteString(valStr)
			case bool:
				valStr, err := table.GetValAsStringByColIndex(colIndex, rowIndex)
				if err != nil {
					where()
					return "", err
				}
				buf.WriteString(valStr)
			case []byte:
				valStr, err := table.GetValAsStringByColIndex(colIndex, rowIndex)
				if err != nil {
					where()
					return "", err
				}
				// Insert comma delimiters between slice elements.
				valStr = strings.ReplaceAll(valStr, " ", ",")
				buf.WriteString(valStr)
			default:
				buf.WriteString(`"TYPE UNKNOWN"`)
			}
			if colIndex < len(table.colNames)-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte('}')
		if rowIndex < len(table.rows)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteString("]}")

	jsonString = buf.String()

	return
}

func (tableSet *TableSet) GetTableSetAsJSON() (jsonString string, err error) {

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(`{ "%s":`, tableSet.tableSetName))

	buf.WriteByte('[')
	for tableIndex := 0; tableIndex < len(tableSet.tables); tableIndex++ {

		var table *Table
		table, err = tableSet.TableByTableIndex(tableIndex)
		if err != nil {
			return "", err
		}

		var jsonTableString string
		jsonTableString, err = table.getTableAsJSON()
		if err != nil {
			return "", err
		}

		buf.WriteString(jsonTableString)

		if tableIndex < len(tableSet.tables)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')

	buf.WriteByte('}')

	jsonString = buf.String()

	return
}
