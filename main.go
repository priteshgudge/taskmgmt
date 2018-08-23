package main

import (
	"sync"
	"taskmgmt/operators"
)

func main() {
	queue := operators.GetNewQueue()
	consumer := operators.GetNewConsumer(queue)
	go consumer.RunConsumer()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		//defer wg.Done()
		queue.AddMessage(operators.Message{"one"})
	}()
	go func() {
		//defer wg.Done()
		queue.AddMessage(operators.Message{"two"})
	}()
	go func() {
		//defer wg.Done()
		queue.AddMessage(operators.Message{"three"})
	}()
	wg.Wait()

}
