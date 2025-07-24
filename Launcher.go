package main

import (
	"HospitalQOps/logic/doctor"
	"HospitalQOps/logic/headdoctor"
	"HospitalQOps/logic/patient"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=====================================")
	fmt.Println("DoctorOperations Terminal System")
	fmt.Println("Manage doctor-patient operations in a hospital")
	fmt.Println("Technologies: Go + TCP + PostgreSQL")
	fmt.Println("=====================================")
	fmt.Println("Choose your role:")
	fmt.Println("1. Doctor")
	fmt.Println("2. Patient")
	fmt.Println("3. Head Doctor")
	fmt.Print("Enter choice [1-3]: ")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		doctor.Main()
	case "2":
		patient.Main()
	case "3":
		headdoctor.Main()
	default:
		fmt.Println("Invalid choice. Exiting.")
	}
}
