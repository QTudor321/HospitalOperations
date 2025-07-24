package patient
import(
	"HospitalQOps/database"
	"HospitalQOps/errorspacket"
	"HospitalQOps/model"
	"fmt"
	"strings"
	"time"
	"strconv"
)
func AppointmentHub(pat *model.Patient){
	fmt.Println("Appointment Hub")
	fmt.Println("Options:")
	fmt.Println("1. Appointment List")
	fmt.Println("2. Schedule")
	fmt.Println("BACK")
	inputAdm,err:=reader.ReadString('\n')
	if err!=nil{
		fmt.Println(errorspacket.InvalidInputError)
		return
	}
	inputAdm=strings.TrimSpace(inputAdm)
	inputAdmU:=strings.ToUpper(inputAdm)
	switch inputAdmU{
	case "1":
		fmt.Println("List of appointments:")
		ListAppPat,err:=database.GetAppointmentsByPatient(pat.ID)
		if err!=nil{
			fmt.Println("Error getting appointments")
			return
		}
		fmt.Println(ListAppPat)
	case "2":
		fmt.Println("Patients ID (You): %d",pat.ID)
		fmt.Println("Type the doctors ID you wish to schedule an appointment with: ")
		docID,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		docID=strings.TrimSpace(docID)
		docIDF,err:=strconv.Atoi(docID)
		if err!=nil{
			fmt.Println("Invalid format")
			return
		}
		fmt.Println("Choose a room ID: ")
		roomID,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		roomID=strings.TrimSpace(roomID)
		roomIDF,err:=strconv.Atoi(roomID)
		if err!=nil{
			fmt.Println("Invalid format")
			return
		}
		fmt.Println("Insert the date with format [YYYY-MM-DD HH:MM:SS]: ")
		dateAppP,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		dateAppP=strings.TrimSpace(dateAppP)
		dateAppPF,err:=time.Parse("2006-01-02 15:04:05", dateAppP)
		if err!=nil{
			fmt.Println("Invalid appointment date format")
			return
		}
		err=database.InsertAppointment(pat.ID,docIDF,roomIDF,dateAppPF)
		if err != nil {
			fmt.Println("Error creating appointment")
		} else {
			fmt.Println("Appointment processed and inserted.")
		}
	case "BACK":
		fmt.Println("Exiting Appointment Hub... Back to menu")
	default:
		fmt.Println("Choose a viable option")
	}
}