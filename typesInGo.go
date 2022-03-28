package main

import "fmt"

// Variables declaration
const caption string = "The following is the account information."
const firstName string = "Luke"
const lastName string = "Skywalker"
const ageInYears int8 = 20

var weightInKg float32 = 73.0
var heightInM float32 = 1.72
var remainingCreditsInDollars float64 = 123.55

const accountName string = "admin"

var accountPassword string = "password"
var isSubscribed bool = true

func tellMeTypes() {
	// print out the variable values and data types
	fmt.Println(caption)
	fmt.Println("=========================================")
	fmt.Printf("First name : %s\n", firstName)
	fmt.Printf("Data type is a %T\n", firstName)
	fmt.Println("----------")
	fmt.Printf("Last name : %s\n", lastName)
	fmt.Printf("Data type is a %T\n", lastName)
	fmt.Println("----------")
	fmt.Printf("Age : %d Years Old\n", ageInYears)
	fmt.Printf("Data type is a %T\n", ageInYears)
	fmt.Println("----------")
	fmt.Printf("Weight (kg) : %f\n", weightInKg)
	fmt.Printf("Data type is a %T\n", weightInKg)
	fmt.Println("----------")
	fmt.Printf("Height (m) : %f\n", heightInM)
	fmt.Printf("Data type is a %T\n", heightInM)
	fmt.Println("----------")
	fmt.Printf("Remining Credits ($) : %f\n", remainingCreditsInDollars)
	fmt.Printf("Data type is a %T\n", remainingCreditsInDollars)
	fmt.Println("----------")
	fmt.Printf("Account name : %s\n", accountName)
	fmt.Printf("Data type is a %T\n", accountName)
	fmt.Println("----------")
	fmt.Printf("Account password : %s\n", accountPassword)
	fmt.Printf("Data type is a %T\n", accountPassword)
	fmt.Println("----------")
	fmt.Printf("Subscribed user : %t\n", isSubscribed)
	fmt.Printf("Data type is a %T\n", isSubscribed)
	fmt.Println("----------")
}

func main() {
	tellMeTypes()
}
