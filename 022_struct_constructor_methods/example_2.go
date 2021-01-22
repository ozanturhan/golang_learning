package main

import (
	"fmt"
	"strconv"
)

func main() {
	// create user v1

	fmt.Println("Create User v1")

	user1 := &User{
		ID:        1,
		FirstName: "Ozan",
		LastName:  "Turhan",
		UserName:  "moturhan",
		Age:       29,
		Pay: &Payment{
			Salary: 5500,
			Bonus:  1000,
		},
	}

	fmt.Println(user1)
	fmt.Println(user1.GetFullName())
	fmt.Println(user1.GetPayment())
	fmt.Println("Maaş: " + strconv.FormatFloat(user1.GetPayment(), 'g', -1, 64))

	// create use v2
	fmt.Println("Create User v2")
	user2 := NewUser()
	user2.FirstName = "Ozan"
	user2.LastName = "Turhan"
	user2.Age = 28
	user2.UserName = "motruhan"
	user2.Pay = &Payment{Salary: 755, Bonus: 2500}

	fmt.Println(user2.GetPayment())
	fmt.Println(user2.GetFullName())

}

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

// User Constructor
func NewUser() *User {
	user := new(User)
	user.Pay = NewPayment()

	return user
}

type Payment struct {
	Salary float64
	Bonus  float64
}

// Payment Constructor
func NewPayment() *Payment {
	return new(Payment)
}

// Methods
func (user User) GetFullName() string {
	return user.FirstName + " " + user.LastName
}

func (user *User) GetUserName() string {
	return user.UserName
}

func (user *User) GetPayment() float64 {
	return user.Pay.Salary + user.Pay.Bonus
}
