package patient
import(
	"HospitalQOps/database"
	"HospitalQOps/errorspacket"
	"HospitalQOps/model"
	"database/sql"
	"fmt"
	"strings"
	"strconv"
)
func PaymentService(pat *model.Patient){
	fmt.Println("Payment Service Menu")
	fmt.Println("Options:")
	fmt.Println("1. View payments")
	fmt.Println("2. Pay")
	fmt.Println("3. View total payments")
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
		paymentsListPatient,err:=database.GetPayments(pat.ID)
		if err!=nil{
			fmt.Println("Error retrieveing patients payments")
			return
		}
		fmt.Println(paymentsListPatient)
	case "2":
		fmt.Println("Payment System")
		fmt.Println("Introduce operation ID or leave empty: ")
		opsID,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		opsID=strings.TrimSpace(opsID)
		var opsIDF sql.NullInt64
		if opsID != "" {
			idVal, err := strconv.Atoi(opsID)
			if err != nil {
				fmt.Println("Invalid operation ID format")
				return
			}
			opsIDF = sql.NullInt64{Int64: int64(idVal), Valid: true}
		}
		fmt.Println("Patients ID (you)",pat.ID)
		fmt.Println("Introduce admission ID or leave empty: ")
		admID,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		admID=strings.TrimSpace(admID)
		var admIDF sql.NullInt64
		if admID != "" {
			idVal, err := strconv.Atoi(admID)
			if err != nil {
				fmt.Println("Invalid admission ID format")
				return
			}
			admIDF = sql.NullInt64{Int64: int64(idVal), Valid: true}
		}
		fmt.Println("Introduce examination ID or leave empty: ")
		examID,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		examID=strings.TrimSpace(examID)
		var examIDF sql.NullInt64
		if examID != "" {
			idVal, err := strconv.Atoi(examID)
			if err != nil {
				fmt.Println("Invalid examination ID format")
				return
			}
			examIDF = sql.NullInt64{Int64: int64(idVal), Valid: true}
		}
		fmt.Println("Introduce the amount you pay:")
		payAmount,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(errorspacket.InvalidInputError)
			return
		}
		payAmount=strings.TrimSpace(payAmount)
		payAmountF, err := strconv.ParseFloat(payAmount, 64)
		if err != nil {
			fmt.Println("Invalid price.")
			return
		}
		err=database.PaymentProcedure(opsIDF,pat.ID,admIDF,examIDF,payAmountF)
		if err != nil {
			fmt.Println("Payment failed:", err)
		} else {
			fmt.Println("Payment recorded successfully!")
		}
	case "3":
		fmt.Println("Total payments for patient: ",pat.LastName,pat.FirstName)
		totalPaymentsPatient,err:=database.SumPaymentsPatient(pat.ID)
		if err!=nil{
			fmt.Println("Failed to retrieve sum of total payments")
			return
		}
		fmt.Println(totalPaymentsPatient)
	case "BACK":
		fmt.Println("Exiting Payment Service... Back to menu")
	default:
		fmt.Println("Choose a viable option")
	}
}