package main

import (
	"fmt"

	"github.com/google/uuid"
)
type User struct {
	Customer string
	UUID string
}
func GenerateUUID() string {
	return uuid.New().String()
}
func PrintCustomer(arr *[]User,  ch chan bool ) {
	fmt.Println("Mulai Print Dari PrintCustomer")
	for _, v := range *arr {
		fmt.Printf("%+v\n", v)
	}
	ch <- true
}
func SliceAndGenerateUser(count int, arr *[]User, ch chan bool)  {
	for i := 0; i < count; i++ {
		*arr = append(*arr, User{fmt.Sprintf("Customer-%d",i), GenerateUUID()})
	}
	ch <- true
}
func main()  {
	var arr []User
	fmt.Println("Mulai Main Program", arr)
	slice := make(chan bool)
	print := make(chan bool)
	go SliceAndGenerateUser(20, &arr, slice)
	result := <-slice
	fmt.Println("Data Diterima?", result)
	
	// fmt.Println("Mulai Print Dari Main")
	// for _, v := range arr {
	// 	fmt.Printf("%+v\n", v)
	// }
	go PrintCustomer(&arr, print)
	<-print
}
