rm -f junk

# CPU profile
  go test -bench=. -cpuprofile=cpu.out				# generate profile results into cpu.out
  go tool pprof cpu.out								# view output. Use command: top

  # After quit from pprof, will open junk
  go tool pprof -top -cum cpu.out > junk; vi junk
