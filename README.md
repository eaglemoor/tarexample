# Install

1. Install [Docker](https://www.docker.com/community-edition#/download)
1. Install [Go 1.10](https://golang.org/dl/)
1. Install [Dep](https://github.com/golang/dep#installation)
1. Clone project to _$GOPATH_ `cd $GOPATH/src && git clone https://github.com/eaglemoor/tarexample`
1. Install Dependency `cd $GOPATH/src/tarexample && dep ensure`
1. Use `make` 

# Make

```bash
make build-tarantool
```

Build new Docker image from tarantool source v 1.10 branch [gh-3320-config-msg-max](https://github.com/tarantool/tarantool/tree/gh-3320-config-msg-max)

---


```bash
make db
```

Create & run tarantool in docker. See [Dockerfile](./Dockerfile)

---

```bash
make run
```

Run program for test


# Config

You can change tarantool version in [Dockerfile](./Dockerfile)

You can change `RateLimit` in [main.go](./main.go)


# Metrics

```
tarantool 1.9
RateLimit 300

make run
go run main.go
2018/05/02 17:33:56 proccess send 118579, get 118304 records by 1s
2018/05/02 17:33:57 proccess send 113862, get 113950 records by 1s
2018/05/02 17:33:58 proccess send 114286, get 114173 records by 1s
2018/05/02 17:33:59 proccess send 118576, get 118576 records by 1s
2018/05/02 17:34:00 proccess send 121132, get 121132 records by 1s
2018/05/02 17:34:01 proccess send 94026, get 94026 records by 1s
2018/05/02 17:34:02 tarantool: connection localhost:3301 got unexpected resultId (759366) in response
2018/05/02 17:34:02 tarantool: connection localhost:3301 got unexpected resultId (759367) in response
2018/05/02 17:34:02 tarantool: connection localhost:3301 got unexpected resultId (759368) in response
```

```
tarantool 1.10.0-137-gf12f5ccad
Intel(R) Xeon(R) CPU E3-1245 v6 @ 3.70GHz
256GB Intel Pro 5400s SSD

$ make run
go run main.go
2018/05/04 13:22:19 proccess send 473141, get 470930 records by 1s
2018/05/04 13:22:20 proccess send 469064, get 469438 records by 1s
2018/05/04 13:22:21 proccess send 456111, get 456441 records by 1s
2018/05/04 13:22:22 proccess send 472051, get 469906 records by 1s
2018/05/04 13:22:23 proccess send 475745, get 476567 records by 1s
2018/05/04 13:22:24 proccess send 470603, get 471612 records by 1s
2018/05/04 13:22:25 proccess send 479104, get 478790 records by 1s
2018/05/04 13:22:26 [ERR] Failed to allocate 51 bytes in slab allocator for memtx_tuple (0x2)
```
