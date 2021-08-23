# GO simple port scanner

## Introduction
A simple implementation of a port scanner, comparing the speed between a linear scan and with the usage of goroutines.
To test it, just replace the `nGoroutines` variables with the desired value.

**Note**: Be careful with the max number of goroutines and the target you test. You can generate some issues.

## Execute
Simple usage:
```sh
go run portscan.go
```

For execution only with goroutines, just pass a value (ex: 1):
```sh
go run portscan.go 1
```

## TODO
- [ ] Receive parameters as input (host, number of goroutines)
- [ ] Receive port range as input (ex: 1-1000)
- [ ] Store results into a file
- [ ] Sort results by port
