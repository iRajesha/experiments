package panic

import "fmt"

func CheckPanicInGo() {
	fmt.Printf("Inside CheckPanitInGo function\n")

	raiseAndHandlePanic()
	fmt.Println("Statement after Panic")

}

func raiseAndHandlePanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic Handelled \n Error is %v\n", err)
		}
	}()
	panic("Panic Raised")

}
