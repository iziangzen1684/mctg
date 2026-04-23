package main

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

func if_exist(val rune, str string) bool {
	for _, v := range str {
		if val == v {
			return true
		}
	}
	return false
}

// Tokenizer for Markov chain
func cusparse(text string) []string {
	var temp string
	var temp_list []string
	for _, s := range text {
		// Define a range of characters to be treated as tokens
		if if_exist(s, "\n\t ()[]{}!,.;:<>?/\\@#$%^&*_-+=~`\"'“”‘’") {
			if temp != "" {
				temp_list = append(temp_list, temp)
				temp = ""
			}
			temp_list = append(temp_list, string(s))
		} else {
			temp += string(s)
		}
	}
	// Check if temp is not empty
	if temp != "" {
		temp_list = append(temp_list, temp)
	}
	return temp_list
}

// Initiallize a Markov table
func create_markov_table() map[string][]string {
	markov_table := make(map[string][]string)
	return markov_table
}

// Train a Markov table
func train_markov_table(data string, markov_table map[string][]string) {
	tokens := cusparse(data)
	pt := order
	data_length := len(tokens)

	var key string

	if data_length < order+1 {
		log.Fatalf("Size of tokenized data (%v) should be larger than order size (%v).\n", data_length, order)
	}

	for pt < data_length {
		key = strings.Join(tokens[pt-order:pt], "")
		markov_table[key] = append(markov_table[key], tokens[pt])
		pt++
	}
}

func generate_text(markov_table map[string][]string, seq_len int) []string {

	if seq_len <= 0 {
		log.Fatalf("Sequence length must be larger than 0.\n")
	}

	table_size := len(markov_table)
	pt := order
	var sequence []string

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randint := r.Intn(table_size)

	keys := make([]string, 0, table_size)
	for k := range markov_table {
		keys = append(keys, k)
	}

	key := keys[randint]
	sequence = append(sequence, cusparse(key)...)

	for i := 0; i < seq_len; i++ {
		if pt > len(sequence) {
			break
		}

		key = strings.Join(sequence[pt-order:pt], "")

		randint = r.Intn(len(markov_table[key]))
		sequence = append(sequence, markov_table[key][randint])
		pt++
	}

	return sequence
}
