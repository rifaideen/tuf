package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// We use a map to store all valid words from wordList (excluding beginWord)
	// This acts as our "unvisited" set — once we use a word in a transformation,
	// we'll remove it to avoid revisiting and forming cycles.
	// Note: We skip beginWord because it's not necessarily in wordList, but we don't want to block it.
	hashmap := map[string]bool{}
	for _, word := range wordList {
		if word != beginWord {
			hashmap[word] = true
		}
	}

	// Early termination: if endWord isn't in the dictionary, we can't reach it
	if !hashmap[endWord] && endWord != beginWord {
		return 0
	}

	// BFS queue: each element is [currentWord, transformationSteps]
	// We use BFS because we're looking for the shortest path in an unweighted graph:
	// - Each word is a node
	// - An edge exists between words that differ by exactly one letter
	// Starting with beginWord at step 1 (the sequence includes the starting word)
	queue := [][]interface{}{
		{beginWord, 1},
	}

	// BFS loop: process each word level by level
	for len(queue) > 0 {
		// Dequeue the next word and the number of steps taken to reach it
		word := queue[0][0].(string)
		steps := queue[0][1].(int)
		queue = queue[1:]

		// If we've reached the target word, return the number of steps
		// This is guaranteed to be the shortest path because BFS explores all paths level by level
		if word == endWord {
			return steps
		}

		// Convert the word into a rune slice for easy character manipulation
		current := []rune(word)

		// Try changing each character in the word to explore all possible one-letter mutations
		for i := 0; i < len(current); i++ { // Fix: was range len(word) → invalid syntax
			original := current[i] // Save the original character

			// Try replacing current character with every letter from 'a' to 'z'
			for char := 'a'; char <= 'z'; char++ {
				if char == original {
					continue // Skip if no change — no point in checking the same word
				}

				current[i] = char          // Mutate: change the i-th character
				newWord := string(current) // Create the new transformed word

				// If this new word exists in our dictionary (i.e., it's a valid next step),
				// then we can enqueue it as part of a valid transformation sequence
				if hashmap[newWord] {
					// Remove the word from the map to mark it as "visited"
					// This prevents cycles and redundant processing, ensuring each word is used only once
					delete(hashmap, newWord)

					// Enqueue the new word with step count + 1
					// This represents extending the transformation sequence by one step
					queue = append(queue, []interface{}{
						newWord,
						steps + 1,
					})
				}
			}

			// Restore the original character before moving to the next position
			// This ensures we only mutate one character at a time
			current[i] = original
		}
	}

	// If we exhaust the BFS without reaching endWord, no valid transformation sequence exists
	return 0
}
