package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	basic_commands := `basic commands:
  uuid          generate a valid uuid

usage:
  generator [command]
`
	if len(os.Args) != 2 {
		fmt.Println(basic_commands)
	} else {
		switch os.Args[1] {
		case "uuid":
			fmt.Println(generate_uuid())
		default:
			fmt.Println(basic_commands)

		}
	}
}

func generate_uuid() (uuid string) {
	values := []string{
		generate_n_random_char(8),
		generate_n_random_char(4),
		generate_n_random_char(4),
		generate_n_random_char(4),
		generate_n_random_char(12),
	}

	return strings.Join(values, "-")

}

func generate_n_random_char(n int) (s string) {
	for i := 0; i < n; i++ {
		s += get_random_char()
	}

	return s

}

func get_random_char() (char string) {
	i := rand.Intn(16)
	switch i {
	case 10:
		return "a"
	case 11:
		return "b"
	case 12:
		return "c"
	case 13:
		return "d"
	case 14:
		return "e"
	case 15:
		return "f"
	default:
		return strconv.Itoa(i)
	}
}
