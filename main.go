package main

import (
	"fmt"
	"strings"
)


func main(){
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50


	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Println("We have",conferenceTickets,"number of tickets and",remainingTickets,"are still available.")
	fmt.Println("Get your tickets here to attend.")

	var bookings []string

	for{
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("slice type: %T\n", bookings)
		fmt.Printf("slice length: %v\n", len(bookings))
		fmt.Printf("Thank you %v %v booked %v tickets. You will receive confirmation email at %v\n",firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for conference.\n",remainingTickets)

		firstNames := []string{}
		for _, booking := range bookings{
			//fields will separate the string on basis of space
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		

	}
	
	
	}
