package main

import (
	"fmt"
	"github.com/akashrzp/go-bootcamp/internal/secondTasks"
	"sync"
)

func main() {
	fmt.Printf("Day Two tasks.....\n\n")

	fmt.Printf("Task - Char Frequency -- Start.....\n")

	inputs := []string{"quick", "brown", "fox", "lazy", "dog"}
	results := make(chan map[rune]int, len(inputs))

	var wg sync.WaitGroup
	wg.Add(len(inputs))

	for _, input := range inputs {
		go func(input string) {
			defer wg.Done()
			secondTasks.CountLetters(input, results)
		}(input)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	finalResult := make(map[rune]int)
	for result := range results {
		finalResult = secondTasks.MergeResults(finalResult, result)
	}

	for char, count := range finalResult {
		fmt.Printf("%c: %d\n", char, count)
	}
	fmt.Printf("Task - Char Frequency -- End.....\n\n")

	fmt.Printf("Task - Rating Class Teacher -- Start.....\n")
	var wg2 sync.WaitGroup
	numStudents := 200
	ratings := make(chan int, numStudents)

	// Launch goroutines for each student to rate the teacher
	for i := 0; i < numStudents; i++ {
		wg2.Add(1)
		go secondTasks.RateTeacher(ratings, &wg2)
	}

	// Close the ratings channel after all ratings are received
	go func() {
		wg2.Wait()
		close(ratings)
	}()

	// Collect ratings and calculate the average
	totalRating := 0
	numRatings := 0
	for rating := range ratings {
		totalRating += rating
		numRatings++
	}

	// Print the average rating
	if numRatings > 0 {
		averageRating := float64(totalRating) / float64(numRatings)
		fmt.Printf("Average rating of the class teacher: %.2f\n", averageRating)
	} else {
		fmt.Println("No ratings received.")
	}
	fmt.Printf("Task - Rating Class Teacher -- End.....\n\n")

	fmt.Printf("Task - Bank Account -- Start.....\n")
	account := secondTasks.NewBankAccount(500)
	var wg3 sync.WaitGroup

	// Perform concurrent deposit and withdrawal
	wg3.Add(2)
	go account.Deposit(200, &wg3)
	go account.Withdraw(300, &wg3)
	wg3.Wait()

	fmt.Printf("Final balance: Rs.%d\n", account.Balance)
	fmt.Printf("Task - Bank Account -- End.....\n\n")

}
