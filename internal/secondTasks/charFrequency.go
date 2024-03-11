package secondTasks

func CountLetters(input string, results chan<- map[rune]int) {
	freq := make(map[rune]int)
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			freq[char]++
		}
	}
	results <- freq
}

func MergeResults(results ...map[rune]int) map[rune]int {
	finalResult := make(map[rune]int)
	for _, result := range results {
		for char, count := range result {
			finalResult[char] += count
		}
	}
	return finalResult
}
