package doctor

import (
	"HospitalQOps/database"
	"HospitalQOps/errorspacket"
	"HospitalQOps/logger"
	"HospitalQOps/model"
	"HospitalQOps/network"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)

func AppointmentsManager(doc *model.Doctor) {
	fmt.Println("Appointments Manager Service")
	fmt.Println("Options:")
	fmt.Println("1. Appointments List")
	fmt.Println("2. Schedule appointment")
	fmt.Println("BACK")
	inputApp,err:=reader.ReadString('\n')
	if err!=nil{
		fmt.Println(errorspacket.InvalidInputError)
		return
	}
	inputApp=strings.TrimSpace(inputApp)
	inputAppU:=strings.ToUpper(inputApp)
	switch inputAppU{
	case "1":
		fmt.Println("List of appointments:")
		ListApp,err:=database.GetAppointments()
		if err!=nil{
			fmt.Println("Error getting appointments list")
			return
		}
		fmt.Println(ListApp)
	case "2":
		fmt.Println("Type the appointment ID to validate:")
		appointmentInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		appointmentInput = strings.TrimSpace(appointmentInput)
		appointmentID, err := strconv.Atoi(appointmentInput)
		if err != nil {
			fmt.Println("Invalid appointment ID format")
			return
		}

		fmt.Println("Doctor response [approved/denied]:")
		responseInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		responseInput = strings.TrimSpace(strings.ToLower(responseInput))

		err = database.DoctorValidatesAppointment(doc.ID, appointmentID, responseInput)
		if err != nil {
			fmt.Println("Failed to validate appointment")
		} else {
			fmt.Println("Appointment validation processed.")
		}
	case "BACK":
		fmt.Println("Returning to menu")
	default:
		fmt.Println("Choose a viable option")
	}
}
func HospitalEmailCenter(doc *model.Doctor, connection net.Conn) {
	fmt.Println("Welcome to the Hospital Email Center")
	fmt.Printf("List of Patients for %s %s\n", doc.LastName, doc.FirstName)
	patientsListing, err := database.GetPatientsListByDoctor(*doc)
	if err != nil {
		fmt.Println("Error retrieving patients")
		return
	}
	fmt.Println(patientsListing)
	fmt.Println("Options:")
	fmt.Println("1. Message")
	fmt.Println("2. Inbox")
	fmt.Println("3. Conversations")
	fmt.Println("BACK")
	var message model.Message
	input5, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(errorspacket.InvalidInputError)
		return
	}
	input5 = strings.TrimSpace(input5)
	input5U := strings.ToUpper(input5)
	switch input5U {
	case "1":
		fmt.Println("Type the individuals Email here: ")
		emailInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		emailInput = strings.TrimSpace(emailInput)
		fmt.Println("Title here: ")
		titleInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		titleInput = strings.TrimSpace(titleInput)
		fmt.Println("Message here: ")
		messageInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		messageInput = strings.TrimSpace(messageInput)
		message = model.Message{
			Sender:   doc.Email,
			Receiver: emailInput,
			Title:    titleInput,
			Content:  messageInput,
		}
		err = network.SendJSON(connection, message)
		if err != nil {
			fmt.Println("Error inserting message:", err)
		}
		err = database.InsertMessage(message)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Message sent to %s\n", emailInput)
	case "2":
		fmt.Println("Inbox:")
		InboxList, err := database.GetDoctorInbox(*doc)
		if err != nil {
			fmt.Println("Error retrieving doctors inbox")
			return
		}
		for i, msg := range InboxList {
			fmt.Printf("[%d] Sender: %s | Title: %s\n", i+1, msg.Sender, msg.Title)
		}
	case "3":
		fmt.Println("Type the individuals Email here to view the conversation: ")
		EmailInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		EmailInput = strings.TrimSpace(EmailInput)
		conversation, err := database.GetConversationEmail(*doc, EmailInput)
		for i, msg := range conversation {
			fmt.Printf("\nIndex [%d]: Sender %s Receiver %s Title %s\n", i+1, msg.Sender, msg.Receiver, msg.Title)
			fmt.Printf("Message: %s", msg.Content)
		}
	case "BACK":
	default:
		fmt.Println("Choose a viable option")
	}
}
func DoctorOperationsMenu(doc *model.Doctor, connection net.Conn) bool{
	fmt.Printf("Welcome back Doctor %s %s!\n", doc.LastName, doc.FirstName)
	fmt.Println("Options:")
	fmt.Println("1.Departments")
	fmt.Println("2.Patients")
	fmt.Println("3.Medications")
	fmt.Println("4.Prescriptions")
	fmt.Println("5.Appointments")
	fmt.Println("6.Examinations")
	fmt.Println("7.Admissions")
	fmt.Println("EXIT")
	input4, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(errorspacket.InvalidInputError)
		return false
	}
	input4 = strings.TrimSpace(input4)
	input4U := strings.ToUpper(input4)
	switch input4U {
	case "1":
		departmentsListing, err := database.GetNursesAndDoctorsByDepartments()
		if err != nil {
			fmt.Println("Error retrieving department data.", err)
		}
		fmt.Println(departmentsListing)
		fmt.Println("Returning to menu...")
	case "2":
		HospitalEmailCenter(doc, connection)
	case "3":
		medicationsListing, err := database.GetMedications()
		if err != nil {
			fmt.Println("Error retrieving medications list")
			return false
		}
		fmt.Println(medicationsListing)
		fmt.Println("Returning to menu...")
	case "4":
		fmt.Println("List of prescriptions:")
		ListPrescriptions,err:=database.GetPrescriptions()
		if err!=nil{
			fmt.Println(err)
			return false
		}
		fmt.Println(ListPrescriptions)
		fmt.Println("Choose an option:")
		fmt.Println("1.Insert a prescription for a patient after the examination")
		fmt.Println("2.Delete a prescription")
		fmt.Println("3.Visualize the average number of medications prescribed")
		fmt.Println("BACK")
		inputP,err:=reader.ReadString('\n')
		if err != nil {
			fmt.Println(errorspacket.InvalidInputError)
			return false
		}
		inputP=strings.TrimSpace(inputP)
		inputPU:=strings.ToUpper(inputP)
		switch inputPU{
		case "1":
			fmt.Println("Type the examination ID: ")
			inputPID,err:=reader.ReadString('\n')
			if err != nil {
				fmt.Println(errorspacket.InvalidInputError)
				return false
			}
			inputPID=strings.TrimSpace(inputPID)
			fmt.Println("Notes for medications dosage: ")
			inputPNot,err:=reader.ReadString('\n')
			if err != nil {
				fmt.Println(errorspacket.InvalidInputError)
				return false
			}
			inputPNot=strings.TrimSpace(inputPNot)
			fmt.Println("Date Format [20YY-MM-DD 00:00:00]: ")
			inputPD,err:=reader.ReadString('\n')
			if err != nil {
				fmt.Println(errorspacket.InvalidInputError)
				return false
			}
			inputPD=strings.TrimSpace(inputPD)
			fmt.Println("Price: ")
			inputPPr,err:=reader.ReadString('\n')
			if err != nil {
				fmt.Println(errorspacket.InvalidInputError)
				return false
			}
			inputPPr=strings.TrimSpace(inputPPr)
			err=database.InsertPrescription(inputPID, inputPNot, inputPD, inputPPr)
			if err!=nil{
				fmt.Println("Error inserting prescription")
				return false
			}
			LastPrescriptionID,err:=database.GetLastPrescription()
			if err!=nil{
				fmt.Println("Error getting last prescription")
				return false
			}
			fmt.Println("The 3 medications ID's: ")
			med1,err:=reader.ReadString('\n')
			if err != nil {
				fmt.Println(errorspacket.InvalidInputError)
				return false
			}
			med1=strings.TrimSpace(med1)
			med2,err:=reader.ReadString('\n')
			if err != nil {
				fmt.Println(errorspacket.InvalidInputError)
				return false
			}
			med2=strings.TrimSpace(med2)
			med3,err:=reader.ReadString('\n')
			if err != nil {
				fmt.Println(errorspacket.InvalidInputError)
				return false
			}
			med3=strings.TrimSpace(med3)
			fmt.Println("General notes: ")
			noteInput2,err:=reader.ReadString('\n')
			if err != nil {
				fmt.Println(errorspacket.InvalidInputError)
				return false
			}
			noteInput2=strings.TrimSpace(noteInput2)
			database.InsertJunctionMedPrescription(LastPrescriptionID,med1,noteInput2)
			database.InsertJunctionMedPrescription(LastPrescriptionID,med2,noteInput2)
			database.InsertJunctionMedPrescription(LastPrescriptionID,med3,noteInput2)
			fmt.Println("Returning to menu...")
		case "2":
			fmt.Println("Type the prescription ID you wish to delete: ")
			inputDelID,err:=reader.ReadString('\n')
			if err != nil {
				fmt.Println(errorspacket.InvalidInputError)
				return false
			}
			inputDelID=strings.TrimSpace(inputDelID)
			err = database.DeletePrescription(inputDelID)
			if err!=nil{
				fmt.Println("Error in deleting prescription")
				return false
			} else{
				fmt.Println("Deletion of prescription successful")
			}
			fmt.Println("Returning to menu...")
		case "3":
			fmt.Println("Average number of medications prescribed: ")
			AveragePres,err:=database.AverageNumberOfMedsPrescribed()
			if err!=nil{
				fmt.Println("Error selecting the average of prescriptions")
				return false
			}
			fmt.Println(AveragePres)
			fmt.Println("Returning to menu...")
		case "BACK":
			fmt.Println("Returning to menu...")
		default:
			fmt.Println("Provide a viable option.")
		}
	case "5":
		AppointmentsManager(doc)
	case "6":
		ExaminationServiceManager(doc)
	case "7":
		AdmissionLobby(doc)
	case "EXIT":
		fmt.Println("Exiting Doctor Interface...")
		return false
	default:
		fmt.Println("Provide a viable option.")
	}
	return true
}

func Main() {
	fmt.Println("--Doctor Interface--")
	logger.InitLogger()
	var ipAddress string
	fmt.Println("Introduce the IP Address and Port of the server you're connecting to: ")
	ipAddress, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(errorspacket.InvalidInputError)
		return
	}
	ipAddress = strings.TrimSpace(ipAddress)
	defer logger.CloseLogger()
	con := network.HospitalServerConnect(ipAddress)
	logger.Info("Doctor Client tries to connect to central server with IP: " + con.RemoteAddr().String())
	defer con.Close()
	fmt.Println("Connecting to the central QHospital server with IP Address ", con.RemoteAddr().String())
	for {
		fmt.Println("Choose the service you wish to proceed with: [1] Login or [2] Register")
		var input string
		fmt.Scanln(&input)
		switch input {
		case "1":
			fmt.Println("Login Service - Introduce your credentials to proceed: ")
			fmt.Println("Last Name: ")
			input1, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(errorspacket.InvalidInputError)
				return
			}
			input1 = strings.TrimSpace(input1)
			fmt.Println("First Name: ")
			input2, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(errorspacket.InvalidInputError)
				return
			}
			input2 = strings.TrimSpace(input2)
			fmt.Println("Password: ")
			input3, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(errorspacket.InvalidInputError)
				return
			}
			input3 = strings.TrimSpace(input3)
			var doctor model.Doctor
			doctor, err = database.LoginDoctor(input1, input2, input3)
			if err != nil {
				fmt.Println(errorspacket.DoctorCredentialsError, err)
				return
			}
			go func() {
				for {
					msg, err := network.ReadJSON(con)
					if err != nil {
						fmt.Println("Error receiving message or connection closed.")
						return
					}
					fmt.Printf("\n[New message received!]\nFrom: %s\nTitle: %s\n%s\n> ", msg.Sender, msg.Title, msg.Content)
				}
			}()
			for {
				continueMenu := DoctorOperationsMenu(&doctor, con)
				if !continueMenu {
					return
				}
			}
		case "2":
			fmt.Println("Register Service - Introduce your credentials")
			fmt.Println("Last Name:")
			lastName,err:=reader.ReadString('\n')
			if err!=nil{
				fmt.Println(errorspacket.InvalidInputError)
				return
			}
			lastName=strings.TrimSpace(lastName)
			fmt.Println("First Name:")
			firstName,err:=reader.ReadString('\n')
			if err!=nil{
				fmt.Println(errorspacket.InvalidInputError)
				return
			}
			firstName=strings.TrimSpace(firstName)
			fmt.Println("Specialty:")
			specialtyD,err:=reader.ReadString('\n')
			if err!=nil{
				fmt.Println(errorspacket.InvalidInputError)
				return
			}
			specialtyD=strings.TrimSpace(specialtyD)
			fmt.Println("Email:")
			emailD,err:=reader.ReadString('\n')
			if err!=nil{
				fmt.Println(errorspacket.InvalidInputError)
				return
			}
			emailD=strings.TrimSpace(emailD)
			fmt.Println("Phone:")
			phoneD,err:=reader.ReadString('\n')
			if err!=nil{
				fmt.Println(errorspacket.InvalidInputError)
				return
			}
			phoneD=strings.TrimSpace(phoneD)
			fmt.Println("Departments List you can choose from: ")
			DepListReg,err:=database.GetDepartmentsID()
			if err!=nil{
				fmt.Println("Error retrieving departments list")
				return
			}
			fmt.Println(DepListReg)
			fmt.Println("Department ID: ")
			depIDD,err:=reader.ReadString('\n')
			if err!=nil{
				fmt.Println(errorspacket.InvalidInputError)
				return
			}
			depIDD=strings.TrimSpace(depIDD)
			depIDInt,err:=strconv.Atoi(depIDD)
			if err!=nil{
				fmt.Println("Invalid format on department ID")
				return
			}
			fmt.Println("Password:")
			passwordD,err:=reader.ReadString('\n')
			if err!=nil{
				fmt.Println(errorspacket.InvalidInputError)
				return
			}
			passwordD=strings.TrimSpace(passwordD)
			doctorRegister:=model.Doctor{
				ID: 100,
				LastName: lastName,
				FirstName: firstName,
				Specialty: specialtyD,
				Email: emailD,
				Phone: phoneD,
				DepartmentID: depIDInt,
				IsHeaddoctor: false,
				Credit: 0.00,
			}
			err=database.RegisterDoctor(doctorRegister)
			if err != nil {
				fmt.Println("Registration failed:", err)
			} else {
				fmt.Printf("\nDoctor successfully registered with password %s\n!",passwordD)
				fmt.Printf("Doctor Last and First Name: %s %s\n",doctorRegister.LastName, doctorRegister.FirstName)
				fmt.Printf("Email: %s\n",doctorRegister.Email)
				fmt.Printf("Registration complete!")
			}
		default:
			fmt.Println("Provide a viable option.")
		}
	}
}
