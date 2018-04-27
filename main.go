package main

import (
	"time"

	"github.com/tarantool/go-tarantool"
	"log"
	"sync/atomic"
)

type respTupl struct {
	i int
	fut *tarantool.Future
}

var I int64
var counter int64
var resp chan respTupl

func conn(address string) (*tarantool.Connection, error) {
	opts := tarantool.Opts{
		User: "guest",
		Timeout:       500 * time.Millisecond,
		Reconnect:     1 * time.Second,
		MaxReconnects: 3,
	}
	conn, err := tarantool.Connect(address,  opts)

	return conn, err
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := conn("localhost:3301")
	check(err)

	timer()

	resp = make(chan respTupl, 1000000)

	go readAsync()

	var i int
	for {
		i++
		pushAsync(conn, i, []int{i, i, i, i, i, i})
	}
}

func pushAsync(conn *tarantool.Connection, i int, data []int) {
	tup := conn.InsertAsync("test", data)
	resp <- respTupl{
		i: i,
		fut: tup,
	}
}

func readAsync() {
	for r := range resp {
		_, err  := r.fut.Get()
		if err != nil {
			log.Printf("[ERR] %s", err)
			continue
		}
		counter++
	}
}

func timer() {
	go func() {
		ti := time.NewTicker(time.Second)
		var c int64

		for range ti.C {
			val := atomic.LoadInt64(&counter)

			log.Printf("proccess %d records by %s\n", val-c, time.Second)

			c = val
		}
	}()
}