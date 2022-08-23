package main

import (
	"fmt"
)


func main() {
   const conferenceName = "Go Conference"
   const conferenceTickets = 50
   var remainingTickets = conferenceTickets
   

   fmt.Printf("Welcome to %v booking application\n", conferenceName)
   fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
   fmt.Printf("Get your tickets here to attend\n")

   var firstName string
   var lastName string
   var email string
   var userTickets int

   fmt.Println("Enter your first name: ")
   fmt.Scan(&firstName)

   fmt.Println("Enter your last name: ")
   fmt.Scan(&lastName)

   fmt.Println("Enter your email address: ")
   fmt.Scan(&email)

   fmt.Println("Enter number of tickets: ")
   fmt.Scan(&userTickets)
    
   fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)

   

}