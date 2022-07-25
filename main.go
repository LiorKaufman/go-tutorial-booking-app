package main

import (
	"fmt"
	"strings"
)

func main() {

	const conferenceName = "GolangConf"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	for {

		fmt.Println("Welcome to", conferenceName, "booking application")
		fmt.Println("We have a total of", conferenceTickets, "tickets and we have", remainingTickets, "remaining")
		var userName string
		var userTickets uint

		fmt.Println("Please enter your name")
		fmt.Scan(&userName)
		fmt.Printf("Hello %s, how many tickets would you like to purchase?\n", userName)
		fmt.Scan(&userTickets)
		fmt.Printf("User %v has purchased %v tickets\n", userName, userTickets)
		remainingTickets = remainingTickets - userTickets
		fmt.Println("We have a total of", conferenceTickets, "tickets and we have", remainingTickets, "remaining")

		bookings = append(bookings, userName)

		if remainingTickets == 0 {
			fmt.Println("Sorry, we are sold out")
			break
		}
	}
	namesWithNoSpecialChars := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		namesWithNoSpecialChars = append(namesWithNoSpecialChars, names[0])
	}
	fmt.Println("The following names have been booked: ", namesWithNoSpecialChars)
}
