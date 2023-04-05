package partyrobot

import (
    "fmt"
)

func Welcome(name string) string {
    welcomeMsg := "Welcome to my party, "
    return welcomeMsg + name + "!"
}

func HappyBirthday(name string, age int) string {
    hbdMsg := "Happy birthday " + name + "! "
    b := fmt.Sprintf("You are now %d years old!", age)
    return hbdMsg + b
}


func AssignTable(name string, table int, neighbor string, direction string, distance float64) string {
	paddedTable := fmt.Sprintf("%03d", table)
	distanceFormatted := fmt.Sprintf("%.1f", distance)

	message := Welcome(name) + "\n" +
		"You have been assigned to table " + paddedTable + ". Your table is " + direction + ", exactly " + distanceFormatted + " meters from here.\n" +
		"You will be sitting next to " + neighbor + "."

	return message
}

