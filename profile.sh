  rm -f junk

# CPU profile
# go test -cpuprofile=cpu.out
# go tool pprof cpu.out
# go tool pprof -top -cum cpu.out > junk; vi junk

# Memory profile
  go test -memprofile=mem.profile
# go tool pprof -top -alloc_objects -cum mem.profile > junk; vi junk
  go tool pprof -top -alloc_space -cum mem.profile > junk; vi junk
