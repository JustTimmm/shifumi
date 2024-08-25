package main

import "fmt"

func main() {
	var (
		age     int
		name    string
		student bool
	)

	age = 15
	name = "Timéo"
	student = true

	fmt.Printf("Bonjour je suis %v ", name)
	fmt.Printf("j'ai %vans ", age)

	if student == true {
		fmt.Printf("et je suis actuellement étudiant.")
	} else {
		fmt.Printf("et je ne suis pas étudiant.")
	}
}
