package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type MenuItem struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	PreparationTime  int    `json:"preparation-time"`
	Complexity       int    `json:"complexity"`
	CookingApparatus string `json:"cooking-apparatus"`
}
type CookingDetails struct {
	FoodId int `json:"food_id"`
	CookId int `json:"cook_id"`
}

type Order struct {
	OrderId       int    `json:"id"`
	TableId       int `json:"table_id"`
	WaiterId       int `json:"table_id"`
	Items    []int `json:"items"`
	Priority int    `json:"priority"`
	MaxWait  int    `json:"maxWait"`
	PickTime int `json:"pick_time"`
}

type OrderDistribution struct {
	OrderId   int `json:"order_id"`
	TableId   int `json:"table_id"`
	Waiter_id int `json:"waiter_id"`
	Items           []int                  `json:"items"`
	Priority     int `json:"priority"`
	MaxWait      int `json:"max_wait"`
	PickUpTime      int                    `json:"pick_up_time"`
	CookingTime    int              `json:"cooking_time"`
	CookingDetails []CookingDetails `json:"cooking_details"`
}

type Menu []MenuItem
type OrdersList []Order
var orderIdentif = 0

func sendOrder() {
	orderIdentif++
	items := []int{3, 4, 4, 5}
	order := Order{
		orderIdentif, 1, 1, items, 3, 45, 1631453140,
	}
	jsonData, err := json.Marshal(order)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8082/orders", "application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(time.Now().Clock())
	fmt.Printf(": Order number %d placed. Status: %d\n", orderIdentif, resp.StatusCode)
}

func generateOrder() {
	i := 1
	max := 10
	for i <= max {
		// wait for 3-10 seconds betwwen placing orders
		preparation_time := rand.Intn(5)
		time.Sleep(time.Duration(preparation_time) * time.Second)
		sendOrder()
		i += 1
	}
}

type orders []int                                      //a list of orders
func indexPage(w http.ResponseWriter, r *http.Request) {}

func postDishes(w http.ResponseWriter, r *http.Request) {
	var prepared OrderDistribution
	err := json.NewDecoder(r.Body).Decode(&prepared)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Print(time.Now().Clock())
	fmt.Printf(": Dishes received. Order id: %d\n", prepared.OrderId)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", indexPage)
	myRouter.HandleFunc("/distribution", postDishes).Methods("POST")
	log.Fatal(http.ListenAndServe(":8083", myRouter))
}

func main() {
	rand.Seed(3000)
	go generateOrder()
	rand.Seed(time.Now().UnixNano())
	handleRequests()
}
