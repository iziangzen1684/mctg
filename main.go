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
<<<<<<< HEAD
	if help == true {
=======
	if help != 0 {
>>>>>>> 9572ec20b2c6849f4ff9a8f582d2463acf00f3b8
		helper()
	}

	// Additional check
	if !file_exists(data_path) {
<<<<<<< HEAD
		log.Fatalf("File %v doesn't exist\n", data_path)
=======
		log.Fatalf("File %v not exist\n", data_path)
>>>>>>> 9572ec20b2c6849f4ff9a8f582d2463acf00f3b8
	}

	// Retrieve file content
	file, err := os.ReadFile(data_path)
	if err != nil {
		log.Fatalf("Can't read file %v: %v\n", data_path, err)
	}
<<<<<<< HEAD

=======
>>>>>>> 9572ec20b2c6849f4ff9a8f582d2463acf00f3b8
	data := string(file)

	// Main table logic
	markov_table := create_markov_table()
	train_markov_table(data, markov_table)
	fmt.Println(strings.Join(generate_text(markov_table, token_count), ""))
}
