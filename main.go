package main

import (
	"fmt"

	"github.com/ornlu-is/go_functional_options/person"
)

func main() {
	unknownPerson := person.New(1)
	aragorn := person.New(
		2,
		person.WithName("Aragorn II"),
		person.WithAddress("Rivendell"),
		person.WithAge(118),
		person.WithEmail("aragorn@mithrilmail.com"),
	)

	fmt.Printf("%+v\n", *unknownPerson)
	fmt.Printf("%+v\n", *aragorn)
}
