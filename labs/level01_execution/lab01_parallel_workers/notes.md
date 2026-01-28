# Testing concurrent behavior

## can test with go's race detector
`go test -race`

## add scheduling pressure by reducing the max processors available to go 
`runtime.GOMAXPROCS(1)`

## add timeouts to catch any deadlocks
`time.After(2 * time.Second)`

## test invariants

