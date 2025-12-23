package main

import (
	"fmt"
	"os"

	"ADP/Employee"
	"ADP/Gym"
	"ADP/Hotel"
	"ADP/Wallet"
)

func main() {
	for {
		fmt.Println("\nMAIN MENU")
		fmt.Println("1. Hotel")
		fmt.Println("2. Employees")
		fmt.Println("3. Gym")
		fmt.Println("4. Wallet")
		fmt.Println("5. Exit")
		fmt.Print("Choice: ")

		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "1":
			runHotel()
		case "2":
			runEmployees()
		case "3":
			runGym()
		case "4":
			runWallet()
		case "5":
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Wrong choice")
		}
	}
}


func runHotel() {
	hotel := Hotel.NewHotel()

	for {
		fmt.Println("\nHOTEL MENU")
		fmt.Println("1. Add room")
		fmt.Println("2. Check in")
		fmt.Println("3. Check out")
		fmt.Println("4. Show vacant rooms")
		fmt.Println("5. Back")
		fmt.Print("Choice: ")

		var cmd string
		fmt.Scan(&cmd)

		switch cmd {
		case "1":
			var room Hotel.Room

			fmt.Print("Room number: ")
			fmt.Scan(&room.RoomNumber)

			fmt.Print("Room type: ")
			fmt.Scan(&room.Type)

			fmt.Print("Price per night: ")
			fmt.Scan(&room.PricePerNight)

			room.IsOccupied = false
			hotel.AddRoom(room)

		case "2":
			var num string
			fmt.Print("Room number: ")
			fmt.Scan(&num)
			hotel.CheckIn(num)

		case "3":
			var num string
			fmt.Print("Room number: ")
			fmt.Scan(&num)
			hotel.CheckOut(num)

		case "4":
			hotel.ListVacantRooms()

		case "5":
			return

		default:
			fmt.Println("Wrong input")
		}
	}
}


func runEmployees() {
	var employees []Employee.Employee

	for {
		fmt.Println("\nEMPLOYEES MENU")
		fmt.Println("1. Add full-time employee")
		fmt.Println("2. Add part-time employee")
		fmt.Println("3. Show salaries")
		fmt.Println("4. Back")
		fmt.Print("Choice: ")

		var cmd string
		fmt.Scan(&cmd)

		switch cmd {
		case "1":
			var emp Employee.FullTime
			fmt.Print("Monthly salary: ")
			fmt.Scan(&emp.MonthlySalary)

			fmt.Print("Bonus rate: ")
			fmt.Scan(&emp.BonusRate)

			employees = append(employees, emp)

		case "2":
			var emp Employee.PartTime
			fmt.Print("Hourly rate: ")
			fmt.Scan(&emp.HourlyRate)

			fmt.Print("Hours worked: ")
			fmt.Scan(&emp.HoursWorked)

			employees = append(employees, emp)

		case "3":
			for i := 0; i < len(employees); i++ {
				fmt.Println("Employee", i+1, "salary:",
					employees[i].CalculateSalary())
			}

		case "4":
			return

		default:
			fmt.Println("Wrong input")
		}
	}
}


func runGym() {
	gym := Gym.NewGym()
	id := uint64(1)

	for {
		fmt.Println("\nGYM MENU")
		fmt.Println("1. Add basic member")
		fmt.Println("2. Add premium member")
		fmt.Println("3. List members")
		fmt.Println("4. Back")
		fmt.Print("Choice: ")

		var cmd string
		fmt.Scan(&cmd)

		switch cmd {
		case "1":
			var m Gym.BasicMember
			fmt.Print("Name: ")
			fmt.Scan(&m.Name)

			gym.AddMember(id, m)
			id++

		case "2":
			var m Gym.PremiumMember
			fmt.Print("Name: ")
			fmt.Scan(&m.Name)

			gym.AddMember(id, m)
			id++

		case "3":
			gym.ListMembers()

		case "4":
			return

		default:
			fmt.Println("Wrong input")
		}
	}
}


func runWallet() {
	w := Wallet.NewWallet()

	for {
		fmt.Println("\nWALLET MENU")
		fmt.Println("Balance:", w.GetBalance())
		fmt.Println("1. Add money")
		fmt.Println("2. Spend money")
		fmt.Println("3. Back")
		fmt.Print("Choice: ")

		var cmd string
		fmt.Scan(&cmd)

		switch cmd {
		case "1":
			var amount float64
			fmt.Print("Amount: ")
			fmt.Scan(&amount)
			w.AddMoney(amount)

		case "2":
			var amount float64
			fmt.Print("Amount: ")
			fmt.Scan(&amount)
			w.SpendMoney(amount)

		case "3":
			return

		default:
			fmt.Println("Wrong input")
		}
	}
}
