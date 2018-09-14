[![Build Status](https://travis-ci.org/vikyd/go-cpu-load.svg?branch=master)](https://travis-ci.org/vikyd/go-cpu-load)

# go-cpu-load

Generate CPU load on Windows/Linux.

# Install

```sh
go get -u github.com/vikyd/go-cpu-load
```

or download binary file directly

# Usage

example 01: run 30% of all CPU cores for 10 seconds

```sh
go-cpu-load -p 30 -t 10
```

example 02: run 30% of all CPU cores forver

```sh
go-cpu-load -p 30
```

example 03: run 30% of 2 of CPU cores for 10 seconds

```sh
go-cpu-load -p 30 -t 10
```

- `all CPU load` = (num of para `c` _ num of `p`) / (all cores count of CPU _ 100)
- may not specify cores run the load only, it just promise the `all CPU load`, and not promise each cores run the same load

# Parameters

```
--coresCount value, -c value   how many cores (optional, default: 8)

--timeSeconds value, -t value  how long (optional, default: 2147483647)

--percentage value, -p value   percentage of each specify cores (required)

--help, -h                     show help
```

# Build

```sh
go build
```

# test

```sh
go test -v
```

> currently only provide Windows testing

# How it runs

- Giving a range of time(e.g. 100ms)
- Want to run 30% of all CPU cores
  - 30ms: run (CPU 100%)
  - 70ms: sleep(CPU 0%)
