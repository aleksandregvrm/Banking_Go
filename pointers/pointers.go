package pointers

import "fmt"

func Pointer() {
	age := 32
	var agePointer *int
	agePointer = &age
	getAdultYears(agePointer)
	fmt.Println("Age:", age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
