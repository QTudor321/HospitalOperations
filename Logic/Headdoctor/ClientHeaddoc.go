package headdoctor
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
func HospitalEmailCenterHeaddoctor(hdoc *model.Doctor, connection net.Conn) {
	fmt.Println("Welcome to the Hospital Email Center")
	fmt.Printf("List of Patients for %s %s\n", hdoc.LastName, hdoc.FirstName)
	patientsListing, err := database.GetPatientsListByDoctor(*hdoc)
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
			Sender:   hdoc.Email,
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
		InboxList, err := database.GetDoctorInbox(*hdoc)
		if err != nil {
			fmt.Println("Error retrieving headdoctors inbox")
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
		conversation, err := database.GetConversationEmail(*hdoc, EmailInput)
		for i, msg := range conversation {
			fmt.Printf("\nIndex [%d]: Sender %s Receiver %s Title %s\n", i+1, msg.Sender, msg.Receiver, msg.Title)
			fmt.Printf("Message: %s", msg.Content)
		}
	case "BACK":
	default:
		fmt.Println("Choose a viable option")
	}
}
func HeadDoctorOperationsMenu(hdoc *model.Doctor, connection net.Conn) bool{
	fmt.Printf("Welcome back HeadDoctor %s %s!\n", hdoc.LastName, hdoc.FirstName)
	fmt.Println("Options:")
	fmt.Println("1.Patients")
	fmt.Println("2.Doctors")
	fmt.Println("3.Nurses")
	fmt.Println("4.Email")
	fmt.Println("5.Operations")
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
		patientsListing, err := database.GetPatientsListByDoctor(*hdoc)
		if err != nil {
			fmt.Println("Error retrieving patients data.", err)
		}
		fmt.Println(patientsListing)
		fmt.Println("List of patients registered yet:")
		patientsAllList,err:=database.GetAllPatients()
		if err != nil {
			fmt.Println("Error retrieving all patients data.", err)
		}
		fmt.Println(patientsAllList)
	case "2":
		fmt.Println("Doctor Options:")
		fmt.Println("1.View all doctors")
		fmt.Println("2.View appointments")
		fmt.Println("3.View examinations")
		fmt.Println("4.View admissions")
		fmt.Println("5.View rooms")
		inputHeadC,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return false
		}
		inputHeadC=strings.TrimSpace(inputHeadC)
		switch inputHeadC{
		case "1":
			doctorsList,err:=database.GetAllDoctors()
			if err!=nil{
				fmt.Println("Error retrieving doctors data.",err)
			}
			fmt.Println(doctorsList)
			fmt.Println("Returning to menu...")
		case "2":
			appointmentsList,err:=database.GetAppointments()
			if err!=nil{
				fmt.Println("Error retrieving appointments data.",err)
			}
			fmt.Println(appointmentsList)
			fmt.Println("Returning to menu...")
		case "3":
			examinationsList,err:=database.GetExaminations()
			if err!=nil{
				fmt.Println("Error retrieving examinations data.",err)
			}
			fmt.Println(examinationsList)
			fmt.Println("Returning to menu...")
		case "4":
			admissionsList,err:=database.GetAdmissions()
			if err!=nil{
				fmt.Println("Error retrieving admissions data.",err)
			}
			fmt.Println(admissionsList)
			fmt.Println("Returning to menu...")
		case "5":
			roomsList,err:=database.GetRooms()
			if err!=nil{
				fmt.Println("Error retrieving rooms data",err)
			}
			fmt.Println(roomsList)
			fmt.Println("Returning to menu...")
		default:
			fmt.Println("Choose a viable option.")
		}
	case "3":
		fmt.Println("Options:")
		fmt.Println("1.View Nurses")
		fmt.Println("2.View Notes")
		inputCase3,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return false
		}
		inputCase3=strings.TrimSpace(inputCase3)
		switch inputCase3{
		case "1":
			fmt.Println("Nurse List:")
			nursesListing, err := database.GetNursesAndDoctorsByDepartments()
			if err != nil {
				fmt.Println("Error retrieving nurses and doctors data.", err)
			}
			fmt.Println(nursesListing)
			fmt.Println("Returning to menu...")
		case "2":
			junctionDoctorsNurses,err:=database.GetNursesDoctors()
			if err!=nil{
				fmt.Println("Error retrieving doctors and nurses")
			}
			fmt.Println(junctionDoctorsNurses)
		default:
			fmt.Println("Choose a viable option")
		}
	case "4":
		HospitalEmailCenterHeaddoctor(hdoc, connection)
	case "5":
		OperationsServiceManager(hdoc)
	case "EXIT":
		fmt.Println("Exiting HeadDoctor Interface...")
		return false
	default:
		fmt.Println("Provide a viable option.")
	}
	return true
}
func Main(){
	fmt.Println("--Head Doctor Interface--")
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
	logger.Info("Headdoctor Client tries to connect to central server with IP: " + con.RemoteAddr().String())
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
				continueMenu := HeadDoctorOperationsMenu(&doctor, con)
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
				IsHeaddoctor: true,
				Credit: 10000.00,
			}
			err=database.RegisterHeaddoctor(doctorRegister)
			if err != nil {
				fmt.Println("Registration failed:", err)
			} else {
				fmt.Printf("\nHead Doctor successfully registered with password %s\n!",passwordD)
				fmt.Printf("Doctor Last and First Name: %s %s\n",doctorRegister.LastName, doctorRegister.FirstName)
				fmt.Printf("Email: %s\n",doctorRegister.Email)
				fmt.Printf("Headdoctor: %s\n",doctorRegister.IsHeaddoctor)
				fmt.Printf("Registration complete!")
			}
		default:
			fmt.Println("Provide a viable option.")
		}
	}
}