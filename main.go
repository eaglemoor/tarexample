package main

import (
	"time"

	"github.com/tarantool/go-tarantool"
	"log"
	"sync/atomic"
)

var I int64
var counter int64

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

	// dataChan := make(chan []int, 100000)
	// for i := 0; i <= 1000; i ++ {
	// 	push(conn, dataChan)
	// }

	timer()

	var i int
	for {
		i++
		// dataChan <- []int{i, i, i, i, i}
		pushAsync(conn, []int{i, i, i, i, i, i})
	}
}

func push(conn *tarantool.Connection, data chan []int) {
	go func() {
		for {
			row := <- data

			fut := conn.InsertAsync("test", row)
			_, err := fut.Get()
			check(err)

			atomic.AddInt64(&counter, 1)
		}
	}()
}

func pushAsync(conn *tarantool.Connection, data []int) {
	conn.InsertAsync("test", data)
	counter++
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