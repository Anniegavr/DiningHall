package utils


import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Single struct {

	values []int
}
var singleInstance *Single

func getInstance(rate int) *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creting Single Instance Now")
			singleInstance = &Single{}
			singleInstance.values = []int{}
		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}

	return singleInstance
}

func (single *Single) Add_rate(rate int) {
	lock.Lock()

	single.values = append(single.values, rate)

	fmt.Println(single.values)
	finalRating := 0.0
	for _,v:=range single.values {
		finalRating += float64(v)
	}
	fmt.Println("Overall rating:", finalRating/float64(len(single.values)))

	lock.Unlock()

}
