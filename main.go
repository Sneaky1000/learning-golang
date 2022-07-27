package main // Turns the project into a package (required)

import (
	"fmt"                    // A library containing different usable functionalities in regards to inputs and outputs
	"learning_golang/common" // Other file within the folder
	"sync"                   // A library that adds features related to asynchronous programming
	"time"                   // A library that adds features related to time (ex. delays)
)

// This application will go over all the basics of Go
// Comments will be shown to explain syntax and features

// Variables can be set up using "var" or ":=" or "const"
// Scoped variables cannot be made on a package level
var conferenceName string = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint8 = 50

// Arrays in Go have a fixed size and data type (which must be defined as shown below—50 elements in this case)
// var bookings [50]string
// Can also be set up with: bookings := [50]string{}

// Slices are an abstraction of arrays in Go. It has a dynamic size (resizes when needed) and is index-based
// Simply remove the size from the brackets to turn the array into a slice (as shown below)
// Can also be set up with: bookings := []string{}
// Note: The code was edited to be a slice of maps (the 0 is the initial size of the slice): var bookings = make([]map[string]string, 0)
// Note: In the brackets (above), the key-type is stated. Outside the brackets to the left is the value-type
// Note: The code was then edited to be a slice of structs as shown below
var bookings = make([]userData, 0)

// Structs (structures) allow us to collect mixed data types unlike maps in Go. They are similar to classes in object-oriented programming languages
// The "type" keyword creates a new type. Here, we created a type called userData based on a struct of firstName, lastName, etc
type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint8
}

// WaitGroup{} waits for the launched goroutine to finish before terminating/stopping the program
var wg = sync.WaitGroup{}

func main() {

	// Calling functions in Go is the same as any other programming language
	greetUsers()

	// // Loops are simplified in Go. There is no for-each, while, or do-while loop. Just for-loop (with different types)
	// // You can put conditions before the start of the brackets in the for-loop just like an if/else statement
	// for {
	firstName, lastName, email, userTickets := getUserInput()

	// Multiple variables can be assigned to one function's return value(s) as shown below
	isValidName, isValidEmail, isValidTickets := common.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTickets {
		bookTicket(userTickets, firstName, lastName, email)
		// Adding "go" in front of a called function makes the function run asycronously (concurrent or alongside) the rest of the program
		// It turns this function below into a goroutine, which is a lightweight thread managed by the Go runtime
		// If this entire booking portion wasn't in a for-loop, this function below wouldn't run because the main goroutine will
		// stop once it's complete. It won't wait for this asyncronous function to run. It will simply exit and stop running beforehand
		// Additionally, that wg variable mentioned earlier will be used to fix this problem. Add() is used to add additional threads to be ran
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("These are all our bookings(first names only): %v\n", firstNames)

		// If no tickets remain, end the program
		// The ":=" automatically makes the variable below a "boolean" data type
		// It can be manually done using the following: var noRemainingTickets bool = remainingTickets == 0
		// This step is unnecessary but demonstrates how booleans can be used in Go
		noRemainingTickets := remainingTickets == 0
		if noRemainingTickets {
			fmt.Println("Our conference is fully booked. Come back next year.")
			// // Break ends the loop
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("The name you entered is too short.")
		}
		if !isValidEmail {
			fmt.Println("The email you entered is in an incorrect format.")
		}
		if !isValidTickets {
			fmt.Println("The number of tickets you entered is invalid.")
		}
		// // Continue will skip the rest of the loop and completely start the loop over from the beginning
		// // Continue isn't needed here because of the if/else statement
		// continue
	}
	// This makes sure the application will wait until all additional threads are done running before exiting
	wg.Wait()
}

// // The following code isn't needed in this application but demonstrates how switch statements work in Go
// city := "London"

// switch city {
// case "New York":
// 	// Execute code for booking New York conference tickets
// // Commas work as an "or" statement
// case "London", "Berlin":
// 	// Execute code for booking London or Berlin conference tickets
// case "California":
// 	// Execute code for booking California conference tickets
// case "Texas":
// 	// Execute code for booking Texas conference tickets
// // This works as an "else" sort of statement. It will default to this if none of the cases are true
// default:
// 	fmt.Println("No valid city selected")
// }

// }

// Functions work the same in Go as pretty much every other language
// When entering a parameter, make sure to define its data type
func greetUsers() {
	// Println prints on the next line whilst Print prints on the same line
	// Concatenation with variables can be done using the "+" operator or commas (auto-adds spaces)
	// \t can be used to add an indent (tab) and \n can be used to add a new line (enter)
	// Printf is used for printing formatted data (%v acts as a placeholder for a variable and is defined
	// after the string using a comma as shown below—the value of confName will replace %v)
	// %v = value (in a default format), %s = string, %T = type of value, %t = the word "true" or "false"
	// Listed above are some of the verbs. Link to all of them: https://pkg.go.dev/fmt
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("We have a total of %v tickets of which %v remain\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// Inside the parenthesis are the INPUT parameters. Outside the parenthesis are the OUTPUT parameters (can be put in separate parenthesis)
func getFirstNames() []string {
	firstNames := []string{}
	// This is a "for-each loop. "index" is replaced by an underscore because it's not used in this example
	// To ignore a variable you don't want to use, replace it with an underscore "_"
	for _, booking := range bookings {
		// strings.Fields separates each string by its white space into separate elements (requires the strings package)
		// Note: Replaced to simply take the first name value from the map created: booking["firstName"]
		// Note: Replaced to even more simply take the first name value from the struct created as shown below
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint8) {
	// All values in Go have data types (just like any other language)
	// Go is statically-typed, meaning you must tell the compiler the data type when declaring a variable
	// BUT Go can infer the type when you assign a value to that variable
	// It's important to know what data type to use and when. Link to all types used in go: https://www.tutorialspoint.com/go/go_data_types.htm
	// Constants cannot have their data type defined (ex. uint64)
	var firstName string
	var lastName string
	var email string
	var userTickets uint8

	// Scan can be used to get a user's input
	// A pointer must be used (which points to the memory address of a variable)
	// This shows the compiler where to store the user's input
	// An example is shown below (use "&" before a variable to show its pointer)
	// fmt.Println(&conferenceName) => prints: 0xc000010250 (this hash is the memory address for the conferenceName variable)
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint8, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	// // Creating a map in Go is simple. In the square brackets is the data type of the keys. On the right, outside the brackets, is the
	// // data type of the values. make() creates the map[] as shown below
	// userData := make(map[string]string)

	// This fills in the type (struct in this case) with the data received from the user's inputs
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// // This is the first key-value pair we are storing in the map
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// // Data types cannot be mixed in Go unlike most other languages so we need to convert uint8 to a string to match the rest of the values
	// userData["numberOfTickets"] = strconv.Itoa(int(userTickets))

	// // Adding an element by index (shown below)
	// bookings[0] = firstName + " " + lastName

	// append() adds elements to the end of a slice and grows the slice if need be
	bookings = append(bookings, userData)

	// // For testing
	// fmt.Printf("List of bookings is %v\n", bookings)

	// // Same syntax is used for arrays
	// fmt.Printf("The entire slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("The slice type: %T\n", bookings)
	// // len() gives the size/length of—in this case—the slice
	// fmt.Printf("The slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v ticket(s). You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint8, firstName string, lastName string, email string) {
	// Sleep freezes the application from continuining to run for a set amount of time (in this case, 5 seconds)
	time.Sleep(5 * time.Second)
	// Sprintf is used to store a string such as the one below in a variable
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#######################")
	fmt.Printf("Sending ticket %v to the email address %v\n", ticket, email)
	fmt.Println("#######################")
	// This lets the WaitGroup know that this function/additional thread is done running and removes it
	wg.Done()
}
