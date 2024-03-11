package secondTasks

import (
	"math/rand"
	"sync"
	"time"
)

func RateTeacher(ratings chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate random response time
	responseTime := time.Duration(rand.Intn(100)) * time.Millisecond
	time.Sleep(responseTime)
	// Simulate random rating between 1 and 10
	rating := rand.Intn(10) + 1
	ratings <- rating
}
