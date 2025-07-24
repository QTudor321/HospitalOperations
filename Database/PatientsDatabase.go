package database
import (
	"fmt"
	"database/sql"
	"HospitalQOps/model"
	"HospitalQOps/errorspacket"
	"bufio"
	"os"
	"time"
	"errors"
	"strings"
)
func LoginPatient(lastN string, firstN string, pass string) (model.Patient, error){
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return model.Patient{}, err
	}
	defer db.Close()
	var pat model.Patient
	query := `
		SELECT patient_id, last_name, first_name, age, gender, email, phone, address, credit
		FROM Patients
		WHERE last_name=$1 and first_name=$2
		LIMIT 1`
	err = db.QueryRow(query, lastN, firstN).Scan(&pat.ID, &pat.LastName, &pat.FirstName, &pat.Age, &pat.Gender, &pat.Email, &pat.Phone, &pat.Address, &pat.Credit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Database Credentials: ", errorspacket.PatientCredentialsError)
			return model.Patient{}, errorspacket.PatientCredentialsError
		}
		return model.Patient{}, err
	}
	file, err := os.Open("VaultLogins.txt")
	if err != nil {
		fmt.Println(errorspacket.VaultError)
		return model.Patient{}, errorspacket.VaultError
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

		if fileEmail == pat.Email && filePass == pass {
			return pat, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return model.Patient{}, errors.New("Error reading login file")
	}
	return model.Patient{}, errorspacket.PatientCredentialsError
}
func RegisterPatient(pat model.Patient) error{
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()

	query := `
		INSERT INTO Patients (last_name, first_name, age, gender, email, phone, address, credit)
		VALUES ($1, $2, $3, $4, $5, $6, $7, 700.00)`

	_, err = db.Exec(query, pat.LastName, pat.FirstName, pat.Age, pat.Gender, pat.Email, pat.Phone, pat.Address)
	if err != nil {
		fmt.Println(errorspacket.ExecutionError)
		return err
	}
	return nil
}
func GetDoctorsListByPatient(patient model.Patient) (string, error) {
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
		WHERE Patients.patient_id = $1
		ORDER BY Examinations.examination_date;
	`

	rows, err := db.Query(query, patient.ID)
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
func GetOperationsPatient(patient model.Patient) (string,error){
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "", err
	}
	defer db.Close()

	query := `
		SELECT * FROM Operations
		WHERE patient_id=$1
	`

	rows, err := db.Query(query, patient.ID)
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