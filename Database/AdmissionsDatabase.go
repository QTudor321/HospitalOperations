package database
import (
	"HospitalQOps/errorspacket"
	"HospitalQOps/model"
	"fmt"
	"strings"
	"strconv"
	"time"
	"database/sql"
	"errors"
)


func GetAdmissions() (string, error){
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "",err
	}
	defer db.Close()
	query:=`SELECT admission_id,patient_id,doctor_id,room_id,start_date,operations_number,discharge_date,status,doctor_notes,price
	FROM Admissions`
	rows,err:=db.Query(query)
	if err!=nil{
		fmt.Println(errorspacket.QueryError)
		return "",err
	}
	defer rows.Close()
	var ListAdmissions strings.Builder
	for rows.Next() {
		var admID, patID, docID, roomID, opNo int
		var priceAdm float64
		var dischargeD sql.NullTime
		var startD time.Time
		var statAdm, docNotes string
		if err := rows.Scan(&admID, &patID, &docID, &roomID, &startD, &opNo,&dischargeD,&statAdm,&docNotes,&priceAdm); err != nil {
			return "", errorspacket.QueryError
		}
		ListAdmissions.WriteString(fmt.Sprintf("Admission, Patient, Doctor, Room ID's: %d %d %d %d\n", admID, patID, docID, roomID))
		ListAdmissions.WriteString(fmt.Sprintf("Start Date: [%s]\n", startD.Format("2006-01-02 15:04:05")))
		if dischargeD.Valid {
			ListAdmissions.WriteString(fmt.Sprintf("Discharge Date: [%s]\n", dischargeD.Time.Format("2006-01-02 15:04:05")))
		} else {
			ListAdmissions.WriteString("Discharge Date: N/A\n")
		}
		ListAdmissions.WriteString(fmt.Sprintf("Operations Number: %d\n",opNo))
		ListAdmissions.WriteString(fmt.Sprintf("Status: %s\n",statAdm))
		ListAdmissions.WriteString(fmt.Sprintf("Notes and total price: %s %.2f\n\n",docNotes, priceAdm))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return ListAdmissions.String(), nil
}
func GetRooms() (string,error){
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "",err
	}
	defer db.Close()
	query:=`SELECT * FROM Rooms`
	rows,err:=db.Query(query)
	if err!=nil{
		fmt.Println(errorspacket.QueryError)
		return "",err
	}
	defer rows.Close()
	var ListRooms strings.Builder
	for rows.Next() {
		var roomID, number, depID,capacity int
		var typeR string
		if err := rows.Scan(&roomID, &number, &depID, &typeR, &capacity); err != nil {
			return "", errorspacket.QueryError
		}
		ListRooms.WriteString(fmt.Sprintf("Room ID: %d\n", roomID))
		ListRooms.WriteString(fmt.Sprintf("Number Room and Department ID: %d %d\n", number,depID))
		ListRooms.WriteString(fmt.Sprintf("Type of room and capacity: %s %d\n\n",typeR,capacity))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return ListRooms.String(), nil
}
func InsertAdmissionInstance(patID string, docID int, roomID string, statusAdms string, notesAdm string, priceAdm string) error{
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()
	patientID, err := strconv.Atoi(patID)
	if err != nil {
		fmt.Println("Invalid patient ID format.")
		return errors.New("invalid patient ID")
	}
	roomIDInt, err := strconv.Atoi(roomID)
	if err != nil {
		fmt.Println("Invalid room ID format.")
		return errors.New("invalid room ID")
	}
	priceFloat, err := strconv.ParseFloat(priceAdm, 64)
	if err != nil {
		fmt.Println("Invalid price format.")
		return errors.New("invalid price")
	}
	instruction:=`INSERT INTO Admissions (patient_id, doctor_id, room_id, start_date, operations_number, status, doctor_notes, price)
	VALUES ($1, $2, $3, NOW(), 0, $4, $5, $6)`
	_,err=db.Exec(instruction,patientID,docID, roomIDInt, statusAdms, notesAdm,priceFloat)
	if err!=nil{
		fmt.Println(errorspacket.ExecutionError)
		return errorspacket.ExecutionError
	}
	return nil
}
func UpdateAdmissionStatus(admIDDis string) error{
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()
	admissionID, err := strconv.Atoi(admIDDis)
	if err != nil {
		fmt.Println("Invalid admission ID format.")
		return errors.New("invalid admission ID")
	}
	updatequery:=`UPDATE Admissions
	SET discharge_date=NOW(), status='discharged'
	WHERE admission_id=$1`
	_,err=db.Exec(updatequery,admissionID)
	if err!=nil{
		fmt.Println(errorspacket.ExecutionError)
		return errorspacket.ExecutionError
	}
	return nil
}
func UpdateAdmissionOperationsNo(admIDOp,opNo string) error{
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()
	admissionIDOp, err := strconv.Atoi(admIDOp)
	if err != nil {
		fmt.Println("Invalid admission ID format.")
		return errors.New("invalid admission ID")
	}
	opNumber, err := strconv.Atoi(opNo)
	if err != nil {
		fmt.Println("Invalid operation number format.")
		return errors.New("invalid operation number")
	}
	updatequery:=`UPDATE Admissions
	SET operations_number=$1
	WHERE admission_id=$2`
	_,err=db.Exec(updatequery,opNumber,admissionIDOp)
	if err!=nil{
		fmt.Println(errorspacket.ExecutionError)
		return errorspacket.ExecutionError
	}
	return nil
}
func UpdateAdmissionNotesAndPrice(admIDN,notesAD,priceAD string) error{
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()
	admissionIDN, err := strconv.Atoi(admIDN)
	if err != nil {
		fmt.Println("Invalid admission ID format.")
		return errors.New("invalid admission ID")
	}
	totalPriceAd, err := strconv.ParseFloat(priceAD, 64)
	if err != nil {
		fmt.Println("Invalid admission price")
		return err
	}
	updatequery:=`UPDATE Admissions
	SET doctor_notes=$1,price=$2
	WHERE admission_id=$3`
	_,err=db.Exec(updatequery,notesAD,totalPriceAd,admissionIDN)
	if err!=nil{
		fmt.Println(errorspacket.ExecutionError)
		return errorspacket.ExecutionError
	}
	return nil
}
func GetAdmissionsPatient(pat model.Patient) (string,error){
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "",err
	}
	defer db.Close()
	query:=`SELECT admission_id,patient_id,doctor_id,room_id,start_date,operations_number,discharge_date,status,doctor_notes,price
	FROM Admissions
	WHERE patient_id=$1`
	rows,err:=db.Query(query,pat.ID)
	if err!=nil{
		fmt.Println(errorspacket.QueryError)
		return "",err
	}
	defer rows.Close()
	var ListAdmissions strings.Builder
	for rows.Next() {
		var admID, patID, docID, roomID, opNo int
		var priceAdm float64
		var dischargeD sql.NullTime
		var startD time.Time
		var statAdm, docNotes string
		if err := rows.Scan(&admID, &patID, &docID, &roomID, &startD, &opNo,&dischargeD,&statAdm,&docNotes,&priceAdm); err != nil {
			return "", errorspacket.QueryError
		}
		ListAdmissions.WriteString(fmt.Sprintf("Admission, Patient, Doctor, Room ID's: %d %d %d %d\n", admID, patID, docID, roomID))
		ListAdmissions.WriteString(fmt.Sprintf("Start Date: [%s]\n", startD.Format("2006-01-02 15:04:05")))
		if dischargeD.Valid {
			ListAdmissions.WriteString(fmt.Sprintf("Discharge Date: [%s]\n", dischargeD.Time.Format("2006-01-02 15:04:05")))
		} else {
			ListAdmissions.WriteString("Discharge Date: N/A\n")
		}
		ListAdmissions.WriteString(fmt.Sprintf("Operations Number: %d\n",opNo))
		ListAdmissions.WriteString(fmt.Sprintf("Status: %s\n",statAdm))
		ListAdmissions.WriteString(fmt.Sprintf("Notes and total price: %s %.2f\n\n",docNotes, priceAdm))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return ListAdmissions.String(), nil
}