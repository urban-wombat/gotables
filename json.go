package gotables

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"strings"

	"github.com/urban-wombat/util"
)

var replaceSpaces *regexp.Regexp = regexp.MustCompile(` `)

/*
	Marshal json from the rows of data in this table.

	A *gotables.Table is composed of metadata and data:-
		1. Metadata:-
			* Table name
			* Column names
			* Column types
		2. Data:
			* Rows of data

	To generate json metadata and data:-
		1. Meta: call method table.GetTableMetadataAsJSON()
		2. Data: call method table.GetTableDataAsJSON()
*/
func (table *Table) GetTableDataAsJSON() (jsonDataString string, err error) {

	if table == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(`{"%s":`, table.tableName))
	buf.WriteByte('[')
	for rowIndex := 0; rowIndex < len(table.rows); rowIndex++ {
		buf.WriteByte(123) // Opening brace
		for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
			buf.WriteByte('"')
			buf.WriteString(table.colNames[colIndex])
			buf.WriteString(`":`)
			var val interface{}
			val, err = table.GetValByColIndex(colIndex, rowIndex)
			if err != nil {
				return "", err
			}
			switch val.(type) {
			case string:
				buf.WriteString(`"` + val.(string) + `"`)	// May be faster as 3 writes?
			case bool, int, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64:
				valStr, err := table.GetValAsStringByColIndex(colIndex, rowIndex)
				if err != nil {
					return "", err
				}
				buf.WriteString(valStr)
			case []byte:
				valStr, err := table.GetValAsStringByColIndex(colIndex, rowIndex)
				if err != nil {
					return "", err
				}
				// Insert comma delimiters between slice elements.
				//				valStr = strings.ReplaceAll(valStr, " ", ",")	// New in Go 1.11?
				valStr = replaceSpaces.ReplaceAllString(valStr, ",")
				buf.WriteString(valStr)
			default:
				buf.WriteString(`"TYPE UNKNOWN"`)
			}
			if colIndex < len(table.colNames)-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte(125) // Closing brace
		if rowIndex < len(table.rows)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteString("]}")

	jsonDataString = buf.String()

	return
}

/*
	Marshal json from the metadata in this table.

	A *gotables.Table is composed of metadata and data:-
		1. Metadata:-
			* Table name
			* Column names
			* Column types
		2. Data:
			* Rows of data

	To generate json metadata and data:-
		1. Meta: call method table.GetTableMetadataAsJSON()
		2. Data: call method table.GetTableDataAsJSON()

	Note: The table must have at least 1 col defined (zero rows are okay).
*/
func (table *Table) GetTableMetadataAsJSON() (jsonMetadataString string, err error) {

	if table == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	if table.ColCount() == 0 {
		return "", fmt.Errorf("in table [%s]: cannot marshal json metadata from a table with zero columns", table.Name())
	}

	var buf bytes.Buffer

	buf.WriteByte(123) // Opening brace
	buf.WriteString(fmt.Sprintf(`"%s":`, table.tableName))
	buf.WriteByte('[')
	for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
		buf.WriteByte(123) // Opening brace
		buf.WriteByte('"')
		buf.WriteString(table.colNames[colIndex])
		buf.WriteString(`":"`)
		buf.WriteString(table.colTypes[colIndex])
		buf.WriteByte('"')
		buf.WriteByte(125) // Closing brace
		if colIndex < len(table.colNames)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')
	buf.WriteByte(125) // Closing brace

	jsonMetadataString = buf.String()

	return
}

/*
	Marshal gotables TableSet to JSON

	The TableSet is returned as two parallel slices of JSON:-
		1. A slice of metadata strings: tableName, colNames and colTypes
		2. A slice of data strings: rows of data corresponding to the metadata

	Each slice element of metadata corresponds with (matches) each element of row data.
*/
func (tableSet *TableSet) GetTableSetAsJSON() (jsonMetadataStrings []string, jsonDataStrings []string, err error) {

	if tableSet == nil {
		return nil, nil, fmt.Errorf("%s %s tableSet is <nil>", util.FuncSource(), util.FuncName())
	}

	for tableIndex := 0; tableIndex < len(tableSet.tables); tableIndex++ {

		var table *Table
		table, err = tableSet.TableByTableIndex(tableIndex)
		if err != nil {
			return nil, nil, err
		}

		var jsonMetadataString string
		jsonMetadataString, err = table.GetTableMetadataAsJSON()
		if err != nil {
			return nil, nil, err
		}
		jsonMetadataStrings = append(jsonMetadataStrings, jsonMetadataString)

		var jsonDataString string
		jsonDataString, err = table.GetTableDataAsJSON()
		if err != nil {
			return nil, nil, err
		}
		jsonDataStrings = append(jsonDataStrings, jsonDataString)
	}

	return
}

/*
	Unmarshal a document of JSON metadata and a document of JSON data to a *gotables.Table

	Two JSON documents are required:-
		1. JSON metadata which contains the tableName, colNames and colTypes.
		2. JSON data which contains zero or more rows of data that map to the metadata.

	The two documents must match: the metadata must match the corresponding data.
*/
func NewTableFromJSON(jsonMetadataString string, jsonDataString string) (table *Table, err error) {

	if jsonMetadataString == "" {
		return nil, fmt.Errorf("newTableFromJSON(): jsonMetadataString is empty")
	}

	if jsonDataString == "" {
		return nil, fmt.Errorf("newTableFromJSON(): jsonDataString is empty")
	}

	// Create empty table from metadata.
	/* Note: To preserve column order, we do NOT use JSON marshalling into a map,
	   because iterating over a map returns values in random order.
	   Instead, we use the json decoder. (The data rows (later in this function)
	   ARE decoded using a map.)
	   Actually, the jsonMetadataString is an array, so it probably WOULD work.
	   TODO: Use a map to decode jsonMetadataString
	*/

	dec := json.NewDecoder(strings.NewReader(jsonMetadataString))
	var token json.Token

	// Skip opening brace
	token, err = dec.Token()
	if err == io.EOF {
		return nil, fmt.Errorf("newTableFromJSON(): unexpected EOF")
	}
	if err != nil {
		return nil, fmt.Errorf("newTableFromJSON(): %v", err)
	}

	// Get table name
	token, err = dec.Token()
	if err == io.EOF {
		return nil, fmt.Errorf("newTableFromJSON(): unexpected EOF")
	}
	if err != nil {
		return nil, fmt.Errorf("newTableFromJSON(): %v", err)
	}

	// Get the table name.
	var metadataTableName string
	switch token.(type) {
	case string: // As expected
		metadataTableName = token.(string)
		table, err = NewTable(metadataTableName)
		if err != nil {
			return nil, fmt.Errorf("newTableFromJSON(): %v", err)
		}
	default:
		return nil, fmt.Errorf("newTableFromJSON(): expecting table name but found: %v", reflect.TypeOf(token))
	}

	// Simple parsing flags and values.
	var colNameNext bool = false
	var colName string
	var colTypeNext bool = false
	var colType string
	var prevDelim rune

Loop:
	for {
		token, err = dec.Token()
		if err == io.EOF {
			return nil, fmt.Errorf("newTableFromJSON(): unexpected EOF")
		}
		if err != nil {
			return nil, fmt.Errorf("newTableFromJSON(): %v", err)
		}

		switch token.(type) {
		case json.Delim:
			delim := token.(json.Delim)
			switch delim {
			case 123: // Opening brace
				colNameNext = true
				prevDelim = 123 // Opening brace
			case 125: // Closing brace
				if prevDelim == 125 { // Closing brace: end of JSON metadata object
					// Table metadata is now completely initialised. Now do the rows of data.
					//							return table, nil
					break Loop
				}
				// We now have a colName-plus-colType pair. Add this col to table.
				err = table.AppendCol(colName, colType)
				if err != nil {
					return nil, fmt.Errorf("newTableFromJSON(): %v", err)
				}
				prevDelim = 125 // Closing brace: end of col
			case '[': // Ignore slice signifiers in type names
			case ']': // Ignore slice signifiers in type names
			}
		case string:
			if colNameNext {
				colName = token.(string)
				colNameNext = false
				colTypeNext = true
			} else if colTypeNext {
				colType = token.(string)
				colTypeNext = false
			} else {
				return nil, fmt.Errorf("newTableFromJSON(): expecting colName or colType")
			}
		case bool:
			return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
		case float64:
			return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
		case json.Number:
			return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
		case nil:
			return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
		default:
			fmt.Printf("unknown json token type %T value %v\n", token, token)
		}
	}


	// Append row of table data from JSON.
	/*
	   Note: Here we use a map for rows of data now that we have already preserved col order.
	   Unmarshal does all the parsing for us.
	*/

	var unmarshalled interface{}
	err = json.Unmarshal([]byte(jsonDataString), &unmarshalled)
	if err != nil {
		return nil, fmt.Errorf("newTableFromJSON(): %v", err)
	}

	var tableMap map[string]interface{} = unmarshalled.(map[string]interface{})

	// Check that this JSON data (rows) object table name matches the JSON metadata object table name.
	// (Could have simply used metadataTableName as the key to a lookup.)
	var dataTableName string
	for dataTableName, _ = range tableMap {
		// There should be only one key, and it should be the table name.
	}
	if dataTableName != metadataTableName {
		return nil, fmt.Errorf("newTableFromJSON(): unexpected JSON metadataTableName %q != JSON dataTableName %q",
			metadataTableName, dataTableName)
	}

	var rowsInterface []interface{} = tableMap[dataTableName].([]interface{})
//where(rowsInterface)

	for rowIndex, row := range rowsInterface {
		table.AppendRow()
		var rowMap map[string]interface{} = row.(map[string]interface{})
		for colName, val := range rowMap {
			var colIndex = table.colNamesMap[colName]
			var colType string = table.colTypes[colIndex]
			switch val.(type) {
			case string:
				err = table.SetString(colName, rowIndex, val.(string))
			case float64:
				switch colType {
				case "int":
					err = table.SetInt(colName, rowIndex, int(val.(float64)))
				case "uint":
					err = table.SetUint(colName, rowIndex, uint(val.(float64)))
				case "byte":
					err = table.SetByte(colName, rowIndex, byte(val.(float64)))
				case "int8":
					err = table.SetInt8(colName, rowIndex, int8(val.(float64)))
				case "int16":
					err = table.SetInt16(colName, rowIndex, int16(val.(float64)))
				case "int32":
					err = table.SetInt32(colName, rowIndex, int32(val.(float64)))
				case "int64":
					err = table.SetInt64(colName, rowIndex, int64(val.(float64)))
				case "uint8":
					err = table.SetUint8(colName, rowIndex, uint8(val.(float64)))
				case "uint16":
					err = table.SetUint16(colName, rowIndex, uint16(val.(float64)))
				case "uint32":
					err = table.SetUint32(colName, rowIndex, uint32(val.(float64)))
				case "uint64":
					err = table.SetUint64(colName, rowIndex, uint64(val.(float64)))
				case "float32":
					err = table.SetFloat32(colName, rowIndex, float32(val.(float64)))
				case "float64":
					err = table.SetFloat64(colName, rowIndex, float64(val.(float64)))
				}
			case bool:
				err = table.SetBool(colName, rowIndex, val.(bool))
			case []interface{}: // This cell is a slice
				var interfaceSlice []interface{} = val.([]interface{})
				var byteSlice []byte = []byte{}
				for _, sliceVal := range interfaceSlice {
					byteSlice = append(byteSlice, byte(sliceVal.(float64)))
				}
				err = table.SetByteSlice(colName, rowIndex, byteSlice)
			case nil:
				return nil, fmt.Errorf("newTableFromJSON(): unexpected nil value")
			default:
				return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(val))
			}

			// Single error handler for all the table.Set...() calls.
			if err != nil {
				return nil, fmt.Errorf("newTableFromJSON(): %v", err)
			}
		}
	}

	return table, nil
}

/*
	Unmarshal a slice of JSON metadata documents and a slice of JSON data documents to a *gotables.TableSet

	Two JSON documents are required:-
		1. A slice of JSON metadata which contains the tableName, colNames and colTypes.
		2. A slice of JSON data which contains zero or more rows of data that map to the metadata.

	The two slices must be parallel: each element of metadata must match the corresponding element of data.
*/
func NewTableSetFromJSON(jsonMetadataStrings []string, jsonDataStrings []string) (tableSet *TableSet, err error) {

	if jsonMetadataStrings == nil {
		return nil, fmt.Errorf("jsonMetadataStrings == nil")
	}

	if jsonDataStrings == nil {
		return nil, fmt.Errorf("jsonDataStrings == nil")
	}

	if len(jsonMetadataStrings) != len(jsonDataStrings) {
		return nil, fmt.Errorf("len(jsonMetadataStrings) %d != len(jsonDataStrings) %d",
			len(jsonMetadataStrings), len(jsonDataStrings))
	}

	tableSet, err = NewTableSet("")
	if err != nil {
		return nil, err
	}

	for tableIndex := 0; tableIndex < len(jsonMetadataStrings); tableIndex++ {
		table, err := NewTableFromJSON(jsonMetadataStrings[tableIndex], jsonDataStrings[tableIndex])
		if err != nil {
			return nil, err
		}

		err = tableSet.AppendTable(table)
		if err != nil {
			return nil, err
		}
	}

	return
}
