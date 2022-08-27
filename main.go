package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	"sync"
	
)
	const conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	var bookings = make([]UserData, 0)

	type UserData struct {
		firstName string
		lastName string
		email string
		userTickets uint
	
	}

	var wg = sync.WaitGroup{}

func main() {
	greetings()

		
		
		firstName, lastName, email, userTickets := getUserInputs()

		isValidName, isValidEmail, isValidTicketNum  := helper.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNum {
			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			if remainingTickets == 0 {
				fmt.Println("We are fully booked out!!!.")
				
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
	wg.Wait()
}

func greetings() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames() []string{
	var firstNames = []string{}
			for _, booking := range bookings{
			 
				firstNames = append(firstNames, booking.firstName)
			}

			return firstNames
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

func bookTicket( userTickets uint, firstName string, lastName string,email string) {
	   
	remainingTickets = remainingTickets - userTickets
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		userTickets: userTickets,
	}
	

	bookings = append(bookings,userData)
	fmt.Printf("Details of your booking: %v\n", bookings)

		firstNames := getFirstNames()

		fmt.Printf("Thank you %v  for booking %v ticket(s). You will receive a confirmation email at %v\n", firstName, userTickets,email)

		fmt.Printf("%v tickets, remaining for %v\n", remainingTickets, conferenceName)
		fmt.Println("list of first names of attendees: ", firstNames)
		
}

func sendTicket(userTickets uint, firstName string, lastName string,email string) {
  time.Sleep(10 * time.Second)
  var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
  fmt.Println("###############")
  fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
  fmt.Println("###############")
  wg.Done()


}

