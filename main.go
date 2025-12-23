package main

import (
	"fmt"

	"ADP/Hotel"
	"ADP/Employee"
	"ADP/Gym"
	"ADP/Wallet"
)

func main() {
	
	fmt.Println("\n--- Hotel ---")
	h := Hotel.NewHotel()
	h.AddRoom("101", "Single", 100)
	h.AddRoom("102", "Double", 150)

	h.CheckIn("101")
	h.ListVacantRooms()
	h.CheckOut("101")

	
	fmt.Println("\n--- Employees ---")
	employees := []Employee.Employee{
		Employee.FullTime{MonthlySalary: 300000, BonusRate: 0.1},
		Employee.PartTime{HourlyRate: 2000, HoursWorked: 80},

	}

	for i := 0; i < len(employees); i++ {
    fmt.Println("Salary:", employees[i].CalculateSalary())
    }



	
	fmt.Println("\n--- Gym ---")
	gym := Gym.NewGym()
	gym.AddMember(1, Gym.BasicMember{Name: "Nickole"})
	gym.AddMember(2, Gym.PremiumMember{Name: "Dane"})
	gym.ListMembers()

	fmt.Println("\n--- Wallet ---")
	w := Wallet.NewWallet()
	w.AddMoney(1000)
	w.SpendMoney(300)
	fmt.Println("Balance:", w.GetBalance())
}
