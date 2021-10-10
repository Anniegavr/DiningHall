package main

import (
	//"Table.go"
	_ "fmt"
	_ "time"
)

type Status int

const (
	Free             Status = 0
	Busy                    = 1
	LookingForOrders        = 2
)

type Waiter struct {
	id int
	//assignedTables []*table.Table
	//queue          *queue.Queue
	status chan bool
}

func (waiter *Waiter) GetId() int {
	return waiter.id
}

//func (waiter *Waiter) Run() {
//	for {
//		waiter.update()
//	}
//}

//func (waiter *Waiter) AddTable(table *tables.Table) {
//
//	waiter.assignedTables = append(waiter.assignedTables, table)
//}
//
//func (waiter *Waiter) getTableById(id int) *table.Table {
//	for _, table := range waiter.assignedTables {
//		if table.GetId() == id {
//			return table
//		}
//	}
//	return nil
//}
//func (waiter *Waiter) GetStatus() TableStatus {
//	return waiter.status
//}
//
//func (waiter *Waiter) update() {
//	return waiter.status == 1
//
//}
