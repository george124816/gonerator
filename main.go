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
  cpf           generate a valid cpf
  cnpj          generate a valid cnpj

usage:
  generator [command]
`
	if len(os.Args) != 2 {
		fmt.Println(basic_commands)
	} else {
		switch os.Args[1] {
		case "uuid":
			fmt.Println(generate_uuid())
		case "cpf":
			fmt.Println(generateCPF())
		case "cnpj":
			fmt.Println(generateCNPJ())
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

func generateCPF() string {
	var cpf [11]int
	for i := 0; i < 9; i++ {
		cpf[i] = rand.Intn(10)
	}

	// Calcula os dois dígitos verificadores
	cpf[9] = calculateDigit(cpf[:9], 10)
	cpf[10] = calculateDigit(cpf[:10], 11)

	return fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d", cpf[0], cpf[1], cpf[2], cpf[3], cpf[4], cpf[5], cpf[6], cpf[7], cpf[8], cpf[9], cpf[10])
}

func generateCNPJ() string {
	var cnpj [14]int
	for i := 0; i < 12; i++ {
		cnpj[i] = rand.Intn(10)
	}
	cnpj[12] = calculateDigit(cnpj[:12], 5)
	cnpj[13] = calculateDigit(cnpj[:13], 6)

	return fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d%d%d", cnpj[0], cnpj[1], cnpj[2], cnpj[3], cnpj[4], cnpj[5], cnpj[6], cnpj[7], cnpj[8], cnpj[9], cnpj[10], cnpj[11], cnpj[12], cnpj[13])
}

func calculateDigit(base []int, weightStart int) int {
	sum, weight := 0, weightStart
	for _, num := range base {
		sum += num * weight
		weight--
		if weight < 2 {
			weight = 9
		}
	}
	digit := 11 - (sum % 11)
	if digit >= 10 {
		digit = 0
	}
	return digit
}
