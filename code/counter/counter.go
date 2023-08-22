package counter

import (
	"fmt"
	"time"
)

func countTo(limit int, channel chan [2]int, counterId int) {
	for i := 0; i < limit; i++ {
		channel <- [2]int{i, counterId}
		time.Sleep(time.Second)
	}
}

func SetupParallelCounters(countLimit int, numGoroutines int) chan [2]int {
	countingChannel := make(chan [2]int)

	fmt.Printf("\nParallel counters (%d threads):\n", numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		counterId := i + 1
		go countTo(countLimit, countingChannel, counterId)
	}

	return countingChannel
}
