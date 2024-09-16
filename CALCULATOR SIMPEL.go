package main

import (
	"fmt"
)

func main() {

	//Declare
	var number1 float64
	var number2 float64
	var result float64
	var logic string

	//Text
	fmt.Println("Yuk Berhitung!")
	fmt.Println("Masukin Bilangan Dulu YuK")

	//Operasi
	fmt.Println("Mau Pakai Operasi Apa? (*, +, /, - ) = ")
	fmt.Scanf("%s", &logic)

	//Input 1
	fmt.Println("Masukan Bilangan 1")
	if _, err := fmt.Scanf("%f", &number1); err != nil {
		fmt.Println("Kalkulator Tidak Mengerti Input User")
		return

	}

	//Input 2
	fmt.Println("Masukan Bilangan 2")
	if _, err := fmt.Scanf("%f", &number2); err != nil {

		fmt.Println("Kalkulator Tidak Mengerti Input User")
		return

	}

	//Full Operation
	if logic == "*" {
		result = number1 * number2
		fmt.Println("Hasilnya Adalah = ", result)

	} else if logic == "+" {
		result = number1 + number2
		fmt.Println("Hasilnya Adalah = ", result)

	} else if logic == "/" {
		if number1 == 0 {
			fmt.Println("Kalkulator Tidak Mengerti Input User")
		}
		result = number1 / number2
		fmt.Println("Hasilnya Adalah = ", result)

	} else if logic == "-" {
		if number1 == 0 {
			fmt.Println("Kalkulator Tidak Mengerti Input User")
		}
		result = number1 - number2
		fmt.Println("Hasilnya Adalah = ", result)

	} else {
		fmt.Println("Input Operasi Salah")
		return

	}

}
