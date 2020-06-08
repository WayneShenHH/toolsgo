package services

import (
	"fmt"
	"sync"
	"testing"
)

func Test_InterfaceCorrect(t *testing.T) {
	var animal Animal = &Dog{}
	fmt.Println(animal.Speak())
}

func Test_RangeCorrect(t *testing.T) {
	type student struct {
		Name string
		Age  int
	}

	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for idx := range stus {
		m[stus[idx].Name] = &stus[idx]
	}
}

func Test_WaitGroupCorrect(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("A: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Test_MapCorrect(t *testing.T) {
	m := make(map[int]bool)
	limit := 1000
	var mux sync.Mutex
	wg := sync.WaitGroup{}
	wg.Add(limit * 2)
	for i := 0; i < limit; i++ {
		go func(idx int) {
			mux.Lock()
			defer mux.Unlock()
			defer wg.Done()
			m[idx] = true
		}(i)
	}

	for i := 0; i < limit; i++ {
		go func(idx int) {
			mux.Lock()
			defer mux.Unlock()
			defer wg.Done()
			fmt.Println(fmt.Sprintf("m[%03d]: %v", idx, m[idx]))
		}(i)
	}
	wg.Wait()
}

func Test_ChanCorrect(t *testing.T) {
	messages := make(chan string, 2)

	messages <- "msg1"
	messages <- "msg2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
