package gotables

import (
	"bytes"
	"fmt"
)

// The word size on the machine we're now running on.
func WordSize() int {
	return 32 << (^uint(0) >> 32 & 1)
}

// The word size on the machine the TypeStruct was generated on?
// This may not be right.
func TypeStructWordSize() int {
	return 0
}

func (table *Table) TypeStruct() (string, error) {
	if table == nil {
		return "", fmt.Errorf("table.%s() table is <nil>", funcName())
	}

	var buf bytes.Buffer

	buf.WriteString("type ")
	buf.WriteString(table.Name())
	buf.WriteString(" struct {\n")
	for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
		colName, err := table.ColName(colIndex)
		if err != nil {
			return "", err
		}
		colType, err := table.ColTypeByColIndex(colIndex)
		if err != nil {
			return "", err
		}
		buf.WriteString("\t")
		buf.WriteString(colName)
		buf.WriteString(" ")
		buf.WriteString(colType)
		buf.WriteString("\n")
	}

	buf.WriteString("}\n")

	var typeStruct string = buf.String()
	return typeStruct, nil
}

func (tableSet *TableSet) TypeStructSet() (string, error) {
	if tableSet == nil {
		return "", fmt.Errorf("tableSet.%s() tableSet is <nil>", funcName())
	}
	
	var buf bytes.Buffer
	var delim string = ""

	for tableIndex := 0; tableIndex < tableSet.TableCount(); tableIndex++ {
		table, err := tableSet.TableByTableIndex(tableIndex)
		if err != nil {
			return "", err
		}
		typeStruct, err := table.TypeStruct()
		buf.WriteString(delim)
		buf.WriteString(typeStruct)
		delim = "\n"
	}

	var typeStruct string = buf.String()
	return typeStruct, nil
}
