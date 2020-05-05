package gotables

import (
	"bytes"
	"fmt"
	"strings"
)

func (tableSet *TableSet) GetTableSetAsYAML() (yamlString string, err error) {

	if tableSet == nil {
		return "", fmt.Errorf("%s tableSet.%s: table set is <nil>", UtilFuncSource(), UtilFuncName())
	}

	const twoSpaces string = "  "	// Two spaces

	var buf bytes.Buffer

	buf.WriteString("---\n")	// Start of YAML document

	var visitTableSet = func(tableSet *TableSet) (err error) {
		buf.WriteString(`tableSetName: "`)
		buf.WriteString(tableSet.Name())
		buf.WriteByte('"')
		buf.WriteByte('\n')
		buf.WriteString("tables:\n")
		return
	}

	var visitTable = func(table *Table) (err error) {
		var indent string = strings.Repeat(twoSpaces, table.depth)
		buf.WriteString(indent + "- tableName: ")
		buf.WriteString(table.Name())
		buf.WriteByte('\n')

		indent = strings.Repeat(twoSpaces, table.depth+1)
		buf.WriteString(indent + "metadata:\n")
		for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
			buf.WriteString(indent)
			buf.WriteString(fmt.Sprintf("- %s: %s\n", table.colNames[colIndex], table.colTypes[colIndex]))
/*
			buf.WriteString(table.colNames[colIndex])
			buf.WriteString(`":"`)
			buf.WriteString(table.colTypes[colIndex])
			buf.WriteByte('"')
			if colIndex < len(table.colNames)-1 {
				buf.WriteByte(',')
			}
*/
		}
		indent = strings.Repeat(twoSpaces, table.depth)

		return
	}

	err = tableSet.Walk(visitTableSet, visitTable, nil)
	if err != nil {
		return "", nil
	}

	yamlString = buf.String()

	return
}
