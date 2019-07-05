## Utility executables for gotables

Each utility begins with the prefix *got* for `gotables`.

For Go (golang) programmers:

    go get -u github.com/urban-wombat/gotables

## Here are the utilities ...

* `gotsyntax`
  - `gotsyntax <files>`

	Check the syntax of one or more `gotables` files
	
  - [gotsyntax details](https://github.com/urban-wombat/gotables/tree/master/cmd/gotsyntax)

* `gotecho`
  - `gotecho -f <gotables-file> [-t <this-table-only>] [-r <rotate-table>]`

	Echo a file of `gotables` to stdout, or just one table with -t \<this-table-only\>

	One table only may be rotated using -r \<rotate-table\> to rotate a table syntactically from tablular to struct,
	or from struct to tabular

	Rotate tabular-to-struct is ignored if table has multiple rows, because struct allows only 0 or 1 "rows" of data
  - [gotecho details](https://github.com/urban-wombat/gotables/tree/master/cmd/gotecho)

### Conventional suffix for gotables files ...

`gotables` files by convention are named with a `.got` suffix, but you can call them anything you like.
