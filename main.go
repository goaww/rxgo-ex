package main

import (
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/iterable"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func main() {
	watcher := observer.Observer{

		NextHandler: func(item interface{}) {
			fmt.Printf("Processing: %v\n", item)
		},

		ErrHandler: func(err error) {
			fmt.Printf("Encountered error: %v\n", err)
		},

		DoneHandler: func() {
			fmt.Println("Done!")
		},
	}

	i, _ := iterable.New([]interface{}{1, 2, 35, 6, errors.New("Bang"), 6})

	sequence := observable.From(i)
	<-sequence.Map(func(num interface{}) interface{} {
		fmt.Println("processing: ", num)
		return num
	}).Filter(func(num interface{}) bool {
		switch num.(type) {
		case int:
			return num.(int)%2 == 0
		}
		return true
	}).Subscribe(watcher)

}
