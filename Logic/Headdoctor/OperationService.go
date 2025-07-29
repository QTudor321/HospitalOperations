package headdoctor
import (
	"HospitalQOps/model"
	"HospitalQOps/errorspacket"
	"HospitalQOps/database"
	"fmt"
	"strings"
	"strconv"
	"time"
	"database/sql"
)
func OperationsServiceManager(headdoctor *model.Doctor){
	fmt.Println("Operations Service Manager")
	fmt.Println("Options:")
	fmt.Println("1.View operations")
	fmt.Println("2.Operate patient")
	fmt.Println("3.Number of operations per month")
	fmt.Println("BACK")
	inputO,err:=reader.ReadString('\n')
	if err!=nil{
		fmt.Println(errorspacket.InvalidInputError)
		return
	}
	inputO=strings.TrimSpace(inputO)
	inputOU:=strings.ToUpper(inputO)
	switch inputOU{
	case "1":
		operationsList,err:=database.GetOperations()
		if err!=nil{
			fmt.Println("Error retrieving operations data")
		}
		fmt.Println(operationsList)
	case "2":
		fmt.Println("Operation Utility")
		fmt.Println("Introduce admission ID:")
		adID,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		adID=strings.TrimSpace(adID)
		adIDF,err:=strconv.Atoi(adID)
		if err!=nil{
			fmt.Println("Invalid format")
			return
		}
		fmt.Println("Type the room ID:")
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
		fmt.Println("Type the patient ID:")
		patID,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		patID=strings.TrimSpace(patID)
		patIDF,err:=strconv.Atoi(patID)
		if err!=nil{
			fmt.Println("Invalid format")
			return
		}
		fmt.Println("Headdoctors ID - Surgeon ID: ",headdoctor.ID)
		fmt.Println("Schedule date of operation of format (YYYY-MM-DD HH:MM:SS):")
		schDate,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		schDate=strings.TrimSpace(schDate)
		schDateF,err:=time.Parse("2006-01-02 15:04:05", schDate)
		if err!=nil{
			fmt.Println("Invalid format")
			return
		}
		fmt.Println("Information:")
		info,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		info=strings.TrimSpace(info)
		fmt.Println("Assistant nurses:")
		nursesA,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		nursesA=strings.TrimSpace(nursesA)
		fmt.Println("Assistant doctors:")
		doctorsA,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		doctorsA=strings.TrimSpace(doctorsA)
		fmt.Println("Next operation date (fill or leave empty):")
		nextOP,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		nextOP=strings.TrimSpace(nextOP)
		var nextOperationDate sql.NullTime
		if nextOP != "" {
			parsed, err := time.Parse("2006-01-02 15:04:05", nextOP)
			if err != nil {
				fmt.Println("Invalid next-operation date format.")
				return
			}
			nextOperationDate = sql.NullTime{Time: parsed, Valid: true}
		} else {
			nextOperationDate = sql.NullTime{Valid: false}
		}
		fmt.Println("Finished date:")
		finishOp,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		finishOp=strings.TrimSpace(finishOp)
		finishOpF,err:=time.Parse("2006-01-02 15:04:05", finishOp)
		if err!=nil{
			fmt.Println("Invalid format")
			return
		}
		fmt.Println("Price:")
		priceO,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		priceO=strings.TrimSpace(priceO)
		priceOF, err := strconv.ParseFloat(priceO, 64)
		if err != nil {
			fmt.Println("Invalid price.")
			return
		}
		fmt.Println("Operation Notes:")
		notesO,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		notesO=strings.TrimSpace(notesO)
		err=database.OperationProcedure(adIDF,roomIDF,patIDF,headdoctor.ID,schDateF,info,nursesA,doctorsA,nextOperationDate,finishOpF,priceOF,notesO)
		if err != nil {
			fmt.Println("Operation failed:", err)
		} else {
			fmt.Println("Operation logged successfully!")
		}
	case "3":
		fmt.Println("Number of operations per month: ")
		statsOpsM,err:=database.OperationStatistics()
		if err!=nil{
			fmt.Println("Error retrieving operation statistics")
		}
		fmt.Println(statsOpsM)
	case "BACK":
		fmt.Println("Exiting Operations Utility...")
	default:
		fmt.Println("Choose a viable option!")
	}
}
