package Employee

type Employee interface {
	CalculateSalary() float64
}

type FullTime struct {
	MonthlySalary float64
	BonusRate     float64
}

func (f FullTime) CalculateSalary() float64 {
	return f.MonthlySalary + f.MonthlySalary*f.BonusRate
}

type PartTime struct {
	HourlyRate   float64
	HoursWorked int
}

func (p PartTime) CalculateSalary() float64 {
	return p.HourlyRate * float64(p.HoursWorked)
}
