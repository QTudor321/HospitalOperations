package patient
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"net"
	"HospitalQOps/logger"
	"HospitalQOps/errorspacket"
	"HospitalQOps/model"
	"HospitalQOps/database"
	"HospitalQOps/network"
)
var reader = bufio.NewReader(os.Stdin)
func HospitalEmailCenterPatient(pat *model.Patient, connection net.Conn) {
	fmt.Println("Welcome to the Hospital Email Center")
	fmt.Printf("List of Doctors for %s %s\n", pat.LastName, pat.FirstName)
	doctorsListing, err := database.GetDoctorsListByPatient(*pat)
	if err != nil {
		fmt.Println("Error retrieving doctors list")
		return
	}
	fmt.Println(doctorsListing)
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
			Sender:   pat.Email,
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
		InboxList, err := database.GetPatientInbox(*pat)
		if err != nil {
			fmt.Println("Error retrieving patients inbox")
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
		conversation, err := database.GetConversationEmailPatient(*pat, EmailInput)
		for i, msg := range conversation {
			fmt.Printf("\nIndex [%d]: Sender %s Receiver %s Title %s\n", i+1, msg.Sender, msg.Receiver, msg.Title)
			fmt.Printf("Message: %s", msg.Content)
		}
	case "BACK":
	default:
		fmt.Println("Choose a viable option")
	}
}
func PatientOperationsMenu(pat *model.Patient, connection net.Conn) bool{
	fmt.Printf("Welcome back Patient %s %s!\n", pat.LastName, pat.FirstName)
	fmt.Println("Options:")
	fmt.Println("1.Departments")
	fmt.Println("2.Email")
	fmt.Println("3.Doctors")
	fmt.Println("4.Rooms")
	fmt.Println("5.Appointments")
	fmt.Println("6.Examinations")
	fmt.Println("7.Admissions")
	fmt.Println("8.Operations")
	fmt.Println("9.Payment")
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
		departmentsListing, err := database.GetDepartmentsID()
		if err != nil {
			fmt.Println("Error retrieving department data.", err)
		}
		fmt.Println(departmentsListing)
		fmt.Println("Returning to menu...")
	case "2":
		HospitalEmailCenterPatient(pat, connection)
	case "3":
		doctorsListing3, err := database.GetDoctorsListByPatient(*pat)
		if err != nil {
			fmt.Println("Error retrieving doctors data.", err)
		}
		fmt.Println(doctorsListing3)
		fmt.Println("Returning to menu...")
	case "4":
		roomList, err:=database.GetRooms()
		if err!=nil{
			fmt.Println("Error retrieveing rooms",err)
		}
		fmt.Println(roomList)
	case "5":
		AppointmentHub(pat)
	case "6":
		examinationsPatientList,err:=database.GetExaminationsPrescriptionsPatient(*pat)
		if err!=nil{
			fmt.Println("Failed to retrieve patients examinations")
		}
		fmt.Println(examinationsPatientList)
	case "7":
		admissionListPat,err:=database.GetAdmissionsPatient(*pat)
		if err!=nil{
			fmt.Println("Failed to retrieve patients admissions")
		}
		fmt.Println(admissionListPat)
	case "8":
		fmt.Println("Operations results:")
		operationsResultPatient,err:=database.GetOperationsPatient(*pat)
		if err!=nil{
			fmt.Println("Failed to retrieve patients admissions")
		}
		fmt.Println(operationsResultPatient)
	case "9":
		PaymentService(pat)
	case "EXIT":
		fmt.Println("Exiting Patient Interface...")
		return false
	default:
		fmt.Println("Provide a viable option.")
	}
	return true
}
func Main(){
	fmt.Println("--Patient Interface--")
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
	logger.Info("Patient Client tries to connect to central server with IP: " + con.RemoteAddr().String())
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
			var patient model.Patient
			patient, err = database.LoginPatient(input1, input2, input3)
			if err != nil {
				fmt.Println(errorspacket.PatientCredentialsError, err)
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
				continueMenu := PatientOperationsMenu(&patient, con)
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
			fmt.Println("Age:")
			ageD,err:=reader.ReadString('\n')
			if err!=nil{
				fmt.Println(errorspacket.InvalidInputError)
				return
			}
			ageD=strings.TrimSpace(ageD)
			ageDInt,err:=strconv.Atoi(ageD)
			if err!=nil{
				fmt.Println("Invalid format on age")
				return
			}
			fmt.Println("Gender:")
			genderP,err:=reader.ReadString('\n')
			if err!=nil{
				fmt.Println(errorspacket.InvalidInputError)
				return
			}
			genderP=strings.TrimSpace(genderP)
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
			fmt.Println("Address: ")
			addressPat,err:=reader.ReadString('\n')
			if err!=nil{
				fmt.Println(errorspacket.InvalidInputError)
				return
			}
			addressPat=strings.TrimSpace(addressPat)
			fmt.Println("Password:")
			passwordP,err:=reader.ReadString('\n')
			if err!=nil{
				fmt.Println(errorspacket.InvalidInputError)
				return
			}
			passwordP=strings.TrimSpace(passwordP)
			patientRegister:=model.Patient{
				ID: 100,
				LastName: lastName,
				FirstName: firstName,
				Age: ageDInt,
				Gender: genderP,
				Email: emailD,
				Phone: phoneD,
				Address: addressPat,
				Credit: 700.00,
			}
			err=database.RegisterPatient(patientRegister)
			if err != nil {
				fmt.Println("Registration failed:", err)
			} else {
				fmt.Printf("\nPatient successfully registered with password %s\n!",passwordP)
				fmt.Printf("Patient Last and First Name: %s %s\n",patientRegister.LastName, patientRegister.FirstName)
				fmt.Printf("Email: %s\n",patientRegister.Email)
				fmt.Printf("Registration complete!")
			}
		default:
			fmt.Println("Provide a viable option.")
		}
	}
}