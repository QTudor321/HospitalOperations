package database
import (
	"HospitalQOps/errorspacket"
	"fmt"
	"strings"
	"time"
	"database/sql"
)
func OperationProcedure(adIDF int, roomIDF int, patIDF int, surgeonID int, 
	schDateF time.Time, info string, nursesA string, doctorsA string,
	nextOperationDate sql.NullTime, finishOpF time.Time, priceOF float64, notesO string) error{
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()
	query := `INSERT INTO Operations (
		admission_id, room_id, patient_id, surgeon_id, 
		scheduled_date, information, assistant_nurses, assistant_doctors, next_operation, finished_date,
		price, operation_notes
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err = db.Exec(query, adIDF, roomIDF, patIDF, surgeonID, schDateF, info, nursesA, doctorsA, nextOperationDate, finishOpF, priceOF, notesO)
	if err != nil {
		fmt.Println(errorspacket.ExecutionError)
		return errorspacket.ExecutionError
	}
	return nil
}
func OperationStatistics() (string,error){
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "",err
	}
	defer db.Close()
	query:=`SELECT DATE_TRUNC('month', scheduled_date) AS month, COUNT(*) AS total_operations
		FROM Operations
		GROUP BY month
		ORDER BY month;
	`
	rows,err:=db.Query(query)
	if err!=nil{
		fmt.Println(errorspacket.QueryError)
		return "",err
	}
	defer rows.Close()
	var OpStats strings.Builder
	for rows.Next() {
		var monthOp time.Time
		var numberOp int
		if err := rows.Scan(&monthOp, &numberOp); err != nil {
			return "", errorspacket.QueryError
		}
		OpStats.WriteString(fmt.Sprintf("Month and number of operations: %s %d\n", monthOp.Format("2006-01-02 15:04:05"),numberOp))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return OpStats.String(), nil
}
func GetOperations() (string,error){
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "", err
	}
	defer db.Close()

	query := `
		SELECT * FROM Operations
	`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(errorspacket.QueryError)
		return "", errorspacket.QueryError
	}
	defer rows.Close()

	var listOps strings.Builder

	for rows.Next() {
		var opsID, admissionID, roomID, patID, docID int
		var schedOpDate time.Time
		var opInfo, assistantNurses, assistantDocs string
		var nextOpsDate, finishedDate sql.NullTime
		var opsPrice float64
		var notesOps string
		if err := rows.Scan(&opsID, &admissionID, &roomID, &patID, &docID, &schedOpDate,&opInfo,&assistantNurses,&assistantDocs,&nextOpsDate,&finishedDate,&opsPrice,&notesOps); err != nil {
			return "", errorspacket.QueryError
		}
		listOps.WriteString(fmt.Sprintf("Operation, Admission, Room, Patient and Surgeon ID's: %d %d %d %d %d\n", opsID, admissionID, roomID, patID, docID))
		listOps.WriteString(fmt.Sprintf("Operation date, Information, Assistant Nurses: [%s] %s %s\n", schedOpDate.Format("2006-01-02 15:04:05"),opInfo,assistantNurses))
		listOps.WriteString(fmt.Sprintf("Assistant doctors: %s\n", assistantDocs))
		if nextOpsDate.Valid {
			listOps.WriteString(fmt.Sprintf("Next Operation Date: %s\n", nextOpsDate.Time.Format("2006-01-02 15:04:05")))
		} else {
			listOps.WriteString("Next Operation Date: NULL\n")
		}
		if finishedDate.Valid {
			listOps.WriteString(fmt.Sprintf("Finished Date: %s\n", finishedDate.Time.Format("2006-01-02 15:04:05")))
		} else {
			listOps.WriteString("Finished Date: NULL\n")
		}
		listOps.WriteString(fmt.Sprintf("Operation price: %.2f\n", opsPrice))
		listOps.WriteString(fmt.Sprintf("Operation notes: %s\n", notesOps))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return listOps.String(), nil
}