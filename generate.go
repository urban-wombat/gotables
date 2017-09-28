package gotables

import (
	"bytes"
	"fmt"
	"strings"
	"time"
	"unicode"
)

// The word size on the machine we're now running on.
func WordSize() int {
	return 32 << (^uint(0) >> 32 & 1)
}

// The word size on the machine the TypeStruct was generated on?
// This may not be right.
func GenerateWordSize() string {
//	var buf bytes.Buffer

	var s string = ""
	return s
}

/*
	Generate a Go struct (as a string) for storing a gotables.Table as a slice of struct.

	Compile the Go struct into your own programs.

	See also: GenerateTypeStructSet()

	See also: GenerateTypeStructSliceFromTable()

	See also: GenerateTypeStructSliceFromTableSet()

	Source table:

		[MyTable]
			  f b       i str       bb
		float32 bool  int string    []byte
			4.4 true   32 "Hello!"  [3 2 1 0]
			5.5 true  -32 "Goodie!" [4 5 6 7 8]
			6.6 false   0 "Great!"  [0 1 2]

	Generated Go struct:

		type MyTable struct {
		        f float32
		        b bool
		        i int
		        str string
		        bb []byte
		}
*/
func (table *Table) GenerateTypeStruct() (string, error) {
	if table == nil {
		return "", fmt.Errorf("table.%s() table is <nil>", funcName())
	}

	var buf bytes.Buffer
	tableName := table.Name()

	buf.WriteString("/*\n")
	buf.WriteString(fmt.Sprintf("\tAutomatically generated source code. DO NOT MODIFY. Generated %s.\n",
		time.Now().Format("15:04:05 Monday 2 Jan 2006")))
	buf.WriteString(fmt.Sprintf("\ttype %s struct generated from *gotables.Table [%s] for direct access in your code.\n",
		tableName, tableName))
	buf.WriteString("*/\n")

	buf.WriteString("type ")
	buf.WriteString(tableName)
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

/*
	Generate a set of Go structs (as a string) for storing a gotables.TableSet as slices of struct.

	Compile the Go structs into your own programs.

	See also: GenerateTypeStruct() [has sample code]

	See also: GenerateTypeStructSliceFromTable() [has sample code]

	See also: GenerateTypeStructSliceFromTableSet()

*/
func (tableSet *TableSet) GenerateTypeStructSet() (string, error) {
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
		typeStruct, err := table.GenerateTypeStruct()
		buf.WriteString(delim)
		buf.WriteString(typeStruct)
		delim = "\n"
	}

	var typeStruct string = buf.String()
	return typeStruct, nil
}

/*
	Generate a Go function (as a string) to convert a gotables.Table to a slice of struct.

	Compile the Go function into your own programs.

	See also: GenerateTypeStruct()

	See also: GenerateTypeStructSet()

	See also: GenerateTypeStructSliceFromTableSet()

	Source table:
		
		[MyTable]
		      f b       i str       bb
		float32 bool  int string    []byte
		    4.4 true   32 "Hello!"  [3 2 1 0]
		    5.5 true  -32 "Goodie!" [4 5 6 7 8]
		    6.6 false   0 "Great!"  [0 1 2]
		
	Previously-generated struct:
		
		type MyTable struct {
		        f float32
		        b bool
		        i int
		        str string
		        bb []byte
		}
		
	Generated Go function:
		
		func TypeStructSlice_MyTable_FromTable(table *gotables.Table) ([]MyTable, error) {
		        if table == nil {
		                funcName := "TypeStructSlice_MyTable_FromTable"
		                return nil, fmt.Errorf("%s(table *gotables.Table) table is <nil>", funcName)
		        }
		
		        var MyTable []MyTable = make([]MyTable, table.RowCount())
		
		        for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
		                f, err := table.GetFloat32("f", rowIndex)
		                if err != nil {
		                        return nil, err
		                }
		                MyTable[rowIndex].f = f
		
		                b, err := table.GetBool("b", rowIndex)
		                if err != nil {
		                        return nil, err
		                }
		                MyTable[rowIndex].b = b
		
		                i, err := table.GetInt("i", rowIndex)
		                if err != nil {
		                        return nil, err
		                }
		                MyTable[rowIndex].i = i
		
		                str, err := table.GetString("str", rowIndex)
		                if err != nil {
		                        return nil, err
		                }
		                MyTable[rowIndex].str = str
		
		                bb, err := table.GetByteSlice("bb", rowIndex)
		                if err != nil {
		                        return nil, err
		                }
		                MyTable[rowIndex].bb = bb
		        }
		
		        return MyTable, nil
		}

	Snippets of code as they might appear in your own program:

		type MyTable struct {
		        f float32
		        b bool
		        i int
		        str string
		        bb []byte
		}

		var a []MyTable
		var err error
		a, err = TypeStructSlice_MyTable_FromTable(table)
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(a); i++ {
			fmt.Printf("a[%d] = %v\n", i, a[i])
		}

	Output:

		a[0] = {4.4 true 32 Hello! [3 2 1 0]}
		a[1] = {5.5 true -32 Goodie! [4 5 6 7 8]}
		a[2] = {6.6 false 0 Great! [0 1 2]}
*/
func (table *Table) GenerateTypeStructSliceFromTable() (string, error) {
	if table == nil {
		return "", fmt.Errorf("table.%s() table is <nil>", funcName())
	}

	var buf bytes.Buffer

	tableName := table.Name()
	funcName := fmt.Sprintf("TypeStructSlice_%s_FromTable", tableName)
	buf.WriteString("/*\n")
	buf.WriteString(fmt.Sprintf("\tAutomatically generated source code. DO NOT MODIFY. Generated %s.\n",
		time.Now().Format("15:04:05 Monday 2 Jan 2006")))
	buf.WriteString(fmt.Sprintf("\tGenerate a slice of type %s struct from *gotables.Table [%s] for direct access in your code.\n",
		tableName, tableName))
	buf.WriteString("*/\n")
	buf.WriteString(fmt.Sprintf("func %s(table *gotables.Table) ([]%s, error) {\n", funcName, tableName))
		buf.WriteString("\tif table == nil {\n")
			buf.WriteString(fmt.Sprintf("\t\tfuncName := %q\n", funcName))
			buf.WriteString("\t\treturn nil, fmt.Errorf(\"%s(table *gotables.Table) table is <nil>\", funcName)\n")
		buf.WriteString("\t}\n\n")

		buf.WriteString(fmt.Sprintf("\tvar %s []%s = make([]%s, table.RowCount())\n\n", tableName, tableName, tableName))

		buf.WriteString("\tfor rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {\n")
			var sep string = ""
			for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
				colName, err := table.ColName(colIndex)
				if err != nil {
					return "", err
				}
				colType, err := table.ColTypeByColIndex(colIndex)
				if err != nil {
					return "", err
				}
				buf.WriteString(sep)
				buf.WriteString(fmt.Sprintf("\t\t%s, err := table.%s(%q, rowIndex)\n", colName, getterName(colType), colName))
				buf.WriteString("\t\tif err != nil {\n")
				buf.WriteString("\t\t\treturn nil, err\n")
				buf.WriteString("\t\t}\n")
				buf.WriteString(fmt.Sprintf("\t\t%s[rowIndex].%s = %s\n", tableName, colName, colName))
				sep = "\n"
			}
		
		buf.WriteString("\t}\n\n")

		buf.WriteString(fmt.Sprintf("\treturn %s, nil\n", tableName))

	buf.WriteString("}\n")

	var typeStruct string = buf.String()
	return typeStruct, nil
}


/*
	Generate a set of Go functions (as a string) to convert a gotables.TableSet to slices of struct.

	Compile the Go functions into your own programs.

	See also: GenerateTypeStruct() [has sample code]

	See also: GenerateTypeStructSet()

	See also: GenerateTypeStructSliceFromTable() [has sample code]
*/
func (tableSet *TableSet) GenerateTypeStructSliceFromTableSet() (string, error) {
	if tableSet == nil {
		return "", fmt.Errorf("tableSet.%s() tableSet is <nil>", funcName())
	}

	var tableSetGenerated string = ""
	var sep string = ""
	for tableIndex := 0; tableIndex < tableSet.TableCount(); tableIndex++ {
		tableSetGenerated += sep
		sep = "\n"
		table, err := tableSet.TableByTableIndex(tableIndex)
		if err != nil {
			return "", err
		}
		tableGenerated, err := table.GenerateTypeStructSliceFromTable()
		tableSetGenerated += tableGenerated
	}

	return tableSetGenerated, nil
}

func typeProper(typeName string) string {
	var buf bytes.Buffer
	var upshifted bool = false
	for i := 0; i < len(typeName); i++ {
		if !upshifted && unicode.IsLetter(rune(typeName[i])) {
			var typeChar []byte = make([]byte, 1)
			typeChar[0] = typeName[i]
			var upper []byte = bytes.ToUpper(typeChar)
			_ = buf.WriteByte(upper[0])
			upshifted = true
		} else {
			_ = buf.WriteByte(typeName[i])
		}
	}

	return buf.String()
}

// Returns getter/setter name without prefix Get/Set
func accessorName(typeName string) string {
	if strings.HasPrefix(typeName, "[]") {
		return fmt.Sprintf("%sSlice", typeProper(typeName[2:]))
	} else {
		return fmt.Sprintf("%s", typeProper(typeName))
	}
}

func getterName(typeName string) string {
	return fmt.Sprintf("Get%s", accessorName(typeName))
}

func setterName(typeName string) string {
	return fmt.Sprintf("Set%s", accessorName(typeName))
}

// Function to harvest a slice of struct generated by this TypeStruct function.
// Generates ascii func definition that returns a table pointer.
func (table *Table) GenerateTypeStructSliceToTable() (string, error) {
	if table == nil {
		return "", fmt.Errorf("table.%s() table is <nil>", funcName())
	}

	var buf bytes.Buffer

	tableName := table.Name()
	funcName := fmt.Sprintf("TypeStructSlice_%s_FromTable", tableName)
	buf.WriteString(fmt.Sprintf("func %s(table *gotables.Table) ([]%s, error) {\n", funcName, tableName))
		buf.WriteString("\tif table == nil {\n")
			buf.WriteString(fmt.Sprintf("\t\tfuncName := %q\n", funcName))
			buf.WriteString("\t\treturn nil, fmt.Errorf(\"%s(table *gotables.Table) table is <nil>\", funcName)\n")
		buf.WriteString("\t}\n\n")

		buf.WriteString(fmt.Sprintf("\tvar %s []%s = make([]%s, table.RowCount())\n\n", tableName, tableName, tableName))

		buf.WriteString("\tfor rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {\n")
			var sep string = ""
			for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
				colName, err := table.ColName(colIndex)
				if err != nil {
					return "", err
				}
				colType, err := table.ColTypeByColIndex(colIndex)
				if err != nil {
					return "", err
				}
				buf.WriteString(sep)
				buf.WriteString(fmt.Sprintf("\t\t%s, err := table.%s(%q, rowIndex)\n", colName, getterName(colType), colName))
				buf.WriteString("\t\tif err != nil {\n")
				buf.WriteString("\t\t\treturn nil, err\n")
				buf.WriteString("\t\t}\n")
				buf.WriteString(fmt.Sprintf("\t\t%s[rowIndex].%s = %s\n", tableName, colName, colName))
				sep = "\n"
			}
		
		buf.WriteString("\t}\n\n")

		buf.WriteString(fmt.Sprintf("\treturn %s, nil\n", tableName))

	buf.WriteString("}\n")

	var typeStruct string = buf.String()
	return typeStruct, nil
}