package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name    string
	Address *Address
}

type Address struct {
	City    string
	Country string
}

func CheckStuffs() {

	name := "Prashant"
	fmt.Printf("In caller func, name : ", name)
	fmt.Println()
	changeString(&name)
	fmt.Printf("After caller func... name : ", name)
	fmt.Println()

	nameSlice := []string{"Backspace", "Enter", "Shift"}
	fmt.Printf("In caller func, nameSlice : ", nameSlice)
	fmt.Println()
	changeSlice(nameSlice)
	fmt.Printf("After caller func.. nameSlice : ", nameSlice)
	fmt.Println()

	var nameInt int
	nameInt = 5
	fmt.Printf("In caller func, Name in int is : ", nameInt)
	fmt.Println()
	changeInt(&nameInt)
	fmt.Printf("After caller func.. Name in int is : ", nameInt)
	fmt.Println()

	u := User{}
	u.Name = "Prashant"
	u.Address = &Address{City: "Mum"}
	fmt.Printf("In caller func, Name in struct is : ", u.Name)
	fmt.Println()
	changeStruct(u)
	fmt.Printf("After caller func.. Name in struct is : %s, City is: %s", u.Name, u.Address.City)
	fmt.Println()
}

const (
	NAME = "ABC"
)

func changeString(name *string) {
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(name)
	*name = "Prashant Singh"
	fmt.Printf("In called func, name : %s", *name)
	fmt.Println()
}

func changeInt(name *int) {
	fmt.Println(name)
	*name = 2
	fmt.Printf("In called func, name : %d", *name)
	fmt.Println()
}

func changeSlice(nameSlice []string) {

	nameSlice[0] = "Backspace key"
	fmt.Printf("In called func, nameSlice : ", nameSlice)
	fmt.Println()
}
func changeStruct(u User) {

	fmt.Println(reflect.TypeOf(u))
	fmt.Printf("%+v", u)
	u.Name = "Prashant Singh"
	u.Address = &Address{City: "DEL"}
	//u.Address.City = "DEL"
	fmt.Printf("In called func, Name in struct is : %s, City is: %s", u.Name, u.Address.City)
	fmt.Println()
}

func main() {

	CheckStuffs()
}
