# gominer
GPU miner for bitcoinx(BCX) in go

All available opencl capable GPU's are detected and used in parallel.

## Binary releases

[Binaries for Windows and Linux are available in the corresponding releases](https://github.com/bitcoinx-project/gominer/releases)


## Installation from source

### Prerequisites
* go version 1.4.2 or above (earlier version might work or not), check with `go version`
* opencl libraries on the library path
* gcc

```
go get github.com/bitcoinx-project/gominer
```

## Run
```
gominer
```

Usage:
```
  -url string
        for stratum servers, use `stratum+tcp://<host>:<port>`
  -user string
        username, most stratum servers take this in the form [payoutaddress].[rigname]
        This is optional, if solo mining bcx, this is not needed
  -I int
    	Intensity (default 28)
  -E string
        Exclude GPU's: comma separated list of devicenumbers
  -cpu
    	If set, also use the CPU for mining, only GPU's are used by default
  -v	Show version and exit
```

See what intensity gives you the best hashrate, increasing the intensity also increases the stale rate though.

## Stratum support

Stratum support is implemented as defined on [stratum_mining_protocol](https://en.bitcoin.it/wiki/Stratum_mining_protocol)
