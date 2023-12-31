package main

import (
	"agency/src"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Println("\nCommand:")
		fmt.Println("1. list - List agencies in a region")
		fmt.Println("2. get - Get details of an agency by ID")
		fmt.Println("3. create - Create a new agency")
		fmt.Println("4. edit - Edit details of an agency by ID")
		fmt.Println("5. status - Get status of agencies in a region")
		fmt.Println("6. exit - Exit the program")

		fmt.Print("Enter Command: ")
		command := src.GetUserInput()

		switch command {
		case "exit":

			return

		case "list":
			fmt.Print("Enter region: ")
			region := src.GetUserInput()
			src.ListAgencies(region)

		case "get":
			fmt.Print("Enter Agency ID: ")
			agencyID, _ := strconv.Atoi(src.GetUserInput())
			src.GetAgencyDetail(agencyID)

		case "create":
			fmt.Print("Enter agency name: ")
			name := src.GetUserInput()
			fmt.Print("Enter region: ")
			region := src.GetUserInput()
			fmt.Print("Enter address: ")
			address := src.GetUserInput()
			fmt.Print("Enter phone: ")
			phone := src.GetUserInput()
			fmt.Print("Enter employeeCount: ")
			employeeCount := src.GetUserInput()
			src.CreateAgencies(name, region, address, phone, employeeCount)

		case "edit":
			fmt.Print("Enter agency ID to Edit:")
			agencyID, _ := strconv.Atoi(src.GetUserInput())
			src.EditAgency(agencyID)

		case "status":
			fmt.Print("Enter region: ")
			region := src.GetUserInput()
			src.Status(region)

		default:
			fmt.Println(strings.Repeat("-", 30))
			fmt.Print(strings.Repeat("\n", 1))
			fmt.Println("Invalid command. Please try again.")
			fmt.Print(strings.Repeat("\n", 1))
			fmt.Println(strings.Repeat("-", 30))
		}
	}
}
