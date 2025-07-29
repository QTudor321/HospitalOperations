package database

import (
	"HospitalQOps/errorspacket"
	"HospitalQOps/model"
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"strconv"
)
func LoginDoctor(i1 string, i2 string, i3 string) (model.Doctor, error) {
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return model.Doctor{}, err
	}
	defer db.Close()
	var doc model.Doctor
	query := `
		SELECT doctor_id, last_name, first_name, specialty, email, phone, department_id, is_Headdoctor, credit
		FROM Doctors
		WHERE last_name=$1 and first_name=$2
		LIMIT 1`
	err = db.QueryRow(query, i1, i2).Scan(&doc.ID, &doc.LastName, &doc.FirstName, &doc.Specialty, &doc.Email, &doc.Phone, &doc.DepartmentID, &doc.IsHeaddoctor, &doc.Credit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Database Credentials: ", errorspacket.DoctorCredentialsError)
			return model.Doctor{}, errorspacket.DoctorCredentialsError
		}
		return model.Doctor{}, err
	}
	file, err := os.Open("VaultLogins.txt")
	if err != nil {
		fmt.Println(errorspacket.VaultError)
		return model.Doctor{}, errorspacket.VaultError
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, ",")
		if len(parts) != 4 {
			continue
		}
		_ = strings.TrimSpace(parts[0])
		_ = strings.TrimSpace(parts[1])
		fileEmail := strings.TrimSpace(parts[2])
		filePass := strings.TrimSpace(parts[3])

		if fileEmail == doc.Email && filePass == i3 {
			return doc, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return model.Doctor{}, errors.New("Error reading login file")
	}
	return model.Doctor{}, errorspacket.DoctorCredentialsError
}

func RegisterDoctor(doc model.Doctor) error {
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()

	query := `
		INSERT INTO Doctors (last_name, first_name, specialty, email, phone, department_id, is_Headdoctor, credit)
		VALUES ($1, $2, $3, $4, $5, $6, false, 0.00)`

	_, err = db.Exec(query, doc.LastName, doc.FirstName, doc.Specialty, doc.Email, doc.Phone, doc.DepartmentID)
	if err != nil {
		fmt.Println(errorspacket.ExecutionError)
		return err
	}
	return nil
}
func RegisterHeaddoctor(doc model.Doctor) error{
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()
	query := `
		INSERT INTO Doctors (last_name, first_name, specialty, email, phone, department_id, is_Headdoctor, credit)
		VALUES ($1, $2, $3, $4, $5, $6, true, 10000.00)`

	_, err = db.Exec(query, doc.LastName, doc.FirstName, doc.Specialty, doc.Email, doc.Phone, doc.DepartmentID)
	if err != nil {
		fmt.Println(errorspacket.ExecutionError)
		return err
	}
	return nil
}
func GetPatientsListByDoctor(doctor model.Doctor) (string, error) {
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "", err
	}
	defer db.Close()

	query := `
		SELECT Patients.last_name, Patients.first_name,
		       Doctors.last_name, Doctors.first_name, Doctors.specialty,
		       Examinations.information
		FROM Patients
		INNER JOIN Appointments ON Appointments.patient_id = Patients.patient_id
		INNER JOIN Doctors ON Doctors.doctor_id = Appointments.doctor_id
		INNER JOIN Examinations ON Appointments.appointment_id = Examinations.appointment_id
		WHERE Doctors.doctor_id = $1
		ORDER BY Examinations.examination_date;
	`

	rows, err := db.Query(query, doctor.ID)
	if err != nil {
		fmt.Println(errorspacket.QueryError)
		return "", errorspacket.QueryError
	}
	defer rows.Close()

	var list strings.Builder

	for rows.Next() {
		var patLast, patFirst, docLast, docFirst, specialty, examInfo string
		if err := rows.Scan(&patLast, &patFirst, &docLast, &docFirst, &specialty, &examInfo); err != nil {
			return "", errorspacket.QueryError
		}
		list.WriteString(fmt.Sprintf("Patient: %s %s\n", patFirst, patLast))
		list.WriteString(fmt.Sprintf("Doctor: %s %s [%s]\n", docFirst, docLast, specialty))
		list.WriteString(fmt.Sprintf("Examination Info: %s\n", examInfo))
		list.WriteString("-----\n")
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return list.String(), nil
}
func GetPrescriptions() (string,error) {
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "",err
	}
	defer db.Close()
	query:=`SELECT * FROM Prescriptions`
	rows,err:=db.Query(query)
	if err!=nil{
		fmt.Println(errorspacket.QueryError)
		return "",errorspacket.QueryError
	}
	defer rows.Close()
	var ListP strings.Builder
	for rows.Next(){
		var presID, medNo int
		var examID sql.NullInt32
		var notes string
		var presDate time.Time
		var totalPrice float64
		if err:=rows.Scan(&presID,&examID,&medNo,&notes,&presDate,&totalPrice); err!=nil{
			return "",errorspacket.QueryError
		}
		ListP.WriteString(fmt.Sprintf("Prescription ID and Examination ID: %d %v\n", presID, examID))
		ListP.WriteString(fmt.Sprintf("Medications Number and Date: %d %s\n", medNo, presDate.Format("2006-01-02 15:04:05")))
		ListP.WriteString(fmt.Sprintf("Notes: %s\n", notes))
		ListP.WriteString(fmt.Sprintf("Total price: %.2f\n\n", totalPrice))

	}
	if err=rows.Err();err!=nil{
		return "",err
	}
	return ListP.String(),nil
}
func InsertPrescription(inputPID, inputPNot, inputPD, inputPPr string) error{
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()
	examID, err := strconv.Atoi(inputPID)
	if err != nil {
		fmt.Println("Invalid examination ID HERE")
		return err
	}
	prescriptionDate, err := time.Parse("2006-01-02 15:04:05", inputPD)
	if err != nil {
		fmt.Println("Invalid date format (expected YYYY-MM-DD)")
		return err
	}
	totalPrice, err := strconv.ParseFloat(inputPPr, 64)
	if err != nil {
		fmt.Println("Invalid total price")
		return err
	}
	instruction:=`INSERT INTO Prescriptions (examination_id,number_medications,notes,prescription_date,total_price)
	VALUES ($1, $2, $3, $4, $5)`
	_,err=db.Exec(instruction,examID,3,inputPNot,prescriptionDate,totalPrice)
	if err!=nil{
		fmt.Println(errorspacket.ExecutionError)
		return errorspacket.ExecutionError
	}
	return nil
}
func GetLastPrescription() (string, error) {
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "", err
	}
	defer db.Close()
	query := `SELECT prescription_id 
	FROM Prescriptions 
	ORDER BY prescription_id 
	DESC LIMIT 1`
	var lastPID int
	err = db.QueryRow(query).Scan(&lastPID)
	if err != nil {
		fmt.Println(errorspacket.QueryError)
		return "", errorspacket.QueryError
	}
	return fmt.Sprintf("%d", lastPID), nil
}
func InsertJunctionMedPrescription(LastPrescriptionID,med1,noteInput2 string) error{
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()
	lastpresID, err := strconv.Atoi(LastPrescriptionID)
	if err != nil {
		fmt.Println("Invalid prescription ID")
		return err
	}
	medicationID, err := strconv.Atoi(med1)
	if err != nil {
		fmt.Println("Invalid medication ID")
		return err
	}
	instruction:=`INSERT INTO prescriptions_medications (prescriptions_prescription_id, medications_medication_id, notes)
	VALUES ($1, $2, $3)`
	_,err=db.Exec(instruction,lastpresID,medicationID,noteInput2)
	if err!=nil{
		fmt.Println(errorspacket.ExecutionError)
		return errorspacket.ExecutionError
	}
	return nil
}
func DeletePrescription(inputDelID string) error{
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()
	presID, err := strconv.Atoi(inputDelID)
	if err != nil {
		fmt.Println("Invalid prescription ID format")
		return err
	}
	deletequery:=`DELETE From Prescriptions 
	WHERE prescription_id=$1`
	_,err=db.Exec(deletequery,presID)
	if err!=nil{
		fmt.Println(errorspacket.ExecutionError)
		return errorspacket.ExecutionError
	}
	return nil
}
func AverageNumberOfMedsPrescribed() (string, error) {
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "", err
	}
	defer db.Close()

	query := `SELECT AVG(number_medications) FROM Prescriptions`
	row := db.QueryRow(query)

	var avg float64
	if err := row.Scan(&avg); err != nil {
		fmt.Println(errorspacket.QueryError)
		return "", errorspacket.QueryError
	}

	return fmt.Sprintf("Average medications per prescription: %.2f", avg), nil
}
func GetAllDoctors() (string,error){
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "", err
	}
	defer db.Close()
	query := `SELECT * FROM Doctors`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(errorspacket.QueryError)
		return "", errorspacket.QueryError
	}
	defer rows.Close()

	var listD strings.Builder

	for rows.Next() {
		var docLast, docFirst, specialty, emailD, phone string
		var docID, depID int
		var isHead bool
		var credit float64
		if err := rows.Scan(&docID, &docLast, &docFirst, &specialty, &emailD, &phone, &depID, &isHead, &credit); err != nil {
			return "", errorspacket.QueryError
		}
		listD.WriteString(fmt.Sprintf("Doctor and Department ID: %d %d\n", docID, depID))
		listD.WriteString(fmt.Sprintf("Last and First Name: %s %s\n", docLast, docFirst))
		listD.WriteString(fmt.Sprintf("Specialty and Email: %s %s \n", specialty, emailD))
		listD.WriteString(fmt.Sprintf("Phone: (%s)\n", phone))
		listD.WriteString(fmt.Sprintf("Is headdoctor and credit: %t %.2f\n\n", isHead, credit))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return listD.String(), nil
}
func GetAllPatients() (string,error){
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "", err
	}
	defer db.Close()
	query := `SELECT * FROM Patients`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(errorspacket.QueryError)
		return "", errorspacket.QueryError
	}
	defer rows.Close()

	var listP strings.Builder

	for rows.Next() {
		var patLast, patFirst, gender, email, phone, address string
		var patID, age int
		var credit float64
		if err := rows.Scan(&patID, &patLast, &patFirst, &age, &gender, &email, &phone, &address, &credit); err != nil {
			return "", errorspacket.QueryError
		}
		listP.WriteString(fmt.Sprintf("Patient ID, Last and First Name: %d %s %s\n", patID, patLast, patFirst))
		listP.WriteString(fmt.Sprintf("Age, Gender, Email: %d %s %s\n", age, gender, email))
		listP.WriteString(fmt.Sprintf("Phone: (%s)\n", phone))
		listP.WriteString(fmt.Sprintf("Address: {%s}\n", address))
		listP.WriteString(fmt.Sprintf("credit: %.2f\n\n", credit))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return listP.String(), nil
}
func GetNursesDoctors() (string,error){
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "", err
	}
	defer db.Close()
	query := `SELECT * FROM nurses_doctors`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(errorspacket.QueryError)
		return "", errorspacket.QueryError
	}
	defer rows.Close()

	var listDN strings.Builder

	for rows.Next() {
		var notes string
		var dnID, nurseID, docID int
		if err := rows.Scan(&dnID, &nurseID, &docID, &notes); err != nil {
			return "", errorspacket.QueryError
		}
		listDN.WriteString(fmt.Sprintf("Identifier: %d\n", dnID))
		listDN.WriteString(fmt.Sprintf("Nurse and Doctor ID: %d %d\n", nurseID, docID))
		listDN.WriteString(fmt.Sprintf("Notes: %s\n", notes))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return listDN.String(), nil
}