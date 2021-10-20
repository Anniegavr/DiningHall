package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Anniegavr/Lobby/Lobby/item"
	"github.com/Anniegavr/Lobby/Lobby/utils"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
)


//type CookingDetails struct {
//	FoodId int `json:"food_id"`
//	CookId int `json:"cook_id"`
//}

//type Order struct {
//	OrderId  int   `json:"id"`
//	TableId  int   `json:"table_id"`
//	WaiterId int   `json:"table_id"`
//	Items    []int `json:"items"`
//	Priority int   `json:"priority"`
//	MaxWait  int   `json:"maxWait"`
//	PickTime int   `json:"pick_time"`
//}



//type TableIdCounter struct {
//	id int
//}
//func New() *TableIdCounter {
//	return &TableIdCounter{
//		id: 0,
//	}
//}
//
//func (counter *TableIdCounter) getIdRef() *int {
//	return &counter.id
//}
//
//func (counter *TableIdCounter) Get() int {
//	idRef := counter.getIdRef()
//	defer func() {
//		*idRef++
//	}()
//
//	return *idRef
//}

type Menu []item.Item
type OrdersList []utils.OrderData

var orderIdentif = 0

//var sendOrder SendOrder = new
func generateOrder() {
	i := 1
	max := 10
	for i <= max {
		// wait for 3-10 seconds betwwen placing orders
		preparation_time := rand.Intn(5)
		time.Sleep(time.Duration(preparation_time) * time.Second)
		//sendOrder.SendOrder()
		i += 1
	}
}

type orders []int                                      //a list of orders
func indexPage(w http.ResponseWriter, r *http.Request) {}



func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", indexPage)
	myRouter.HandleFunc("/distribution", SendRequest()).Methods("POST")
	log.Fatal(http.ListenAndServe(":8083", myRouter))
}

func main() {
	rand.Seed(3000)
	go generateOrder()
	rand.Seed(time.Now().UnixNano())
	handleRequests()
}
