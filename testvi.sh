rm -f junk
go test &> junk
vi +?"gotables_test.go" junk
