package main

import (
	"fmt"
	"reflect"

	"github.com/iRajesha/experiments/src/panic"
)

type Persons struct {
	name     string `required:"true"`
	age      int
	siblings []string
}

func main() {

	var stringType string
	fmt.Printf("%v,%T\n", stringType, stringType)

	mySlice := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("mySlice %v \n", mySlice[:])
	fmt.Printf("mySlice %v \n", mySlice[0:])
	fmt.Printf("mySlice %v \n", mySlice[1:])
	fmt.Printf("mySlice %v \n", mySlice[1:3])
	fmt.Printf("mySlice %v \n", mySlice[:3])

	myNewSlice := make([]int, 2)
	myNewSlice[0] = 123
	myNewSlice[1] = 111
	myNewSlice = append(myNewSlice, 123)
	fmt.Printf("myNewSlice %v", myNewSlice)
	myNewSlice = append(myNewSlice[:1], myNewSlice[2:]...)
	fmt.Printf("myNewSlice %v\n", myNewSlice)

	ray := Persons{
		name: "ray",
		age:  30,
		siblings: []string{
			"Chota Ray",
			"Bada Ray"}}
	fmt.Printf("%v\n", ray)

	rayMap := make(map[string]string)
	rayMap["name"] = "ray"
	fmt.Printf("%v\n", rayMap)
	rayRefectType := reflect.TypeOf(ray)
	rayRefectField, _ := rayRefectType.FieldByName("name")
	fmt.Println(rayRefectField.Tag.Get("required"))

	if _, ok := rayMap["name"]; ok {
		fmt.Printf("Value exists -- > %v\n", rayMap["name"])

	}

	playWithEmptyInterface(&ray)

	i := 20
	switch {
	case i <= 20:
		fmt.Printf("Less than 20\n")

	case i >= 20:
		fmt.Printf("greater than 20\n")
		fallthrough
	default:
		fmt.Printf("Printing default\n")
	}
backToFor:

	for i := 0; i < 5; {
		fmt.Println(i)
		break backToFor
	}

	fmt.Printf("Right after for loop\n")

	panic.CheckPanicInGo()
	i = 1

	str1 := "String"
	str1 = str1 + fmt.Sprintf("_%v", i)
	fmt.Printf("Concatinated value is %v", str1)

	extractInterfaceType(ray)

}

func playWithEmptyInterface(spreadedArgs interface{}) {
	fmt.Printf("Address -- %v\n", spreadedArgs)

	newName := spreadedArgs.(*Persons)

	fmt.Printf("Address -- %v\n", newName)
}

func extractInterfaceType(i interface{}) {

	switch i := i.(type) {
	case Persons:
		fmt.Printf("Found a person %v\n", i.name)
	}

}
