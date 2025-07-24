package doctor
import(
	"HospitalQOps/database"
	"HospitalQOps/errorspacket"
	"HospitalQOps/model"
	"fmt"
	"strings"
	"strconv"
	"time"
	"database/sql"
)
func ExaminationServiceManager(doc *model.Doctor){
	fmt.Println("Examination Service Menu")
	fmt.Println("Options:")
	fmt.Println("1. View examinations list")
	fmt.Println("2. Examine patient")
	fmt.Println("3. Order")
	fmt.Println("BACK")
	inputE,err:=reader.ReadString('\n')
	if err!=nil{
		fmt.Println(errorspacket.InvalidInputError)
		return
	}
	inputE=strings.TrimSpace(inputE)
	inputEU:=strings.ToUpper(inputE)
	switch inputEU{
	case "1":
		fmt.Println("Examinations list: ")
		ExamList,err:=database.GetExaminations()
		if err!=nil{
			fmt.Println("Failed to fetch examinations")
			return
		}
		fmt.Println(ExamList)
	case "2":
		fmt.Println("Examination Procedure initiated")
		fmt.Println("Type the appointment ID: ")
		appID,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		appID=strings.TrimSpace(appID)
		appIDF,err:=strconv.Atoi(appID)
		if err!=nil{
			fmt.Println("Invalid format")
			return
		}
		fmt.Println("Examination Date of format (YYYY-MM-DD HH:MM:SS): ")
		examDate,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		examDate=strings.TrimSpace(examDate)
		examDateF,err:=time.Parse("2006-01-02 15:04:05", examDate)
		if err!=nil{
			fmt.Println("Invalid format")
			return
		}
		fmt.Println("Informational notes: ")
		infoN,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		infoN=strings.TrimSpace(infoN)
		fmt.Println("Nurses list which assist you: ")
		nurseList,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		nurseList=strings.TrimSpace(nurseList)
		fmt.Println("Diagnosis: ")
		diag,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		diag=strings.TrimSpace(diag)
		fmt.Println("Result [admission/reexamination/prescription]: ")
		resultE,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		resultE=strings.TrimSpace(resultE)
		fmt.Println("Reexamination Date (fill or leave empty if its null): ")
		reexamDate,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		reexamDate=strings.TrimSpace(reexamDate)
		var reexamDateF sql.NullTime
		if reexamDate != "" {
			parsed, err := time.Parse("2006-01-02 15:04:05", reexamDate)
			if err != nil {
				fmt.Println("Invalid re-examination date format.")
				return
			}
			reexamDateF = sql.NullTime{Time: parsed, Valid: true}
		} else {
			reexamDateF = sql.NullTime{Valid: false}
		}
		fmt.Println("Price: ")
		priceE,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		priceE=strings.TrimSpace(priceE)
		priceF, err := strconv.ParseFloat(priceE, 64)
		if err != nil {
			fmt.Println("Invalid price.")
			return
		}
		fmt.Println("Level of severity[none/mild/moderate/severe/critical/terminal]: ")
		levelE,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		levelE=strings.TrimSpace(levelE)
		fmt.Println("Examination notes by doctor %s %s: ",doc.FirstName,doc.LastName)
		notesE,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		notesE=strings.TrimSpace(notesE)
		fmt.Println("Nurse ID: ")
		nurseIDJ,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		nurseIDJ=strings.TrimSpace(nurseIDJ)
		nurseIDJF,err:=strconv.Atoi(nurseIDJ)
		if err!=nil{
			fmt.Println("Invalid format")
			return
		}
		fmt.Println("Further examination notes involving nurses help: ")
		notesNurse,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		notesNurse=strings.TrimSpace(notesNurse)
		err=database.InsertJunctionNurseDoctors(nurseIDJF,doc.ID,notesNurse)
		if err!=nil{
			fmt.Println("Error logging into junction table nurses help: ",err)
		} else {
			fmt.Println("Nurses involvement logged!")
		}
		err=database.ExaminationProcedure(appIDF,examDateF,infoN,nurseList,diag,resultE,reexamDateF,priceF,levelE,notesE)
		if err != nil {
			fmt.Println("Examination failed:", err)
		} else {
			fmt.Println("Examination recorded successfully!")
		}
	case "3":
		fmt.Println("Choose the ordering option [Result/Level]: ")
		orderChoice,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		orderChoice=strings.TrimSpace(orderChoice)
		orderChoiceU:=strings.ToUpper(orderChoice)
		switch orderChoiceU{
		case "RESULT":
			fmt.Println("Ordering examination list by result: ")
			ExamResultList,err:=database.OrderExamResult()
			if err!=nil{
				fmt.Println("Error ordering examination by result")
				return
			}
			fmt.Println(ExamResultList)
		case "LEVEL":
			fmt.Println("Ordering examination list by level: ")
			ExamLevelList,err:=database.OrderExamLevel()
			if err!=nil{
				fmt.Println("Error ordering examination by level")
				return
			}
			fmt.Println(ExamLevelList)
		default:
			fmt.Println("Choose the available ordering option!")
		}
	case "BACK":
		fmt.Println("Exiting Examination Utility...")
	default:
		fmt.Println("Choose a viable option!")
	}
}
