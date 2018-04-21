# Note: this flag may be one that is by default off even with all turned on.
# See: https://golang.org/cmd/vet
#      Shadowed variables
#      Flag: -shadow=false (experimental; must be set explicitly)

go vet -shadow=true
