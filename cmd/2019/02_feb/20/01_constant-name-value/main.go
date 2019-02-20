package main

import "fmt"

type wordInt int16

const (
	// SUNDAY defines constant SUNDAY
	SUNDAY = wordInt(1)

	// MONDAY defines constant MONDAY
	MONDAY = wordInt(2)

	// TUESDAY defines constant TUESDAY
	TUESDAY = wordInt(3)

	// WEDNESDAY defines constant WEDNESDAY
	WEDNESDAY = wordInt(4)

	// THURSDAY defines constant THURSDAY
	THURSDAY = wordInt(5)

	// FRIDAY defines constant FRIDAY
	FRIDAY = wordInt(6)

	// SATURDAY defines constant SATURDAY
	SATURDAY = wordInt(7)
)

func (wi wordInt) String() string {
	switch wi {
	case 1:
		return "SUNDAY"
	case 2:
		return "MONDAY"
	case 3:
		return "TUESDAY"
	case 4:
		return "WEDNESDAY"
	case 5:
		return "THURSDAY"
	case 6:
		return "FRIDAY"
	case 7:
		return "SATURDAY"
	}
	return "NOT DEFINED"
}

func main() {
	fmt.Printf("%d : %s\n", SUNDAY, SUNDAY)
	fmt.Printf("%d : %s\n", MONDAY, MONDAY)
	fmt.Printf("%d : %s\n", TUESDAY, TUESDAY)
	fmt.Printf("%d : %s\n", WEDNESDAY, WEDNESDAY)
	fmt.Printf("%d : %s\n", THURSDAY, THURSDAY)
	fmt.Printf("%d : %s\n", FRIDAY, FRIDAY)
	fmt.Printf("%d : %s\n", SATURDAY, SATURDAY)
}
