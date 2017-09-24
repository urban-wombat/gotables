package gotables

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Prepare table for GOB encoding, by copying its contents to an exportable (public) table data structure.
func (table *Table) exportTable() (*TableExported, error) {
	if table == nil {
		return nil, fmt.Errorf("table.%s() table is <nil>", funcName())
	}
	var err error
	var elementCount int
	var tableExported *TableExported

	tableExported, err = newTableExported(table.Name())
	if err != nil {
		return nil, err
	}

	var colCount int = table.ColCount()

	tableExported.ColNames = make([]string, colCount)
	if len(tableExported.ColNames) != colCount {
		err = fmt.Errorf("exportTable() [%s] Could not make col names slice of size %d",
			table.Name(), colCount)
		return nil, err
	}
	elementCount = copy(tableExported.ColNames, table.colNames)
	if elementCount != colCount {
		err = fmt.Errorf("exportTable() [%s] expecting to export %d col names but exported: %d",
			table.Name(), colCount, elementCount)
		return nil, err
	}

	tableExported.ColTypes = make([]string, colCount)
	if len(tableExported.ColTypes) != colCount {
		err = fmt.Errorf("exportTable() [%s] Could not make col types slice of size %d",
			table.Name(), colCount)
		return nil, err
	}
	elementCount = copy(tableExported.ColTypes, table.colTypes)
	if elementCount != colCount {
		err = fmt.Errorf("exportTable() [%s] expecting to export %d col types but exported: %d",
			table.Name(), colCount, elementCount)
		return nil, err
	}

	tableExported.ColNamesLookup = map[string]int{}
	for key, val := range table.colNamesLookup {
		tableExported.ColNamesLookup[key] = val
	}

	var rowCount int = table.RowCount()

	tableExported.Rows = make(tableRows, rowCount)
	if len(tableExported.Rows) != rowCount {
		err = fmt.Errorf("exportTable() [%s] Could not make rows slice of size %d",
			table.Name(), rowCount)
		return nil, err
	}
	elementCount = copy(tableExported.Rows, table.rows)
	if elementCount != rowCount {
		err = fmt.Errorf("exportTable() [%s] expecting to export %d rows but exported: %d",
			table.Name(), rowCount, elementCount)
		return nil, err
	}

	tableExported.SortKeys = []SortKeyExported{}
	for keyIndex, _ := range table.sortKeys {
		tableExported.SortKeys[keyIndex] = SortKeyExported{}
		tableExported.SortKeys[keyIndex].ColName = table.sortKeys[keyIndex].colName
		tableExported.SortKeys[keyIndex].ColType = table.sortKeys[keyIndex].colType
		tableExported.SortKeys[keyIndex].Reverse = table.sortKeys[keyIndex].reverse
		tableExported.SortKeys[keyIndex].SortFunc = table.sortKeys[keyIndex].sortFunc
	}

	tableExported.StructShape = table.structShape

	return tableExported, nil
}

// Reconstitute table from GOB decoding.
func (tableExported *TableExported) importTable() (*Table, error) {
	if tableExported == nil {
		return nil, fmt.Errorf("table.%s() table is <nil>", funcName())
	}
	var err error
	var elementCount int
	var table *Table

	var tableName string = tableExported.TableName
	table, err = NewTable(tableName)
	if err != nil {
		return nil, err
	}
	var isValid bool

	var colCount int = len(tableExported.ColNames)

	table.colNames = make([]string, colCount)
	elementCount = copy(table.colNames, tableExported.ColNames)
	if elementCount != colCount {
		err = fmt.Errorf("importTable() [%s] expecting to import %d col names but imported: %d",
			tableName, colCount, elementCount)
		return nil, err
	}

	table.colTypes = make([]string, colCount)
	elementCount = copy(table.colTypes, tableExported.ColTypes)
	if elementCount != colCount {
		err = fmt.Errorf("importTable() [%s] expecting to import %d col types but imported: %d",
			tableName, colCount, elementCount)
		return nil, err
	}

	table.colNamesLookup = map[string]int{}
	for key, val := range tableExported.ColNamesLookup {
		table.colNamesLookup[key] = val
	}

	var rowCount int = len(tableExported.Rows)

	table.rows = make(tableRows, rowCount)
	elementCount = copy(table.rows, tableExported.Rows)
	if elementCount != rowCount {
		err = fmt.Errorf("importTable() [%s] expecting to import %d rows but imported: %d",
			table.Name(), rowCount, elementCount)
		return nil, err
	}

	table.sortKeys = []sortKey{}
	for keyIndex, _ := range table.sortKeys {
		table.sortKeys[keyIndex] = sortKey{}
		table.sortKeys[keyIndex].colName = tableExported.SortKeys[keyIndex].ColName
		table.sortKeys[keyIndex].colType = tableExported.SortKeys[keyIndex].ColType
		table.sortKeys[keyIndex].reverse = tableExported.SortKeys[keyIndex].Reverse
		table.sortKeys[keyIndex].sortFunc = tableExported.SortKeys[keyIndex].SortFunc
	}

	table.structShape = tableExported.StructShape

	isValid, err = table.IsValidTable()
	if !isValid {
		return nil, err
	}

	return table, nil
}

func (table *Table) GobEncode() ([]byte, error) {
	var emptyBuffer []byte

	if table == nil {
		return emptyBuffer, fmt.Errorf("table.%s() table is <nil>", funcName())
	}
	var err error
	var buffer bytes.Buffer
	var enc *gob.Encoder = gob.NewEncoder(&buffer)
	var tableExported *TableExported

	tableExported, err = table.exportTable()
	if err != nil {
		return emptyBuffer, err
	}

	err = enc.Encode(tableExported)
	if err != nil {
		return emptyBuffer, err
	}

	return buffer.Bytes(), nil
}

func (tableSet *TableSet) GobEncode() ([]bytes.Buffer, error) {
	/*
		go vet doesn't like this signature, but I see no other way. If it were an array of byte,
		there would be no way of distinguishing individual tables (arrays of byte) within it.
		$ go vet
		$ method GobEncode() ([]bytes.Buffer, error) should have signature GobEncode() ([]byte, error)
	*/
	var emptyBuffer []bytes.Buffer
	var encodedTableSet []bytes.Buffer
	var err error

	for tableIndex := 0; tableIndex < len(tableSet.tables); tableIndex++ {
		var table *Table = tableSet.tables[tableIndex]
		var encodedTable *bytes.Buffer
		var encodedTableBytes []byte
		encodedTableBytes, err = table.GobEncode()
		if err != nil {
			return emptyBuffer, err
		}
		encodedTable = bytes.NewBuffer(encodedTableBytes)
		encodedTableSet = append(encodedTableSet, *encodedTable)
		if len(encodedTableSet) != tableIndex+1 {
			err = fmt.Errorf("GobEncode(): table [%s] Error appending table to table set",
				table.Name())
			return emptyBuffer, err
		}
	}

	// Add header information to the tail end of the buffer array.
	var tableSetHeader TableSetExported
	tableSetHeader.TableSetName = tableSet.tableSetName
	tableSetHeader.FileName = tableSet.fileName
	var encodedHeader bytes.Buffer
	var enc *gob.Encoder = gob.NewEncoder(&encodedHeader)
	err = enc.Encode(tableSetHeader)
	if err != nil {
		return emptyBuffer, err
	}
	encodedTableSet = append(encodedTableSet, encodedHeader)
	var headerIndex int = len(tableSet.tables)
	if len(encodedTableSet) != headerIndex+1 {
		err = fmt.Errorf("GobEncode(): error appending table set header to table set")
		return emptyBuffer, err
	}

	return encodedTableSet, nil
}

/*
	Reconstruct a TableSet from a slice of []bytes.Buffer

	Each element in the slice is a Gob encoded table as a slice of []byte

	Calls GobDecodeTableSet(buffer)
*/
func NewTableSetFromGob(buffer []bytes.Buffer) (*TableSet, error) {
	return GobDecodeTableSet(buffer)
}

func GobDecodeTableSet(buffer []bytes.Buffer) (*TableSet, error) {
	var tableSet *TableSet
	var err error
	tableSet, err = NewTableSet("")
	if err != nil {
		return nil, err
	}

	var table *Table
	var tableCount = len(buffer) - 1 // The tail end buffer element is the TableSet header.
	for tableIndex := 0; tableIndex < tableCount; tableIndex++ {
		table, err = GobDecodeTable(buffer[tableIndex].Bytes())
		if err != nil {
			return nil, err
		}
		err = tableSet.AppendTable(table)
		if err != nil {
			return nil, err
		}
	}

	// Decode and restore the header.
	var headerIndex int = len(buffer) - 1
	var encodedHeader bytes.Buffer = buffer[headerIndex]
	var dec *gob.Decoder = gob.NewDecoder(&encodedHeader)
	var tableSetHeader TableSetExported
	err = dec.Decode(&tableSetHeader)
	if err != nil {
		return nil, err
	}
	tableSet.tableSetName = tableSetHeader.TableSetName
	tableSet.fileName = tableSetHeader.FileName

	return tableSet, nil
}

/*
	Reconstruct a Table from a slice of []byte

	Calls GobDecodeTable([]byte)
*/
func NewTableFromGob(buf []byte) (*Table, error) {
	return GobDecodeTable(buf)
}

func GobDecodeTable(buf []byte) (*Table, error) {
	var err error
	var tableDecoded *Table
	var buffer *bytes.Buffer
	buffer = bytes.NewBuffer(buf)
	var dec *gob.Decoder = gob.NewDecoder(buffer)
	var tableExported *TableExported

	err = dec.Decode(&tableExported)
	if err != nil {
		return nil, err
	}

	tableDecoded, err = tableExported.importTable()
	if err != nil {
		return nil, err
	}

	return tableDecoded, nil
}
