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

	const twoSpaces string = "t "
	const eightSpaces string = "T       "
	var tableIndent string

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

		tableIndent = strings.Repeat(eightSpaces, table.depth*1)
		var metadataIndent string = "m "
// buf.WriteString(fmt.Sprintf("\t\t\t\ttable[%s].depth = %d\n", table.Name(), table.depth))
// buf.WriteString(fmt.Sprintf("\t\t\t\ttableIndent = %q\n", tableIndent))

		buf.WriteString(tableIndent + "- tableName: ")
		buf.WriteString(table.Name())
		buf.WriteByte('\n')

		buf.WriteString(tableIndent + "metadata:\n")

		for i := 0; i < table.ColCount(); i++ {
			buf.WriteString(tableIndent)
			buf.WriteString(fmt.Sprintf("%s- %s: %s\n", metadataIndent, table.colNames[i], table.colTypes[i]))
		}

		buf.WriteString(tableIndent + "data:\n")

/*
//		var valString string
//
//		for rowIndex := 0; rowIndex < len(table.rows); rowIndex++ {
//
//			for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
//
//				var dataIndent string
//				if colIndex == 0 {
//					dataIndent = tableIndent + "  - - "
//				} else {
//					dataIndent = tableIndent + "    - "
//				}
//				buf.WriteString(tableIndent)
//				buf.WriteString(dataIndent)
//
//				var colName string = table.colNames[colIndex]
//
//				switch table.colTypes[colIndex] {
//				case "string":
//					valString, err = table.GetStringByColIndex(colIndex, rowIndex)
//					if err != nil {
//						return err
//					}
//					buf.WriteString(fmt.Sprintf("%s: %q\n", colName, valString))
//				case "bool", "int", "uint", "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64", "float32", "float64":
//					valString, err = table.GetValAsStringByColIndex(colIndex, rowIndex)
//					if err != nil {
//						return err
//					}
//					buf.WriteString(fmt.Sprintf("%s: %s\n", colName, valString))
//				case "*Table":
//					buf.WriteString(fmt.Sprintf("%s:\n", colName))
//				default:
//					var colType string = table.colTypes[colIndex]
//					err = fmt.Errorf("%s: ERROR IN visitCell(): unknown type: %s\n", UtilFuncSource(), colType)
//				}
//			}
//		}
*/


		return
	}

	var visitRow = func(row Row) (err error) {
		tableIndent = strings.Repeat(eightSpaces, row.Table.depth+1)
		return
	}
	_ = visitRow

	var visitCell = func(cell Cell) (err error) {

// buf.WriteString(fmt.Sprintf("tableIndent = %q\n", tableIndent))
		var dataIndent string
		if cell.ColIndex == 0 {
			dataIndent = tableIndent + "D - - "
		} else {
			dataIndent = tableIndent + "d   - "
		}
		buf.WriteString(dataIndent)

		var valString string
		switch cell.ColType {
		case "string":
			valString, err = cell.Table.GetStringByColIndex(cell.ColIndex, cell.RowIndex)
			if err != nil {
				return err
			}
			buf.WriteString(fmt.Sprintf("%s: %q\n", cell.ColName, valString))
		case "bool", "int", "uint", "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64", "float32", "float64":
			valString, err = cell.Table.GetValAsStringByColIndex(cell.ColIndex, cell.RowIndex)
			if err != nil {
				return err
			}
			buf.WriteString(fmt.Sprintf("%s: %s\n", cell.ColName, valString))
		case "*Table":
			buf.WriteString(fmt.Sprintf("%s:\n", cell.ColName))
		default:
			err = fmt.Errorf("%s: ERROR IN visitCell(): unknown type: %s\n", UtilFuncSource(), cell.ColType)
		}

		return
	}
//	_ = visitCell

	err = tableSet.Walk(visitTableSet, visitTable, nil, visitCell)
	if err != nil {
		return "", nil
	}

	yamlString = buf.String()

	return
}
