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
	var indent string = ""

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

		buf.WriteString(indent + "- tableName: ")
		buf.WriteString(table.Name())
		buf.WriteByte('\n')

		indent = strings.Repeat(twoSpaces, table.depth+1)
		buf.WriteString(indent + "metadata:\n")

		return
	}

	var visitCell = func(cell Cell) (err error) {

		buf.WriteString(indent)
		buf.WriteString(fmt.Sprintf("- %s: %s\n", cell.ColName, cell.ColType))

		return
	}

	err = tableSet.Walk(visitTableSet, visitTable, visitCell)
	if err != nil {
		return "", nil
	}

	yamlString = buf.String()

	return
}
