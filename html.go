package gotable

/*
	Note: View Source in Firefox (not Chrome) to see corresponding line numbers.
*/

import (
	"errors"
	"fmt"
	"golang.org/x/net/html" // From: https://github.com/golang/go/wiki/SubRepositories
	"io"
	"log"
	"net/http"
	"os"
	//	"path"
	"strings"
)

/*
Copyright (c) 2017 Malcolm Gorman

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

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	var err error

	url := "https://en.wikipedia.org/wiki/List_of_government_schools_in_Victoria,_Australia"
	//	url := "http://www.bom.gov.au/vic/observations/melbourne.shtml"
	//	url := "http://www.bom.gov.au/products/IDV60901/IDV60901.94870.shtml"
	//	url := "http://www.marketindex.com.au/asx200"

	var tableSet *GoTableSet
	tableSet, err = NewGoTableSetFromHtmlUrl(url)
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println()
	//	fmt.Fprintf(os.Stderr, "%s", tableSet)
	fmt.Printf("%s", tableSet)
}

const startTag = "START_TAG"
const endTag = "END_TAG"

func NewGoTableSetFromHtmlString(sourceName string, htmlString string) (*GoTableSet, error) {
	//	where("called NewGoTableSetFromHtmlString()")
	/*
		htmlString = `<table>
		<tr> <td>r0 c0</td> <td>r0 c1</td> <td>r0 c2</td> </tr>
		<tr> <td>r1 c0</td> <td>r1 c1</td> <td>r1 c2</td> </tr>
		<tr> <td>r2 c0</td> <td>r2 c1</td> <td>r2 c2</td> </tr>
		</table>`
		sourceName = "DEBUGGING HTML"
	*/

	var stringReader *strings.Reader = strings.NewReader(htmlString)

	return NewGoTableSetFromHtmlReader(sourceName, stringReader)
}

func NewGoTableSetFromHtmlUrl(url string) (*GoTableSet, error) {
	var err error
	var tableSet *GoTableSet

	response, err := http.Get(url)
	if err != nil {
		var errString string = err.Error()
		var errPrepended error = errors.New("In NewGoTableSetFromHtml(url) http.Get(url) FAILED: " + errString)
		return nil, errPrepended
	}
	defer response.Body.Close()

	// Even though the response was received without error, the request for the url may still fail.
	if response.StatusCode != http.StatusOK {
		return tableSet, fmt.Errorf("http get error status: %s : %q", response.Status, url)
	}

	return NewGoTableSetFromHtmlReader(url, response.Body)
}

func NewGoTableSetFromHtmlReader(sourceName string, reader io.Reader) (*GoTableSet, error) {
	var err error
	var tableSet *GoTableSet
	var table *GoTable
	var z *html.Tokenizer
	var tableIndex int = -1
	var rowIndex int = -1
	var colIndex int = -1
	var startToken string = "" // Used for switching upon data.
	var prevStartToken string = ""
	var isFirstRow bool = false // Until we see a row in a new table with <td> tags in it.
	var isTdRow = false
	var isFirstRowDone bool = false
	var insideTable bool = false
	var insideRow bool = false
	var insideCol bool = false
	var accumulated string
	var colName string
	var sourceLineNumber = 1
	const colType = "string" // Always string because we don't know if there are valid numbers and booleans.
	const rowsToPrint = 2

	// Notice when a tag isn't closed.
	var prevTableTag string = endTag // To avoid having to test for "".
	var prevTrTag string = endTag    // To avoid having to test for "".
	var prevTdTag string = endTag    // To avoid having to test for "".

	tableSet, err = NewGoTableSet("")
	if err != nil {
		return nil, err
	}

	//	prevLineNumber := sourceLineNumber

	//	debugging := false
	//	if debugging {
	//		// Debugging
	//		htmlString := `<table>
	//		<tr> <td>r0 c0</td> <td>r0 c1</td> <td>r0 c2</td> </tr>
	//		<tr> <td>r1 c0</td> <td>r1 c1</td> <td>r1 c2</td> </tr>
	//		<tr> <td>r2 c0</td> <td>r2 c1</td> <td>r2 c2</td> </tr>
	//		</table>`
	//		var stringReader *strings.Reader = strings.NewReader(htmlString)
	//		z = html.NewTokenizer(stringReader)
	//		sourceName = "DEBUGGING HTML"
	//	} else {	// Normal execution.
	//		z = html.NewTokenizer(response.Body)
	//where(fmt.Sprintf("response      type: %T", response))
	//where(fmt.Sprintf("response.Body type: %T", response.Body))
	//	}

	//	z = html.NewTokenizer(response.Body)
	z = html.NewTokenizer(reader)

	for {
		var tokenType html.TokenType = z.Next()

		// Count lines in the HTML for error message.
		raw := z.Raw()
		// text := z.Text()	// Don't use text. It swollows token text.
		//		lineGap := sourceLineNumber - prevLineNumber
		// where(fmt.Sprintf("* %s[%d] len(raw)=%d lineGap=%d: %s", path.Base(sourceName), sourceLineNumber, len(raw), lineGap, string(raw)))
		// where(fmt.Sprintf("* %s[%d] len(raw)=%d: %s", path.Base(sourceName), sourceLineNumber, len(raw), string(raw)))
		//		where(fmt.Sprintf("At %s[%d] %s", path.Base(sourceName), sourceLineNumber, string(raw)))
		//		prevLineNumber = sourceLineNumber
		for i := 0; i < len(raw); i++ {
			if raw[i] == 10 {
				sourceLineNumber++
			}
		}

		token := z.Token()
		var tokenData string = token.Data

		switch tokenType {

		case html.StartTagToken:

			switch tokenData {
			case "table": // start
				//				where(fmt.Sprintf("<table>"))
				if prevTableTag != endTag {
					//					where(fmt.Sprintf("unclosed <table>"))
					// Close the previous table.
					//					where("calling endTheRow()")
					prevTrTag, prevTdTag, insideRow, insideCol, colIndex, isTdRow, isFirstRow, isFirstRowDone = endTheRow(isTdRow, isFirstRow, isFirstRowDone)
					//					where("calling endTheTable()")
					prevTableTag, rowIndex, isFirstRowDone, insideTable, err = endTheTable(tableSet, table)
					if err != nil {
						return nil, fmt.Errorf("%s[%d] %v", sourceName, sourceLineNumber, err)
					}
					fmt.Fprintf(os.Stderr, "\n")
					//					where(fmt.Sprintf("\n%s", tableSet))
					fmt.Fprintf(os.Stderr, "\n")
				}

				prevTableTag = startTag
				startToken = "table"
				//				where("calling startTheTable()")
				tableIndex, table, tableSet, rowIndex, err = startTheTable(tableIndex, tableSet)
				if err != nil {
					return nil, fmt.Errorf("%s[%d] %v", sourceName, sourceLineNumber, err)
				}
				//				where(fmt.Sprintf("tableIndex = %d", tableIndex))
				insideTable = true

			case "tr": // start
				//				where(fmt.Sprintf("<tr>"))
				if !insideTable {
					//					where("orphan <tr>")
					//					where("continue")
					continue
				}
				if prevTrTag != endTag {
					//					where(fmt.Sprintf("unclosed <tr>"))
					// Close the previous row.
					//					where("calling endTheRow()")
					prevTrTag, prevTdTag, insideRow, insideCol, colIndex, isTdRow, isFirstRow, isFirstRowDone = endTheRow(isTdRow, isFirstRow, isFirstRowDone)
				}
				prevTrTag = startTag
				startToken = "tr"
				insideRow = true
				//					case "td":	// start
			case "td", "th": // start
				if tokenData == "td" {
					isTdRow = true
				}
				//				where(fmt.Sprintf("<td>"))
				if !insideRow {
					//					where("orphan <td>")
					//					where("continue")
					continue
				}
				prevTdTag = startTag
				insideCol = true
				colIndex++
				//				where(fmt.Sprintf("colIndex++ = %d", colIndex))
				//				where(fmt.Sprintf("isFirstRow = %t", isFirstRow))
				//				where(fmt.Sprintf("??? if   colIndex == 0"))
				if colIndex == 0 {
					// We know it's a new row with data in it, not <th> headings.
					rowIndex++
					//					where(fmt.Sprintf("setting rowIndex++ = %d", rowIndex))
				}
				startToken = "td"

				//				where(fmt.Sprintf("??? if   !isFirstRow=%t && !isFirstRowDone=%t == %t", !isFirstRow, !isFirstRowDone, !isFirstRow && !isFirstRowDone))
				if !isFirstRow && !isFirstRowDone {
					isFirstRow = true
					//					where(fmt.Sprintf("??? then setting isFirstRow = %t", isFirstRow))
				} else {
					//					where(fmt.Sprintf("??? else NOT setting isFirstRow = true"))
				}
				//				where(fmt.Sprintf("??? if   insideRow=%t && isFirstRow=%t == %t", insideRow, isFirstRow, insideRow && isFirstRow))
				if insideRow && isFirstRow {
					// Create a column.
					colName = fmt.Sprintf("col_%d", colIndex)
					//					where(fmt.Sprintf("??? then [%s].AddCol(%s, %s)", table.TableName(), colName, colType))
					err = table.AddCol(colName, colType)
					if err != nil {
						return nil, fmt.Errorf("%s[%d] %v", sourceName, sourceLineNumber, err)
					}
				} else {
					//					where(fmt.Sprintf("??? else NOT [%s].AddCol(%s, %s)", table.TableName(), colName, colType))
				}

				// We can't reliably add rows after each <tr> because it may contain only <th> tags.
				//				where(fmt.Sprintf("??? if   table.RowCount()=%d == rowIndex=%d == %t", table.RowCount(), rowIndex, table.RowCount() == rowIndex))
				if table.RowCount() == rowIndex {
					// Need exactly one new row.
					//					where(fmt.Sprintf("??? then [%s].AddRow()", table.TableName()))
					table.AddRow()
				} else {
					//					where(fmt.Sprintf("??? else NOT [%s].AddRow()", table.TableName()))
				}
			}

		case html.TextToken:
			switch prevStartToken {
			case "td":
				//				where(fmt.Sprintf("case td <td>%s</td>", tokenData))
				//				where(fmt.Sprintf("insideRow=%t && insideCol=%t == %t", insideRow, insideCol, insideRow && insideCol))
				if insideRow && insideCol { // Don't accumulate token data outside of a row or col.
					//					where()
					if colIndex >= table.ColCount() {
						// This row has more cols than previous rows. Not good, but we'll accommodate that.
						// Create additional column.
						colName = fmt.Sprintf("col_%d", colIndex)
						//						where(fmt.Sprintf("??? then [%s].AddCol(%s, %s)", table.TableName(), colName, colType))
						err = table.AddCol(colName, colType)
						if err != nil {
							return nil, fmt.Errorf("%s[%d] %v", sourceName, sourceLineNumber, err)
						}
					}

					// Read what is already accumulated.
					accumulated, err = table.GetStringByColIndex(colIndex, rowIndex)
					//					where()
					if err != nil {
						//						where()
						//								return nil, err
						return nil, fmt.Errorf("%s[%d] %v", sourceName, sourceLineNumber, err)
					}
					//					where()
					const nbspChar = "\u00a0"
					//							const nbspNum = "&#160;"	// See https://en.wikipedia.org/wiki/Non-breaking_space
					//							const trimChars = " \t\n" + nbspChar + nbspNum	// This seems to remove trailing digits!
					const trimChars = " \t\n" + nbspChar
					//					where()
					const punctuation = ".,;:"

					//					where()
					//					where(fmt.Sprintf("tokenData: %q (before Trim)", tokenData))
					// Make sure there is just one space between tokens in accumulated.
					tokenData = strings.Trim(tokenData, trimChars)
					//					where(fmt.Sprintf("tokenData: %q (after  Trim)", tokenData))
					if strings.ContainsAny(tokenData, punctuation) {
						//						where(fmt.Sprintf("tokenData: %q (contains punctuation)", tokenData))
						// Don't insert space.
						accumulated += tokenData
					} else {
						//						where(fmt.Sprintf("tokenData: %q (does NOT contain punctuation)", tokenData))
						//						where(fmt.Sprintf("accumulated: %q (before inserting space)", accumulated))
						//						where(fmt.Sprintf("tokenData  : %q (before inserting space)", tokenData))
						accumulated += " " + tokenData
						//						where(fmt.Sprintf("accumulated: %q (after  inserting space)", accumulated))
					}
					//					where()

					//					where(fmt.Sprintf("accumulated: %q (before Trim)", accumulated))
					accumulated = strings.Trim(accumulated, trimChars) // Remove white space crud.
					// Write back accumulated plus current text token (minus surrounding white space crud).
					err = table.SetStringByColIndex(colIndex, rowIndex, accumulated)
					//					where(fmt.Sprintf("accumulated: %q (after  Trim)", accumulated))
					if err != nil {
						//								return nil, err
						return nil, fmt.Errorf("%s[%d] %v", sourceName, sourceLineNumber, err)
					}
				}
			}

			//			case html.SelfClosingTagToken:
			//				where(fmt.Sprintf("<%s /> (SelfClosingTagToken)", tokenData))

		case html.EndTagToken:
			switch tokenData {
			case "table": // end
				//				where(fmt.Sprintf("</table>"))
				if prevTableTag != startTag {
					// Orphan table. It will have already been ended. So do nothing.
					//					where(fmt.Sprintf("orphan </table>"))
				} else {
					//					where("calling endTheRow()")
					prevTrTag, prevTdTag, insideRow, insideCol, colIndex, isTdRow, isFirstRow, isFirstRowDone = endTheRow(isTdRow, isFirstRow, isFirstRowDone)
					//					where("calling endTheTable()")
					prevTableTag, rowIndex, isFirstRowDone, insideTable, err = endTheTable(tableSet, table)
					if err != nil {
						//						where()
						//								return nil, err
						return nil, fmt.Errorf("%s[%d] %v", sourceName, sourceLineNumber, err)
					}
				}
			case "tr": // end
				//				where(fmt.Sprintf("</tr>"))
				if prevTrTag != startTag {
					//					where(fmt.Sprintf("orphan </tr>"))
				}
				//				where("calling endTheRow()")
				prevTrTag, prevTdTag, insideRow, insideCol, colIndex, isTdRow, isFirstRow, isFirstRowDone = endTheRow(isTdRow, isFirstRow, isFirstRowDone)
			case "td": // end
				//				where(fmt.Sprintf("</td>"))
				if prevTdTag != startTag {
					//					where(fmt.Sprintf("orphan </td>"))
				}
				// NOTE: Do not set insideCol to false. Sometimes there is no html.TextToken tokenData?
				// There's a logic error here somewhere. There are newlines in tokenData which are not in <td></td>
				// Aha! Maybe they are just after the </td>
				// NOTE: This was fixed by gotable setting new rows and new cols to zero values, not leaving them as nil.
				insideCol = false // 7.7.2016
				prevTdTag = endTag
			}

		case html.ErrorToken:
			// End of document.
			if tableIndex == -1 {
				// Not really an error. Caller can check number of tables returned.
				return tableSet, nil
			}
			// Probably the normal exit from the function.
			return tableSet, nil

		case html.CommentToken:
			//			where(fmt.Sprintf("CommentToken: token = %v", token))
			//			//				where(fmt.Sprintf("text: %s", string(text)))
		}
		//		where(fmt.Sprintf("insideTable    = %t", insideTable))
		//		where(fmt.Sprintf("isFirstRow     = %t", isFirstRow))
		//		where(fmt.Sprintf("isFirstRowDone = %t", isFirstRowDone))
		//		where(fmt.Sprintf("insideRow      = %t", insideRow))
		//		where(fmt.Sprintf("insideCol      = %t", insideCol))
		fmt.Fprintf(os.Stderr, "\n")

		prevStartToken = startToken
	}
}

func startTheTable(tableIndex int, tableSet *GoTableSet) (int, *GoTable, *GoTableSet, int, error) {
	//	where("startTheTable()")
	var err error

	tableIndex++
	rowIndex := -1
	//	where(fmt.Sprintf("tableIndex++ = %d", tableIndex))
	table, err := NewGoTable(fmt.Sprintf("table_%d", tableIndex))
	if err != nil {
		return tableIndex, table, tableSet, rowIndex, err
	}
	err = tableSet.AddTable(table)
	if err != nil {
		return tableIndex, table, tableSet, rowIndex, err
	}

	return tableIndex, table, tableSet, rowIndex, nil
}

func endTheTable(tableSet *GoTableSet, table *GoTable) (string, int, bool, bool, error) {
	//	where("endTheTable()")
	prevTableTag := endTag
	rowIndex := -1 // For good measure. In case </tr> not encountered.
	//	where(fmt.Sprintf("rowIndex = %d", rowIndex))
	isFirstRowDone := false
	insideTable := false
	if isValid, err := table.IsValidTable(); !isValid {
		return prevTableTag, rowIndex, isFirstRowDone, insideTable, err
	}

	fmt.Fprintf(os.Stderr, "\n")
	//	where(fmt.Sprintf("\n%s", tableSet))
	fmt.Fprintf(os.Stderr, "\n")

	//	where(fmt.Sprintf("setting isFirstRowDone = %t", isFirstRowDone))
	return prevTableTag, rowIndex, isFirstRowDone, insideTable, nil
}

func endTheRow(isTdRow bool, isFirstRow bool, isFirstRowDone bool) (string, string, bool, bool, int, bool, bool, bool) {
	//	where("endTheRow()")
	prevTrTag := endTag
	prevTdTag := endTag
	insideRow := false
	insideCol := false // For good measure. In case </td> not encountered.
	colIndex := -1
	//	where(fmt.Sprintf("colIndex = %d", colIndex))
	if isFirstRow == true {
		if isTdRow {
			isFirstRow = false
			isFirstRowDone = true
			//			where(fmt.Sprintf("setting isFirstRowDone = %t", isFirstRowDone))
		} else {
			//			where(fmt.Sprintf("Not a real first row! Undo the column headings!"))
		}
	}

	isTdRow = false

	return prevTrTag, prevTdTag, insideRow, insideCol, colIndex, isTdRow, isFirstRow, isFirstRowDone
}
