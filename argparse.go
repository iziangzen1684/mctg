package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Parse command line argument
func parse_arg() {
	argv := os.Args
	argc := len(argv)
	if argc == 1 {
		return
	}
	pt := 1
	var v string

	for pt < argc {
		v = argv[pt]

		if v == "-O" || v == "--order" { // Change the order
			pt++
			if pt >= argc {
				log.Fatalf("%v :Value not specified.\n", v)
			}

			na := argv[pt]
			new_order, err := strconv.Atoi(na)
			if err != nil {
				log.Fatalf("%v :Invalid order length.\n", v)
			}

			order = new_order
		} else if v == "-d" || v == "--data" {
			pt++
			if pt >= argc {
				log.Fatalf("%v :Value not specified.\n", v)
			}

			new_data := argv[pt]

			if !file_exists(new_data) {
				log.Fatalf("%v :File not exist: %v\n", v, new_data)
			}
			data_path = new_data
		} else if v == "-c" || v == "--token_count" {
			pt++
			if pt >= argc {
				log.Fatalf("%v :Value not specified.\n", v)
			}

			na := argv[pt]
			new_count, err := strconv.Atoi(na)
			if err != nil {
				log.Fatalf("%v :Invalid token count.\n", v)
			}

			token_count = new_count
		} else if v == "-h" || v == "--help" {
			help = 1
		} else {
			log.Fatalf("Invalid flag: %v\n", v)
		}
		pt++
	}
}

func file_exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	// other errors (e.g., permission issues)
	return false
}

func helper() {
	help := []string{
		"Usage: mctg [arguments]\n",
		"Available flags:\n",
		"	-O, --order: Change the Markov chain order length.\n",
		"	-d, --data: Change the data path.\n",
		"	-c, --token_count: Change the token count.\n",
		"	-h, --help: Display this help message and exit.\n",
	}
	for _, v := range help {
		fmt.Print(v)
	}
	os.Exit(0)
}
