package database
import (
	"HospitalQOps/errorspacket"
	"fmt"
	"database/sql"
	"strings"
	"time"
)
func nullIntToString(n sql.NullInt64) string {
	if n.Valid {
		return fmt.Sprintf("%d", n.Int64)
	}
	return "N/A"
}
func GetPayments(patID int) (string,error){
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "",err
	}
	defer db.Close()
	query:=`SELECT payment_id,operation_id,patient_id,admission_id,examination_id,amount,date_pay
	FROM Payment
	WHERE patient_id=$1`
	rows,err:=db.Query(query,patID)
	if err!=nil{
		fmt.Println(errorspacket.QueryError)
		return "",err
	}
	defer rows.Close()
	var ListPayments strings.Builder
	for rows.Next() {
		var payID int
		var opsID, admID, examID sql.NullInt64
		var patIDP int
		var amountPay float64
		var datePay time.Time
		if err := rows.Scan(&payID, &opsID, &patIDP, &admID, &examID, &amountPay,&datePay); err != nil {
			return "", errorspacket.QueryError
		}
		ListPayments.WriteString(fmt.Sprintf(
			"Payment ID: %d | Operation ID: %v | Admission ID: %v | Examination ID: %v\n",
			payID,
			nullIntToString(opsID),
			nullIntToString(admID),
			nullIntToString(examID),
		))
		ListPayments.WriteString(fmt.Sprintf("Patient ID: %d\n", patIDP))
		ListPayments.WriteString(fmt.Sprintf("Payment Date: [%s]\n", datePay.Format("2006-01-02 15:04:05")))
		ListPayments.WriteString(fmt.Sprintf("Payment amount: %.2f\n\n",amountPay))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return ListPayments.String(), nil
}
func PaymentProcedure(opsID sql.NullInt64, patID int, admID sql.NullInt64, examID sql.NullInt64, payAmount float64) error{
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()
	instruction:=`INSERT INTO Payment (operation_id, patient_id, admission_id, examination_id, amount, date_pay)
	VALUES ($1, $2, $3, $4, $5, NOW())`
	_,err=db.Exec(instruction, opsID, patID, admID, examID, payAmount)
	if err!=nil{
		fmt.Println(errorspacket.ExecutionError)
		return errorspacket.ExecutionError
	}
	return nil
}
func SumPaymentsPatient(patID int) (float64, error) {
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return 0, err
	}
	defer db.Close()

	query := `
		SELECT SUM(amount)
		From Payment
		Where patient_id=$1;
	`
	var total float64
	err = db.QueryRow(query, patID).Scan(&total)
	if err != nil {
		fmt.Println(errorspacket.QueryError)
		return 0, errorspacket.QueryError
	}
	return total, nil
}