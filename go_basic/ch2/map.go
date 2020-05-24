package main

import "fmt"

type PersonmInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonmInfo
	personDB = make(map[string]PersonmInfo)

	personDB["12345"] = PersonmInfo{"12345", "Tom", "Room 203,..."}
	personDB["1"] = PersonmInfo{"1", "Jack", "Room101,..."}

	person, ok := personDB["1234"]

	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}
}
