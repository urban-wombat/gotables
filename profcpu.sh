rm -f junk

# CPU profile
  go test -cpuprofile=cpu.out
  go tool pprof cpu.out
  go tool pprof -top -cum cpu.out > junk; vi junk
