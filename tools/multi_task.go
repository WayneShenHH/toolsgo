package tools

import (
	"fmt"
	"sync"
)

func MultiStartWithGroup(taskNumber int, job func(index int)) {
	var wg sync.WaitGroup
	wg.Add(taskNumber)
	for i := 0; i < taskNumber; i++ {
		go func(i int) {
			defer wg.Done()
			job(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("Finished for loop")
}
