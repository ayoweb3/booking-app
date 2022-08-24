package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >=2 
		isValidEmail := strings.Contains(email,"@")
		isValidTicketNum := userTickets > 0 && userTickets <= remainingTickets 

		if isValidName && isValidEmail && isValidTicketNum {

			remainingTickets = remainingTickets - userTickets
			
			_ = append(bookings,firstName + " " + lastName)

				
			fmt.Printf("Thank you %v  for booking %v ticket(s). You will receive a confirmation email at %v", firstName, userTickets,email)

			fmt.Printf("%v tickets, remaining for %v\n", remainingTickets, conferenceName)
			// firstNames := []string{}
			// for _, value := range bookings{
			//  var names  = strings.Fields(value)
			// 	firstNames = append(firstNames, names[0])
			// }

			// fmt.Printf("The first names of bookings are: %v\n", firstNames)

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