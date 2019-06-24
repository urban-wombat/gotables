package gotables

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/urban-wombat/util"
)

//
//// The word size (in bits) of the machine we're now running on. Typically 64 or 32 bits.
//func WordSize() int {
//	return 32 << (^uint(0) >> 32 & 1)
//}

/*
	Generate a Go struct (as a string) for storing a gotables.Table as a slice of struct.

	Compile the Go struct into your own programs.

	See also: GenerateTypeStructSliceFromTable()

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
		return "", fmt.Errorf("table.%s table is <nil>", util.FuncName())
	}

	var buf bytes.Buffer
	tableName := table.Name()

	buf.WriteString("/*\n")
	buf.WriteString(fmt.Sprintf("\tAutomatically generated source code. DO NOT MODIFY. Generated %s.\n\n",
		time.Now().Format("3:04 PM Monday 2 Jan 2006")))
	buf.WriteString(fmt.Sprintf("\ttype %s struct generated from *gotables.Table [%s] for including in your code.\n",
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

	UNUSED

*/
func (tableSet *TableSet) generateTypeStructSet() (string, error) {
	if tableSet == nil {
		return "", fmt.Errorf("tableSet.%s tableSet is <nil>", util.FuncName())
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

	See also: GenerateTypeStructSliceFromTableSet()

	Source table:

		[MyTable]
		      f b       i str       bb
		float32 bool  int string    []byte
		    4.4 true   32 "Hello!"  [3 2 1 0]
		    5.5 true  -32 "Goodie!" [4 5 6 7 8]
		    6.6 false   0 "Great!"  [0 1 2]

	Previously-generated type struct - using GenerateTypeStruct()

		type MyTable struct {
		        f float32
		        b bool
		        i int
		        str string
		        bb []byte
		}

	Generated Go function - using GenerateTypeStructSliceFromTable()

        Automatically generated source code. DO NOT MODIFY. Generated 1:09 PM Friday 29 Sep 2017.

        Generate a slice of type MyTable struct from *gotables.Table [MyTable] for including in your code.

		func TypeStructSliceFromTable_MyTable(table *gotables.Table) ([]MyTable, error) {
		        if table == nil {
		                return nil, fmt.Errorf("TypeStructSliceFromTable_MyTable(slice []MyTable) slice is <nil>")
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
		a, err = TypeStructSliceFromTable_MyTable(table)
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
		return "", fmt.Errorf("table.%s table is <nil>", util.FuncName())
	}

	var buf bytes.Buffer

	tableName := table.Name()
	funcName := fmt.Sprintf("TypeStructSliceFromTable_%s", tableName)

	buf.WriteString("/*\n")
	buf.WriteString(fmt.Sprintf("\tAutomatically generated source code. DO NOT MODIFY. Generated %s.\n\n",
		time.Now().Format("3:04 PM Monday 2 Jan 2006")))
	buf.WriteString(fmt.Sprintf("\tGenerate a slice of type %s struct from *gotables.Table [%s] for including in your code.\n",
		tableName, tableName))
	buf.WriteString("*/\n")

	buf.WriteString(fmt.Sprintf("func %s(table *gotables.Table) ([]%s, error) {\n", funcName, tableName))
	buf.WriteString("\tif table == nil {\n")
	buf.WriteString(fmt.Sprintf("\t\treturn nil, fmt.Errorf(\"%s(slice []%s) slice is <nil>\")\n", funcName, tableName))
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

	See also: GenerateTypeStructSliceFromTable() - convert in the opposite direction

	UNUSED
*/
func (tableSet *TableSet) generateTypeStructSliceFromTableSet() (string, error) {
	if tableSet == nil {
		return "", fmt.Errorf("tableSet.%s tableSet is <nil>", util.FuncName())
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

// Returns getter/setter name without Get/Set prefix
func accessorName(typeName string) string {
	if strings.HasPrefix(typeName, "[]") {
		return fmt.Sprintf("%sSlice", typeProper(typeName[2:]))
	} else {
		return fmt.Sprintf("%s", typeProper(typeName))
	}
}

// Return getter name for this type.
func getterName(typeName string) string {
	return fmt.Sprintf("Get%s", accessorName(typeName))
}

// Return setter name for this type.
func setterName(typeName string) string {
	return fmt.Sprintf("Set%s", accessorName(typeName))
}

/*
	Generate a Go function (as a string) to convert a slice of struct to a gotables.Table.

	Compile the Go function into your own programs.

	See also: GenerateTypeStruct()

	See also: GenerateTypeStructSliceFromTable() - convert in the opposite direction

	Source table:

		[MyTable]
		      f b       i str       bb
		float32 bool  int string    []byte
		    4.4 true   32 "Hello!"  [3 2 1 0]
		    5.5 true  -32 "Goodie!" [4 5 6 7 8]
		    6.6 false   0 "Great!"  [0 1 2]

	Previously-generated type struct - using GenerateTypeStruct()

		type MyTable struct {
		        f float32
		        b bool
		        i int
		        str string
		        bb []byte
		}

	Generated Go function - using GenerateTypeStructSliceToTable()

        Automatically generated source code. DO NOT MODIFY. Generated 1:12 PM Friday 29 Sep 2017.
        Generate a gotables Table [MyTable] from a slice of type struct []MyTable for including in your code.

		func TypeStructSliceToTable_MyTable(slice []MyTable) (*gotables.Table, error) {
		        if slice == nil {
		                return nil, fmt.Errorf("TypeStructSliceToTable_MyTable(slice []MyTable) slice is <nil>")
		        }

		        var err error

		        var seedTable string = `
		        [MyTable]
		        f float32
		        b bool
		        i int
		        str string
		        bb []byte
		        `
		        var table *gotables.Table
		        table, err = gotables.NewTableFromString(seedTable)
		        if err != nil {
		                return nil, err
		        }

		        for rowIndex := 0; rowIndex < len(slice); rowIndex++ {
		                err = table.AppendRow()
		                if err != nil {
		                        return nil, err
		                }

		                err = table.SetFloat32("f", rowIndex, slice[rowIndex].f)
		                if err != nil {
		                        return nil, err
		                }

		                err = table.SetBool("b", rowIndex, slice[rowIndex].b)
		                if err != nil {
		                        return nil, err
		                }

		                err = table.SetInt("i", rowIndex, slice[rowIndex].i)
		                if err != nil {
		                        return nil, err
		                }

		                err = table.SetString("str", rowIndex, slice[rowIndex].str)
		                if err != nil {
		                        return nil, err
		                }

		                err = table.SetByteSlice("bb", rowIndex, slice[rowIndex].bb)
		                if err != nil {
		                        return nil, err
		                }
		        }

		        return table, nil
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
		a, err = TypeStructSliceFromTable_MyTable(table)
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(a); i++ {
			fmt.Printf("a[%d] = %v\n", i, a[i])
		}

		fmt.Println()
		a[2].i = 666
		a = append(a, MyTable{f: 7.7, b: true, i: 777, str: "Wonderful!!!", bb: []byte{9, 8, 7, 6, 5} })
		var x MyTable = MyTable{f: 8.8, b: true, i: 888, str: "Wonderful!!!", bb: []byte{1, 1, 1}}
		fmt.Printf("x = %v\n", x)
		a = append(a, x)

		var outTable *gotables.Table
		outTable, err = TypeStructSliceToTable_MyTable(a)
		if err != nil {
			panic(err)
		}

		fmt.Println()
		fmt.Println(outTable)

	Output:

		a[0] = {4.4 true 32 Hello! [3 2 1 0]}
		a[1] = {5.5 true -32 Goodie! [4 5 6 7 8]}
		a[2] = {6.6 false 0 Great! [0 1 2]}

		x = {8.8 true 888 Wonderful!!! [1 1 1]}

		[MyTable]
			  f b       i str            bb
		float32 bool  int string         []byte
			4.4 true   32 "Hello!"       [3 2 1 0]
			5.5 true  -32 "Goodie!"      [4 5 6 7 8]
			6.6 false 666 "Great!"       [0 1 2]
			7.7 true  777 "Wonderful!!!" [9 8 7 6 5]
			8.8 true  888 "Wonderful!!!" [1 1 1]
*/
func (table *Table) GenerateTypeStructSliceToTable() (string, error) {
	if table == nil {
		return "", fmt.Errorf("table.%s table is <nil>", util.FuncName())
	}

	var err error
	var buf bytes.Buffer

	tableName := table.Name()
	funcName := fmt.Sprintf("TypeStructSliceToTable_%s", tableName)

	buf.WriteString("/*\n")
	buf.WriteString(fmt.Sprintf("\tAutomatically generated source code. DO NOT MODIFY. Generated %s.\n\n",
		time.Now().Format("3:04 PM Monday 2 Jan 2006")))
	buf.WriteString(fmt.Sprintf("\tGenerate a gotables Table [%s] from a slice of type struct []%s for including in your code.\n",
		tableName, tableName))
	buf.WriteString("*/\n")

	buf.WriteString(fmt.Sprintf("func %s(slice []%s) (*gotables.Table, error) {\n", funcName, tableName))
	buf.WriteString("\tif slice == nil {\n")
	buf.WriteString(fmt.Sprintf("\t\treturn nil, fmt.Errorf(\"%s(slice []%s) slice is <nil>\")\n", funcName, tableName))
	buf.WriteString("\t}\n\n")

	buf.WriteString("\tvar err error\n\n")

	colNames, colTypes, err := table.GetColInfoAsSlices()
	if err != nil {
		return "", err
	}

	tableName = table.Name()

	buf.WriteString("\tvar table *gotables.Table\n")
	buf.WriteString(fmt.Sprintf("\tvar tableName string = %q\n", tableName))
	buf.WriteString(fmt.Sprintf("\tvar colNames []string = %s\n", stringSliceToLiteral(colNames)))
	buf.WriteString(fmt.Sprintf("\tvar colTypes []string = %s\n", stringSliceToLiteral(colTypes)))

	buf.WriteString("\ttable, err = gotables.NewTableFromMetadata(tableName, colNames, colTypes)\n")
	buf.WriteString("\tif err != nil {\n")
	buf.WriteString("\t\treturn nil, err\n")
	buf.WriteString("\t}\n\n")

	buf.WriteString("\tfor rowIndex := 0; rowIndex < len(slice); rowIndex++ {\n")
	buf.WriteString("\t\terr = table.AppendRow()\n")
	buf.WriteString("\t\tif err != nil {\n")
	buf.WriteString("\t\t\treturn nil, err\n")
	buf.WriteString("\t\t}\n\n")
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
		buf.WriteString(fmt.Sprintf("\t\terr = table.%s(%q, rowIndex, slice[rowIndex].%s)\n",
			setterName(colType), colName, colName))
		buf.WriteString("\t\tif err != nil {\n")
		buf.WriteString("\t\t\treturn nil, err\n")
		buf.WriteString("\t\t}\n")
		sep = "\n"
	}

	buf.WriteString("\t}\n\n")

	buf.WriteString("\treturn table, nil\n")

	buf.WriteString("}\n")

	var typeStruct string = buf.String()
	return typeStruct, nil
}

func indentText(indent string, text string) string {
	var indentedText string = ""
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		indentedText += fmt.Sprintf("%s%s\n", indent, scanner.Text())
	}
	return indentedText
}

func stringSliceToLiteral(slice []string) string {
	var buf bytes.Buffer

	buf.WriteString("[]string{")

	var delim string = ""
	for i := 0; i < len(slice); i++ {
		buf.WriteString(delim)
		delim = ","
		buf.WriteString(fmt.Sprintf("%q", slice[i]))
	}

	buf.WriteString("}")

	return buf.String()
}
