# sauced
Managed Sauce Connect tunnels.

## Install and Run
1. `go get github.com/maxdobeck/sauced`
or
2. Clone the repo and put it in `$GOPATH/github.com/maxdobeck/sauced`.

Move to the directory and run `$ go build`.  Run the binary like `./sauced`.  `.\sauced` on windows.  Be aware of other windows specific behavior like %GOPATH% instead of $GOPATH.

## Testing
Run `$ go test ./...`.  This goes through all directories recursively to run anything like `*_test.go`.