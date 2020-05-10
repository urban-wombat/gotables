package gotables_test

import (
	_ "fmt"
	"testing"

	"github.com/urban-wombat/gotables"
)

/*
Copyright (c) 2018 Malcolm Gorman

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

//func TestTableSet_GetTableSetAsYAML(t *testing.T) {
//
//	var err error
//
//	// Cheat and get our test TableSet from JSON.
//	var jsonString string = `
//	{"tableSetName":"MyTableSetName","tables":[{"tableName":"TypesGalore22","metadata":[{"i":"int"},{"s":"string"},{"right":"*Table"}],"data":[[{"i":0},{"s":"abc"},{"right":{"tableName":"right0","isStructShape":true,"metadata":[{"i":"int"}],"data":[[{"i":32}]]}}],[{"i":1},{"s":"xyz"},{"right":{"tableName":"right1","isStructShape":true,"metadata":[{"s":"string"}],"data":[[{"s":"thirty-two"}]]}}],[{"i":2},{"s":"ssss"},{"right":{"tableName":"right2","metadata":[{"x":"int"},{"y":"int"},{"z":"int"}],"data":[[{"x":1},{"y":2},{"z":3}],[{"x":4},{"y":5},{"z":6}],[{"x":7},{"y":8},{"z":9}]]}}],[{"i":3},{"s":"xxxx"},{"right":{"tableName":"right3","isStructShape":true,"metadata":[{"f":"float32"}],"data":[[{"f":88.8}]]}}],[{"i":4},{"s":"yyyy"},{"right":{"tableName":"right4","isStructShape":true,"metadata":[{"t1":"*Table"}],"data":[[{"t1":null}]]}}]]}]}
//`
//	var tableSet *gotables.TableSet
//	tableSet, err = gotables.NewTableSetFromJSON(jsonString)
//	if err != nil {
//		t.Fatal(err)
//	}
//	// fmt.Println(tableSet.String())
//
//	var yaml string
//	yaml, err = tableSet.GetTableSetAsYAML()
//	if err != nil {
//		t.Fatal(err)
//	}
//	// fmt.Println(yaml)
//
//	expected := `---
//tableSetName: "MyTableSetName"
//tables:
//- tableName: TypesGalore22
//  metadata:
//  - i: int
//  - s: string
//  - right: "*Table"
//  data:
//  - - i: 0
//    - s: "abc"
//    - right:
//        - tableName: right0
//          isStructShape: true
//          metadata:
//          - i: int
//          data:
//          - - i: 32
//          - - i: 1
//            - s: "xyz"
//            - right:
//        - tableName: right1
//          isStructShape: true
//          metadata:
//          - s: string
//          data:
//          - - s: "thirty-two"
//          - - i: 2
//            - s: "ssss"
//            - right:
//        - tableName: right2
//          metadata:
//          - x: int
//          - y: int
//          - z: int
//          data:
//          - - x: 1
//            - y: 2
//            - z: 3
//          - - x: 4
//            - y: 5
//            - z: 6
//          - - x: 7
//            - y: 8
//            - z: 9
//          - - i: 3
//            - s: "xxxx"
//            - right:
//        - tableName: right3
//          isStructShape: true
//          metadata:
//          - f: float32
//          data:
//          - - f: 88.8
//          - - i: 4
//            - s: "yyyy"
//            - right:
//        - tableName: right4
//          isStructShape: true
//          metadata:
//          - t1: "*Table"
//          data:
//          - - t1:
//        - tableName: 
//          metadata:
//          data:
//`
//
//	if yaml != expected {
//		t.Errorf("yaml (long document) does not match expected yaml (long document)")
//	}
//}

/*
//func Test_NewTableSetFromYAML(t *testing.T) {
//
//	var tableSet *gotables.TableSet
//	var err error
//
//	var yaml string = `---
//tableSetName: "MyTableSetName"
//tables:
//- tableName: TypesGalore22
//  metadata:
//  - i: int
//  - s: string
//  - right: "*Table"
//  data:
//  - - i: 0
//    - s: "abc"
//    - right:
//        - tableName: right0
//          isStructShape: true
//          metadata:
//          - i: int
//          data:
//          - - i: 32
//          - - i: 1
//            - s: "xyz"
//            - right:
//        - tableName: right1
//          isStructShape: true
//          metadata:
//          - s: string
//          data:
//          - - s: "thirty-two"
//          - - i: 2
//            - s: "ssss"
//            - right:
//        - tableName: right2
//          metadata:
//          - x: int
//          - y: int
//          - z: int
//          data:
//          - - x: 1
//            - y: 2
//            - z: 3
//          - - x: 4
//            - y: 5
//            - z: 6
//          - - x: 7
//            - y: 8
//            - z: 9
//          - - i: 3
//            - s: "xxxx"
//            - right:
//        - tableName: right3
//          isStructShape: true
//          metadata:
//          - f: float32
//          data:
//          - - f: 88.8
//          - - i: 4
//            - s: "yyyy"
//            - right:
//        - tableName: right4
//          isStructShape: true
//          metadata:
//          - t1: "*Table"
//          data:
//          - - t1:
//        - tableName: 
//          metadata:
//          data:
//`
//
//	tableSet, err = gotables.NewTableSetFromYAML(yaml)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//where(tableSet.String())
//}
*/

func Test_NewTableSetFromYAML(t *testing.T) {

	var err error
	var tableSet1 *gotables.TableSet
//	var tableSet2 *gotables.TableSet
	var tableSetString string
	var yamlString string

	tableSetString = `
	[[TipTopName]]

	[T0]
	f float64 = 99.90001
	u uint8 = 8
	c byte = 100

	[T1]
	a int = 1
	b int = 2
	c int = 3
	y int = 4

	[T2]
	x		y		s
	bool	byte	string
	true	42		"forty two"
	false	55		"fifty-five"
	true	66		"sixty six"
	`
	tableSet1, err = gotables.NewTableSetFromString(tableSetString)
	if err != nil {
		t.Fatal(err)
	}
where("\n" + tableSet1.String())
println()

where()
	yamlString, err = tableSet1.GetTableSetAsYAML()
where(err)
	if err != nil {
		t.Fatal(err)
	}
where()
println()
where("\n" + yamlString)

/*
	tableSet2, err = gotables.NewTableSetFromYAML(yamlString)
	if err != nil {
		t.Fatal(err)
	}
	_ = tableSet2
*/
}
