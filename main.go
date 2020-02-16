package main

import (
	"fmt"
	"github.com/reactivex/rxgo/handlers"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func main() {
	sequence := observable.Just([]int{1, 2, 35, 6, 6})
	<-sequence.FlatMap(func(num interface{}) observable.Observable {
		return observable.Create(func(emitter *observer.Observer, disposed bool) {
			for n := range num.([]int) {
				emitter.OnNext(n)
			}
			emitter.OnDone()
		})
	}, 1).Filter(func(num interface{}) bool {
		return num.(int)%2 == 0
	}).Subscribe(handlers.NextFunc(func(num interface{}) {
		fmt.Println("Result:", num)
	}))
}
