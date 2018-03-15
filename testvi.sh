rm -f junk
# go test &> junk
go test -run TestDeleteRow &> junk
vi +?"gotables_test.go" junk
