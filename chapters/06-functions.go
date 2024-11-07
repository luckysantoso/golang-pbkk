package main

import (
	"fmt"
	"strings"
)

const totalTickets int = 50

var availableTickets uint = 50
var eventTitle = "Go Conference"
var reservations = []string{}

func main() {

	greetUsers()

	for {

		fname, lname, email, ticketCount := getUserInput()
		nameIsValid, emailIsValid, ticketCountIsValid := validateUserInput(fname, lname, email, ticketCount)

		if nameIsValid && emailIsValid && ticketCountIsValid {

			bookTicket(ticketCount, fname, lname, email)

			fnames := printFirstNames()
			fmt.Printf("List of attendees' first names: %v\n", fnames)

			if availableTickets == 0 {
				// end program
				break
			}
		} else {
			if !nameIsValid {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !emailIsValid {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !ticketCountIsValid {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}
}

func printFirstNames() []string {
	fnames := []string{}

	for _, booking := range reservations {
		var names = strings.Fields(booking)
		fnames = append(fnames, names[0])
	}
	return fnames
}

func getUserInput() (string, string, string, uint) {
	var fname string
	var lname string
	var email string
	var ticketCount uint

	fmt.Println("Provide Your First Name: ")
	fmt.Scanln(&fname)

	fmt.Println("Provide Your Last Name: ")
	fmt.Scanln(&lname)

	fmt.Println("Provide Your Email Address: ")
	fmt.Scanln(&email)

	fmt.Println("How many tickets do you need?: ")
	fmt.Scanln(&ticketCount)

	return fname, lname, email, ticketCount
}

func validateUserInput(fname string, lname string, email string, ticketCount uint) (bool, bool, bool) {
	nameIsValid := len(fname) >= 2 && len(lname) >= 2
	emailIsValid := strings.Contains(email, "@")
	ticketCountIsValid := ticketCount > 0 && ticketCount <= availableTickets
	return nameIsValid, emailIsValid, ticketCountIsValid
}

func greetUsers() {
	fmt.Printf("Welcome to %v event registration system.\nWe have total of %v tickets and %v are still available.\nSecure your spot by booking tickets now\n", eventTitle, totalTickets, availableTickets)
}

func bookTicket(ticketCount uint, fname string, lname string, email string) {
	availableTickets = availableTickets - ticketCount
	reservations = append(reservations, fname+" "+lname)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", fname, lname, ticketCount, email)
	fmt.Printf("%v tickets remaining for %v\n", availableTickets, eventTitle)
}
