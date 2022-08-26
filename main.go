package main

import (
	"fmt"
	"strings"
)

	const conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	var bookings []string

func main() {
	greetings()

	for {
		
		firstName, lastName, email, userTickets := getUserInputs()

		isValidName, isValidEmail, isValidTicketNum  := validateUserInputs(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNum {
			bookTicket( userTickets, firstName, lastName, email)

			if remainingTickets == 0 {
				fmt.Println("We are fully booked out!!!.")
				break
			}
			

		} else {
			if !isValidName {
			  fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
			  fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNum {
			  fmt.Println("number of tickets you entered is invalid")
			}
			
		}
	}
}

func greetings() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames(bookings []string) []string{
	firstNames := []string{}
			for _, value := range bookings{
			 var names  = strings.Fields(value)
				firstNames = append(firstNames, names[0])
			}

			return firstNames
}

func validateUserInputs(firstName string, lastName string,
	 email string, userTickets uint)(bool,bool,bool) {
		
	isValidName := len(firstName) >= 2 && len(lastName) >=2 
		isValidEmail := strings.Contains(email,"@")
		isValidTicketNum := userTickets > 0 && userTickets <= remainingTickets 

		return isValidName, isValidEmail, isValidTicketNum

}

func getUserInputs() (string, string, string, uint) {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		return firstName, lastName, email, userTickets

}

func bookTicket( userTickets uint, firstName string, lastName string, email string) {
		remainingTickets = remainingTickets - userTickets
				
		bookings = append(bookings,firstName + " " + lastName)

		firstNames := getFirstNames(bookings)
			
		fmt.Printf("Thank you %v  for booking %v ticket(s). You will receive a confirmation email at %v\n", firstName, userTickets,email)

		fmt.Printf("%v tickets, remaining for %v\n", remainingTickets, conferenceName)
		fmt.Println("list of first names of attendees: ", firstNames)

}