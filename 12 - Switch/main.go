package main

import "fmt"

func dayOfWeek(num uint) string {
	switch num {
	case 0:
		return "Domingo"
	case 1:
		return "Segunda-Feira"
	case 2:
		return "Terça-Feira"
	case 3:
		return "Quarta-Feira"
	case 4:
		return "Quinta-Feira"
	case 5:
		return "Sexta-Feira"
	case 6:
		return "Sábado"
	default:
		return "Invalid number day"
	}
}

func dayOfWeek2(num uint) string {
	switch {
	case num == 0:
		return "Domingo"
	case num == 1:
		return "Segunda-Feira"
	case num == 2:
		return "Terça-Feira"
	case num == 3:
		return "Quarta-Feira"
	case num == 4:
		return "Quinta-Feira"
	case num == 5:
		return "Sexta-Feira"
	case num == 6:
		return "Sábado"
	default:
		return "Invalid number day"
	}
}

func dayOfWeek3(num uint) string {
	switch num * 2 {
	case 0:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 4:
		return "Terça-Feira"
	case 6:
		return "Quarta-Feira"
	case 8:
		return "Quinta-Feira"
	case 10:
		return "Sexta-Feira"
	case 12:
		return "Sábado"
	default:
		return "Invalid number day"
	}
}

func main() {
	fmt.Println(dayOfWeek(12))
	fmt.Println(dayOfWeek2(4))
	fmt.Println(dayOfWeek3(4))
}
