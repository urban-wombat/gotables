`go get -u github.com/urban-wombat/gotables/cmd/gotecho`

Usage1: `gotecho [-f <gotables-file>] [-t <this-table-only>] [-r <rotate-table>]`

Usage2: `cat <gotables-file> | gotecho [-t <this-table-only>] [-r <rotate-table>]`

Echo a file of `gotables` to stdout, or just one table with -t \<this-table-only\>

One table may be rotated using -r \<rotate-table\> to rotate a table syntactically from tablular to struct,
or from struct to tabular

Rotate tabular-to-struct is ignored if table has multiple rows, because struct allows only 0 or 1 "rows" of data

