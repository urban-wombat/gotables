rm -f junk

# Memory profile
  go test -memprofile=mem.profile
# go tool pprof -top -alloc_objects -cum mem.profile > junk; vi junk
  go tool pprof -top -alloc_space -cum mem.profile > junk; vi junk
