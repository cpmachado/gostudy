package partyrobot

import "fmt"

const welcomeFormat = "Welcome to my party, %s!"
const happyBirthdayFormat = "Happy birthday %s! You are now %d years old!"
const assignTableFormat = "Welcome to my party, %s!\n" +
	"You have been assigned to table %03d." +
	" Your table is %s, exactly %.1f meters from here.\n" +
	"You will be sitting next to %s."

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf(welcomeFormat, name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf(happyBirthdayFormat, name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf(assignTableFormat, name, table, direction, distance, neighbor)
}
