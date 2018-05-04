package main

import (
	"time"

	"github.com/tarantool/go-tarantool"
	"log"
	"math/rand"
	"os"
	"sync/atomic"
)

type respTupl struct {
	i   int
	fut *tarantool.Future
}

var I int64
var counterSend, counterGet int64
var resp chan respTupl

func conn(address string) (*tarantool.Connection, error) {
	opts := tarantool.Opts{
		User:          "guest",
		Timeout:       5000 * time.Millisecond,
		Reconnect:     1 * time.Second,
		MaxReconnects: 3,
		RateLimit:     5000,
		RLimitAction:  tarantool.RLimitWait,
	}
	conn, err := tarantool.Connect(address, opts)

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

	rand.Seed(time.Now().Unix())
	rng := rand.New(rand.NewSource(rand.Int63()))

	var i int
	for {
		i++
		pushAsync(conn, i, []int{rng.Int(), i, i, i, i, i})
	}
}

func pushAsync(conn *tarantool.Connection, i int, data []int) {
	tup := conn.InsertAsync("test", data)
	atomic.AddInt64(&counterSend, 1)
	resp <- respTupl{
		i:   i,
		fut: tup,
	}
}

func readAsync() {
	c := 0
	for r := range resp {
		if c > 10 {
			os.Exit(1)
		}

		_, err := r.fut.Get()
		if err != nil {
			log.Printf("[ERR] %s", err)
			c++
			continue
		}
		atomic.AddInt64(&counterGet, 1)
	}
}

func timer() {
	go func() {
		ti := time.NewTicker(time.Second)
		var send, get int64

		for range ti.C {
			valGet := atomic.LoadInt64(&counterGet)
			valSend := atomic.LoadInt64(&counterSend)

			log.Printf("proccess send %d, get %d records by %s\n", valSend-send, valGet-get, time.Second)

			get = valGet
			send = valSend
		}
	}()
}
