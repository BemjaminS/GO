package main

import "fmt"

func main()  {

	confernceName := "Go confernce"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", confernceName)
	fmt.Printf("We have total of %v tickets and %v are remaining\n" ,conferenceTickets , remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

		//assign data to var using input
		// using & to store data in the pointer 
		var firstName string
		var lastName string
		var email string
		var userTicket uint

	 	fmt.Println("Enter your First name")
	 	fmt.Scan(&firstName)
	 	fmt.Println("Enter your Last name")
	 	fmt.Scan(&lastName)
	 	fmt.Println("Enter your Email adress")
	 	fmt.Scan(&email)
	 	fmt.Println("Enter number of tickets")
	 	fmt.Scan(&userTicket)
	
	 	remainingTickets = remainingTickets - userTicket
	 	fmt.Printf("Thank you %v %v for booking %v tickets. You will recive a confirmation email at %v\n" , firstName , lastName , userTicket , email)
	 	fmt.Printf("%v Remaining ticket to %v\n" , remainingTickets , confernceName)
	 	bookings = append(bookings , firstName + " " + lastName)
	 	fmt.Printf("booking list %v , \n", bookings)

	}
}