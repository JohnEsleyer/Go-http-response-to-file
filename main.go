package main

import (
	"os"
	"strconv"
)

type Person struct {
	ID                      int    `json:"id"`
	UID                     int    `json:"uid"`
	Password                string `json:"password"`
	First_name              string `json:"first_name"`
	Last_name               string `json:"last_name"`
	Username                string `json:"username"`
	Email                   string `json:"email"`
	Avatar                  string `json:"avatar"`
	Gender                  string `json:"gender"`
	Phone_number            string `json:"phone_number"`
	Social_insurance_number string `json:"social_insurance_number"`
	Data_of_birth           string `json:"data_of_birth"`
	EmployeeInfo            struct {
		Title     string `json:"title"`
		Key_skill string `json:"key_skill"`
	} `json:"employment"`
	Address struct {
		City           string `json:"city"`
		Street_name    string `json:"street_name"`
		Street_address string `json:"street_address"`
		Zip_code       string `json:"zip_code"`
		State          string `json:"state"`
		Country        string `json:"country"`
		Coordinates    struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"coordinates"`
	} `json:"address"`
	Credit_card struct {
		cc_number string
	} `json:"credit_card"`
	Subscription struct {
		Plan           string `json:"plan"`
		Status         string `json:"status"`
		Payment_method string `json:"payment_method"`
		Term           string `json:"term"`
	} `json:"Subscription"`
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	// open file
	var stat os.FileInfo
	var file1 *os.File
	file1, err = os.OpenFile("output.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	defer file1.Close()

	// Get the data from endpoint many times
	for i := 0; i < n; i++ {

	}
}
