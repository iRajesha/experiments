package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"

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

	checkMarshalling()
	fmt.Printf("\n\n\n\n\n")
	gernateBusinessObjectId()
	//var empty interface{}
	bMarshalled, _ := json.Marshal("CCSNonProd")
	fmt.Printf("Marshhed string(byte) %v\n", string(bMarshalled))
	bytes := []byte("CCSNonProd")
	//json.Unmarshal([]byte(IkNDU05vblByb2Qi), empty)
	//fmt.Printf("%v", empty)
	fmt.Printf("CCSNonProd value -> %v", string(bytes))

}

func gernateBusinessObjectId() {
	var channelId = "0"
	now := time.Now()
	year := now.Year()
	lastTwoDigitsOfYear := year % 1e2
	_, week := now.UTC().ISOWeek()

	businessId := channelId + strconv.Itoa(lastTwoDigitsOfYear) + strconv.Itoa(week) + fmt.Sprintf("%d", get13MiddleDigits("4b537cebb9f67d9bcfe29d38d03d76febe54808a16238815e5ad28b695c197b7"))

	fmt.Printf("Last two digits --> %v\n", lastTwoDigitsOfYear)
	fmt.Printf("Final Business Identifier %v\n", businessId)
}

func get13MiddleDigits(txId string) int64 {
	//hasher := sha256.New()

	bv := []byte(txId)
	result := sha256.Sum256(bv)
	hexString := hex.EncodeToString(result[:])
	last13HexDigits := hexString[50 : len(hexString)-1]
	fmt.Printf(" last13HexDigits ->  %v\n", last13HexDigits)
	decimal, err := strconv.ParseInt(last13HexDigits, 16, 64)
	if err != nil {
		fmt.Printf("Error %v\n ", err)
	}
	last13digits := decimal % 1e13
	fmt.Printf("random 13 digits %v\n", last13digits)

	return last13digits
}

func checkMarshalling() {
	var walletRespPayload = struct {
		MspId  string `json:"mspId"`
		Status string `json:"status"`
	}{}

	type walletRespPayloadStruct struct {
		Username string `json:"username"`
		Secret   string `json:"secret"`
		MspId    string `json:"mspId"`
		Status   string `json:"status"`
	}
	structVar := walletRespPayloadStruct{}
	//payload := "{\"username\":\"NewOrgTMAId\",\"secret\":\"63c1dd951ffedf6f7fd968ad4efa39b8ed584f162f46e715114ee184f8de9201\",\"mspId\":\"TMAFounderTEST\",\"status\":\"TMA Member\"}"
	payload := string(`{"username":"NewOrgTMAId","secret":"63c1dd951ffedf6f7fd968ad4efa39b8ed584f162f46e715114ee184f8de9201","mspId":"TMAFounderTEST","status":"TMA Member"}`)
	bPayload, _ := json.Marshal(payload)

	fmt.Printf("Marshalled payload --> %v\n", string(bPayload))
	json.Unmarshal([]byte(payload), &walletRespPayload)
	fmt.Printf("Unmashalled Payload %v\n", walletRespPayload)
	json.Unmarshal([]byte(payload), &structVar)
	fmt.Printf("Unmashalled Payload %v\n", structVar)

	//value := string(`{"TMAFounderTEST":"Hello"}`)
	value, err := json.Marshal("TMAFounderTEST")
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	fmt.Printf("Value -- %v\n", string(value))
	var obj interface{}
	err = json.Unmarshal([]byte(value), &obj)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

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
