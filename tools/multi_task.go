package tools

import (
	"fmt"
	"sync"
)

func MultiStart(taskNumber int, job func()) {
	var wg sync.WaitGroup
	wg.Add(taskNumber)
	for i := 0; i < taskNumber; i++ {
		go func(i int) {
			defer wg.Done()
			job()
		}(i)
	}
	wg.Wait()
	fmt.Println("Finished for loop")
}
