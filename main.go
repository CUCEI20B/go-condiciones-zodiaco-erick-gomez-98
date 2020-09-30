package main

import "fmt"

func main() {
	var dia uint64
	var mes uint64

	fmt.Scanln(&dia)
	fmt.Scanln(&mes)

	switch mes {
	case 1:
		if dia >= 1 && dia <= 20 {
			fmt.Println("capricornio")
		} else {
			fmt.Println("acuario")
		}
	case 2:
		if dia >= 0 && dia <= 19 {
			fmt.Println("acuario")
		} else {
			fmt.Println("piscis")
		}

	case 3:
		if dia >= 0 && dia <= 20 {
			fmt.Println("piscis")
		} else {
			fmt.Println("aries")
		}
	case 4:
		if dia >= 0 && dia <= 20 {
			fmt.Println("aries")
		} else {
			fmt.Println("tauro")
		}

	case 5:
		if dia >= 0 && dia <= 21 {
			fmt.Println("tauro")
		} else {
			fmt.Println("geminis")
		}
	case 6:
		if dia >= 0 && dia <= 21 {
			fmt.Println("geminis")
		} else {
			fmt.Println("cancer")
		}
	case 7:
		if dia >= 0 && dia <= 23 {
			fmt.Println("cancer")
		} else {
			fmt.Println("leo")
		}
	case 8:
		if dia >= 0 && dia <= 23 {
			fmt.Println("leo")
		} else {
			fmt.Println("virgo")
		}
	case 9:
		if dia >= 0 && dia <= 23 {
			fmt.Println("virgo")
		} else {
			fmt.Println("libra")
		}
	case 10:
		if dia >= 0 && dia <= 23 {
			fmt.Println("libra")
		} else {
			fmt.Println("escorpio")
		}
	case 11:
		if dia >= 0 && dia <= 22 {
			fmt.Println("escorpio")
		} else {
			fmt.Println("sagitario")
		}
	case 12:
		if dia >= 0 && dia <= 21 {
			fmt.Println("sagitario")
		} else {
			fmt.Println("capricornio")
		}
	}
}
