package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Set log flag to 0 and output to stderr
	log.SetOutput(os.Stderr)
	log.SetFlags(0)

	parse_arg()
	if help != 0 {
		helper()
	}

	// Additional check
	if !file_exists(data_path) {
		log.Fatalf("File %v not exist\n", data_path)
	}

	// Retrieve file content
	file, err := os.ReadFile(data_path)
	if err != nil {
		log.Fatalf("Can't read file %v: %v\n", data_path, err)
	}
	data := string(file)

	// Main table logic
	markov_table := create_markov_table()
	train_markov_table(data, markov_table)
	fmt.Println(strings.Join(generate_text(markov_table, token_count), ""))
}
