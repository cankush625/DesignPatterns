package object_pool_design_pattern

import (
	"log"
	"strconv"
)

func main() {
	connections := make([]IPoolObject, 0)
	for i := 0; i < 5; i++ {
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	pool, err := initPool(connections)
	if err != nil {
		log.Fatalf("Init Pool Error: %s", err)
	}
	connection1, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	connection2, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	err = pool.release(connection1)
	if err != nil {
		return
	}
	err = pool.release(connection2)
	if err != nil {
		return
	}
}
