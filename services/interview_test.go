package services

import (
	"fmt"
	"sync"
	"testing"
)

type Animal interface {
	Speak() string
}

type Dog struct{}

func (*Dog) Speak() string {
	return "bark!"
}
func Test_Interface(t *testing.T) {
	// var animal Animal = Dog{}
	// fmt.Println(animal.Speak())
}
func Test_Defer(t *testing.T) {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

func Test_Range(t *testing.T) {
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
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
}
func Test_WaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_Map(t *testing.T) {
	m := make(map[int]bool)
	limit := 1000
	wg := sync.WaitGroup{}
	wg.Add(limit * 2)
	for i := 0; i < limit; i++ {
		go func(idx int) {
			defer wg.Done()
			m[idx] = true
		}(i)
	}

	for i := 0; i < limit; i++ {
		go func(idx int) {
			defer wg.Done()
			fmt.Println(fmt.Sprintf("m[%03d]: %v", idx, m[idx]))
		}(i)
	}
	wg.Wait()
}

func Test_Chan(t *testing.T) {
	messages := make(chan string)

	messages <- "msg1"
	messages <- "msg2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
