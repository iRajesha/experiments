package main

import (
	"fmt"
)

type Persons struct {
	name string
	age  int
}

func main() {
	person := []Persons{{"Ray", 12}, {"Ray", 12}}
	fmt.Println("My first GO")
	personPtr := &person
	for _, eachPerperson := range person {
		fmt.Printf("Person type %T, pointer value is %v \n", eachPerperson, *personPtr)
	}

}
