package main

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

type TableStatus int

const (
	FreeForClients TableStatus = 0
	Ordering                   = 1
	Waiting                    = 2
)

type Table struct {
	id     int
	status TableStatus

	orderStatus chan bool

	mutex sync.Mutex
}

//func OccupyTable(
//	id int,
//	return &Table{
//		id:          id,
//			TableStatus: make(chan bool)
//}

func (table *Table) GetId() int {
	return table.id
}

func (table *Table) GetStatus() TableStatus {
	return table.status
}

func (table *Table) GetStatusMakingOrder() bool {
	return table.status == Ordering
}

func (table *Table) Run() {
	for {
		table.update()
	}
}

func (table *Table) StartMakeOrder() error {
	table.mutex.Lock()

	if table.status != Ordering {
		return errors.New("busy")
	}
	table.status = Waiting

	table.mutex.Unlock()

	return nil
}

//func (table *Table) GetOrder(distribution *OrderDistribution.CookingDetails) {
//	<-table.orderStatus
//}

func (table *Table) update() {
	table.free()
	table.makingOrder()
}

func (table *Table) free() {
	table.status = FreeForClients
	time.Sleep(300)
}

func (table *Table) makingOrder() {
	table.status = Ordering
	table.orderStatus <- true
}

func (table *Table) getPriority() int {

	priority := rand.Int()

	return priority
}

func (table *Table) getOrderCount() int {

	orderItemsCount := rand.Int()

	return orderItemsCount
}
