# Produce a coverage output file for analysis.
go test -coverprofile=coverage.out

# Analyze the output file, remove some strings, sort with least covered functions last.
go tool cover -func=coverage.out | egrep -v '^ok|^PASS|^total:' | sort -k 3 -n -r

# Open a web page.
# Note: Use drop-down menu at top left to switch between source files.
go tool cover -html=coverage.out
