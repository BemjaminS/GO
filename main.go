package main

import "fmt"

func main()  {

	confernceName := "Go confernce"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", confernceName)
	fmt.Printf("We have total of %v tickets and %v are remaining\n" ,conferenceTickets , remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets uint
	//assign data to var
	userName = "Ben"
	userTickets = 2 
	fmt.Printf("User %v booked %v Tickets\n" , userName , userTickets)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("%v Remaining ticket to %v\n" , remainingTickets , confernceName)

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

	 //using array
	 var bookings [50]string
	 bookings[0] = firstName + " " + lastName
	 //using append to dynamic array
	 //bookings = append(bookings , firstName + " " + lastName)
	 fmt.Printf("Print whole array: %v\n", bookings)
	 fmt.Printf("Print First index: %v\n", bookings[0])
	 fmt.Printf("Print Type %T\n", bookings)
	 fmt.Printf("Print array lenght: %v\n", len(bookings))

	 




}