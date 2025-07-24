package doctor
import(
	"HospitalQOps/database"
	"HospitalQOps/errorspacket"
	"HospitalQOps/model"
	"fmt"
	"strings"
)
func AdmissionLobby(doc *model.Doctor){
	fmt.Println("Admission Lobby")
	fmt.Println("Options:")
	fmt.Println("1. Admissions List")
	fmt.Println("2. Admit patient")
	fmt.Println("3. Discharge patient")
	fmt.Println("4. Update operations number")
	fmt.Println("5. Update notes and price")
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
		fmt.Println("List of admissions:")
		ListAdm,err:=database.GetAdmissions()
		if err!=nil{
			fmt.Println("Error getting admissions")
			return
		}
		fmt.Println(ListAdm)
	case "2":
		fmt.Println("Type the patients ID you wish to admit into the Hospital: ")
		patID,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		patID=strings.TrimSpace(patID)
		fmt.Println("Doctors ID (You): %d",doc.ID)
		fmt.Println("Rooms available: ")
		roomList,err:=database.GetRooms()
		if err!=nil{
			fmt.Println("Error fetching rooms")
			return
		}
		fmt.Println(roomList)
		fmt.Println("Choose a room ID: ")
		roomID,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		roomID=strings.TrimSpace(roomID)
		statusAdms:="active"
		fmt.Println("Insert some notes: ")
		notesAdm,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		notesAdm=strings.TrimSpace(notesAdm)
		fmt.Println("Insert the approximate price of the admission: ")
		priceAdm,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		priceAdm=strings.TrimSpace(priceAdm)
		err=database.InsertAdmissionInstance(patID,doc.ID,roomID,statusAdms,notesAdm,priceAdm)
		if err != nil {
			fmt.Println("Error admitting patient")
		} else {
			fmt.Println("Admission processed and inserted.")
		}
	case "3":
		fmt.Println("Admission ID you wish to discharge: ")
		admIDDis,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		admIDDis=strings.TrimSpace(admIDDis)
		err=database.UpdateAdmissionStatus(admIDDis)
		if err!=nil{
			fmt.Println("Error discharging")
		} else{
			fmt.Println("Discharge processed succesfully")
		}
	case "4":
		fmt.Println("Admission ID you wish to update: ")
		admIDOp,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		admIDOp=strings.TrimSpace(admIDOp)
		fmt.Println("Number of operations to add: ")
		opNo,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		opNo=strings.TrimSpace(opNo)
		err=database.UpdateAdmissionOperationsNo(admIDOp,opNo)
		if err!=nil{
			fmt.Println("Error updating operations number")
		} else{
			fmt.Println("Updating operations number succesful")
		}
	case "5":
		fmt.Println("Admission ID you wish to add notes and a new price: ")
		admIDN,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		admIDN=strings.TrimSpace(admIDN)
		fmt.Println("Notes: ")
		notesAD,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		notesAD=strings.TrimSpace(notesAD)
		fmt.Println("Price: ")
		priceAD,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		priceAD=strings.TrimSpace(priceAD)
		err=database.UpdateAdmissionNotesAndPrice(admIDN,notesAD,priceAD)
		if err!=nil{
			fmt.Println("Error updating notes")
		} else{
			fmt.Println("Updating notes and price succesful")
		}
	case "BACK":
		fmt.Println("Exiting Admission Lobby... Back to menu")
	default:
		fmt.Println("Choose a viable option")
	}
}