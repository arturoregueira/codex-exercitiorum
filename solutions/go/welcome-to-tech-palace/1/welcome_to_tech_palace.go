package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	base_welcome := "Welcome to the Tech Palace,"
	cus_text := strings.ToUpper(customer)
	return base_welcome + " " + cus_text
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	return stars + "\n" + welcomeMsg + "\n" + stars
	//panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	text := strings.Trim(oldMsg, "*")
	text = strings.ReplaceAll(text, "*", "")
	text = strings.TrimSpace(text)
	fmt.Println(text)
	return text
	//panic("Please implement the CleanupMessage() function")
}
