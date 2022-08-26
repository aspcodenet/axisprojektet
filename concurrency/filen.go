package concurrency

import (
	"fmt"
	"strconv"
	"time"
)

func Run4() int {
	return 23
}

func Calculate3(workChannel chan string, result chan string) {
	for {
		nextJob := <-workChannel

		result <- "Starting "
		result <- nextJob
		time.Sleep(time.Second * 5)
		result <- "Ending " + nextJob

	}
}

func RunDemo() {
	resultChannel := make(chan string)
	workChannel := make(chan string, 100)

	for i := 0; i < 100; i++ {
		workChannel <- "Jobb" + strconv.FormatInt(int64(i), 10)
	}
	go Calculate3(workChannel, resultChannel)
	go Calculate3(workChannel, resultChannel)
	go Calculate3(workChannel, resultChannel)
	go Calculate3(workChannel, resultChannel)
	go Calculate3(workChannel, resultChannel)
	go Calculate3(workChannel, resultChannel)
	go Calculate3(workChannel, resultChannel)
	for msg := range resultChannel {
		fmt.Println(msg)
	}

	// go func() {

	// 	Calculate("job2")
	// }()
	for msg := range resultChannel {
		fmt.Println(msg)
	}

}

// func RunDemo() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	funk := func() {
// 		Calculate("job1")
// 		wg.Done()
// 	}
// 	go funk()

// 	go func() {
// 		Calculate("job2")
// 		wg.Done()
// 	}()

// 	wg.Wait()

// }

func Calculate(todo string, resultChannel chan string) {
	if todo == "job1" {

	} else if todo == "job2" {

	}
	resultChannel <- "Starting "
	resultChannel <- todo
	time.Sleep(time.Second * 5)
	resultChannel <- "Ending " + todo
	close(resultChannel)
}
