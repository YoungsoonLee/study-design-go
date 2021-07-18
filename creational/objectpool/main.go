package main

import (
	"log"
	"strconv"
)

func main() {
	connection := make([]iPoolObject, 0)

	for i := 0; i < 3; i++ {
		c := &connection{
			id: strconv.Itoa(i),
		}
		connection = append(connection, c)
	}

	pool, err := initPool(connection)
	if err != nil {
		log.Fatalf("Init pool error: %s", err)
	}

	conn1, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool loan error: %s", err)
	}

	conn2, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool loan error: %s", err)
	}
	pool.receive(conn1)
	pool.receive(conn2)
}
