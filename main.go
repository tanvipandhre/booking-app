package main

import (
	"fmt"	
	"booking-app/helper"
	"strconv"
)
//Package level variables
var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
//empty list of maps []
var bookings = make([]map[string]string, 0)



func main(){
	
	greetUsers()
	
	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := userInput()
		
		isValidName, isValidEmail, isValidTicketNumeber := helper.ValidateUserInput(firstName , lastName , email , userTickets, remainingTickets  )
		
		
		if isValidName && isValidEmail && isValidTicketNumeber{
			
			bookTickets(userTickets , firstName , lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			
			if remainingTickets == 0{
				//end program
				fmt.Println("Our conference is booked out. Please come back next year.")
				break
			}

		}else{
			if !isValidName{
				fmt.Println("FirstNmae or LastName you entered is too short")
			}
			if !isValidEmail{
				fmt.Println("Email you entered does not have @ sign")
			}
			if isValidTicketNumeber{
				fmt.Printf("We only have %v tickets. Please enter valid number of tickets.\n", remainingTickets)
			}
			
			//we use continue so that it go to the next iteration
			continue
		}

		

		
		
	}
	
	
	}

	func greetUsers() {
		fmt.Printf("Welcome to %v booking application.\n", conferenceName)
		fmt.Println("We have",conferenceTickets,"number of tickets and",remainingTickets,"are still available.")
		fmt.Println("Get your tickets here to attend.")
	}

	func getFirstNames() []string{
		//slice firstnames	
		firstNames := []string{}
		//for each loop
		for _, booking := range bookings{
			//fields will separate the string on basis of space
			
			firstNames = append(firstNames, booking["firstName"])
		}
		return firstNames
		
		
	}

	

	func userInput() ( string, string, string, uint){
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

		return firstName, lastName, email, userTickets

	}

	func bookTickets(userTickets uint, firstName string, lastName string, email string) {
			remainingTickets = remainingTickets - userTickets
			//create a map for a user
			
			//empty map
			var userData = make(map[string]string)
			userData["firstName"] = firstName
			userData["lastName"] = lastName
			userData["email"] = email
			userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

			bookings = append(bookings, userData)	
			fmt.Printf("List of bookings is %v \n", bookings)		
			fmt.Printf("Thank you %v %v booked %v tickets. You will receive confirmation email at %v\n",firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for conference.\n",remainingTickets)
			
			
	}
