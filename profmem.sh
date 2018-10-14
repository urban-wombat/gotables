rm -f junk

# Memory profile
  go test -memprofile=mem.out
# go tool pprof -top -alloc_objects -cum mem.out > junk; vi junk
  go tool pprof -top -alloc_space -cum mem.out > junk; vi junk
