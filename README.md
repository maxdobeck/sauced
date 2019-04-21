# sauced
Managed Sauce Connect tunnels.

## Install and Run
1. `go get github.com/maxdobeck/sauced`
or
2. Clone the repo and put it in `$GOPATH/github.com/maxdobeck/sauced`.

Move to the directory and run `$ go build`.  Run the binary like `./sauced`.  `.\sauced` on windows.  Be aware of other windows specific behavior like %GOPATH% instead of $GOPATH.

Pass in the config file like so `$ ./sauced ~/.config/sauced.txt`.  The specified file should be read line by line and used to start a Sauce Connect instance.  

## Config File
The config file should have one line for each SC instance.  The first portion of the line should be the path to the SC binary you want to use.  The other arguments should be the flags you would use if you were starting the tunnel manually from a command line.  The scheme should be:

```
/path/to/bin/sc <normal arguments>
```

An example of a single tunnel:
`/home/user/tools/sc-4.5.1-linux/bin/sc -u account-name-here -k api-key-here -v`


An example of a pool of tunnels:

```
/home/user/tools/sc-4.5.1-linux/bin/sc -u account-name -k api-key-here -v --no-remove-colliding-tunnels -N -i main-tunnel-pool --se-port 0 
/home/user/tools/sc-4.5.1-linux/bin/sc -u account-name -k api-key-here -v --no-remove-colliding-tunnels -N -i main-tunnel-pool --se-port 0 
```

## Testing
Run `$ go test ./...`.  This goes through all directories recursively to run anything like `*_test.go`.